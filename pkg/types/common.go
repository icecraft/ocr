package types

type HttpResult struct {
	Code   int         `json:"code,omitempty" bson:"code,omitempty"`
	ErrMsg string      `json:"err_msg,omitempty" bson:"err_msg,omitempty"`
	Record interface{} `json:"record,omitempty" bson:"record,omitempty"`
}
