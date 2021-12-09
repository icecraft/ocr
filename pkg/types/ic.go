package types

import (
	"errors"

	"github.com/icecraft/ocr/pkg/utils"
)

var (
	iTaxCardTargets     = []string{"Permanent Account Number Card", "/Name", "/FathersName", "/Date of Birth"}
	ErrOcrFieldNotFound = errors.New("ocr rec field not found")
)

type ITaxCard struct {
	Name              string `json:"name,omitempty" bson:"name,omitempty"`
	FatherName        string `json:"father_name,omitempty" bson:"father_name,omitempty"`
	ID                string `json:"id,omitempty" bson:"id,omitempty"`
	BirthDay          string `json:"birth_day,omitempty" bson:"birth_day,omitempty"`
	TotalEditDistance int    `json:"total_edit_distance,omitempty" bson:"total_edit_distance,omitempty"`
}

func NewITaxCardFromArr(arr []string) (*ITaxCard, error) {
	dis := getEditDistance(iTaxCardTargets, arr)
	mIdx := len(arr) - 1

	ret := ITaxCard{}

	// id
	pos, n := utils.MinIntArrWithPos(dis[0])
	if pos > -1 {
		if mIdx > pos {
			ret.ID = arr[pos+1]
		} else {
			return nil, ErrOcrFieldNotFound
		}
		ret.TotalEditDistance += n
	}

	// name
	pos, n = utils.MinIntArrWithPos(dis[1])
	if pos > -1 {
		if mIdx > pos {
			ret.Name = arr[pos+1]
		} else {
			return nil, ErrOcrFieldNotFound
		}
		ret.TotalEditDistance += n
	}

	// papa
	pos, n = utils.MinIntArrWithPos(dis[2])
	if pos > -1 {
		if mIdx > pos {
			ret.FatherName = arr[pos+1]
		} else {
			return nil, ErrOcrFieldNotFound
		}
		ret.TotalEditDistance += n
	}

	// Birth
	pos, n = utils.MinIntArrWithPos(dis[3])
	if pos > -1 {
		if mIdx > pos {
			ret.BirthDay = arr[pos+1]
		} else {
			return nil, ErrOcrFieldNotFound
		}
		ret.TotalEditDistance += n
	}
	return &ret, nil
}
