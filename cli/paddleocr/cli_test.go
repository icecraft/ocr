package paddleocr

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	s = []string{"fT", "3eR", "中口", "ER", "INCOMETAXDEPARTMENT", "GOVTOFINDLA", "taRel teaal",
		"Permanent Account Number Card", "ATHPL6094K", "/Name", "LAKHADE CHANDOJI NAMDEC",
		"uaTTE/FathersName", "NAMDEO LAKHADE", "Harfrg/Date of Birll", "分", "01/01/1992", "Chda", "RH8R/Slgnalure"}
)

func TestPaddle(t *testing.T) {
	cli := NewCli("http://127.0.0.1:9998")
	fd, err := os.Open("visa.jpeg")
	assert.NoError(t, err)

	t.Run("ocr", func(t *testing.T) {
		ret, err := cli.Rec(context.Background(), fd)
		assert.NoError(t, err)
		fmt.Printf("%v\n", ret)
	})
}
