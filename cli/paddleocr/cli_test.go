package paddleocr

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
