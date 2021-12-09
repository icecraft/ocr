package types

import (
	"github.com/hbollon/go-edlib"
)

type HttpResult struct {
	Code   int         `json:"code,omitempty" bson:"code,omitempty"`
	ErrMsg string      `json:"err_msg,omitempty" bson:"err_msg,omitempty"`
	Record interface{} `json:"record,omitempty" bson:"record,omitempty"`
}

func getEditDistance(target []string, src []string) [][]int {
	ret := make([][]int, len(target))

	for i := range target {
		tmp := make([]int, len(src))
		for j := range src {
			tmp[j] = edlib.LevenshteinDistance(target[i], src[j])
		}
		ret[i] = tmp
	}
	return ret
}
