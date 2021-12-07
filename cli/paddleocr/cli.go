package paddleocr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	// "github.com/icecraft/ocr/pkg/types"
)

const (
	UserCenterCookie = "cookie"
)

var (
	ErrIllegalToken  = errors.New("illegal token")
	ErrTooManyValues = errors.New("too many values")
)

type rawResponse struct {
	ErrNo  int      `json:"err_no,omitempty" bson:"err_no,omitempty"`
	ErrMsg string   `json:"err_msg,omitempty" bson:"err_msg,omitempty"`
	Key    []string `json:"key,omitempty" bson:"key,omitempty"`
	Value  []string `json:"value,omitempty" bson:"value,omitempty"`
}

type Response rawResponse

type Req struct {
	Value [][]byte `json:"value,omitempty" bson:"value,omitempty"`
	Key   []string `json:"key,omitempty" bson:"key,omitempty"`
}

type Cli struct {
	endpoint string
}

func NewCli(endpoint string) *Cli {
	return &Cli{endpoint: endpoint}
}

func (o *Cli) Rec(ctx context.Context, r io.Reader) (*Response, error) {
	endpoint := fmt.Sprintf("%s/%s", o.endpoint, "ocr/prediction")
	var res rawResponse

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	req := Req{
		Key:   []string{"image"},
		Value: [][]byte{data},
	}
	if err := o.doPost(ctx, endpoint, &req, &res); err != nil {
		return nil, err
	}
	return rawRes2Res(&res)
}

// common function
func (o *Cli) setCookie(ctx context.Context, req *http.Request) {
	cookie := ctx.Value(UserCenterCookie)
	req.Header.Add("Cookie", fmt.Sprintf("%v", cookie))
}

//  request functions
func (o *Cli) doPost(ctx context.Context, endpoint string, postData interface{}, respBody interface{}) error {
	data, err := json.Marshal(postData)
	if err != nil {
		return err
	}
	postBody := bytes.NewBuffer(data)

	req, err := http.NewRequest("POST", endpoint, postBody)
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")
	o.setCookie(ctx, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, respBody)
}

func (o *Cli) doGet(ctx context.Context, endpoint string, respBody interface{}) error {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")
	o.setCookie(ctx, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, respBody)
}

// help functions
func rawRes2Res(in *rawResponse) (*Response, error) {
	ret := Response{ErrNo: in.ErrNo, ErrMsg: in.ErrMsg}
	if in.Value != nil {
		if len(in.Value) == 1 {
			ret.Value = strings.Split(strings.Trim(in.Value[0], "[]"), ",")
		} else if len(in.Value) > 1 {
			return nil, ErrTooManyValues
		}
	}
	return &ret, nil
}
