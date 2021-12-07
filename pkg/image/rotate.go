package image

import (
	"bytes"
	"context"
	"encoding/base64"
	"image"
	_ "image/jpeg"
	"image/png"
	"io"

	"github.com/disintegration/imaging"
)

func Rotate90(ctx context.Context, r io.Reader) (io.Reader, error) {
	return doRotate(ctx, r, imaging.Rotate90)
}

func Rotate180(ctx context.Context, r io.Reader) (io.Reader, error) {
	return doRotate(ctx, r, imaging.Rotate180)
}

func Rotate270(ctx context.Context, r io.Reader) (io.Reader, error) {
	return doRotate(ctx, r, imaging.Rotate270)
}

func doRotate(ctx context.Context, r io.Reader, proc func(image.Image) *image.NRGBA) (io.Reader, error) {
	m, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	nm := proc(m)

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, nm); err != nil {
		return nil, err
	}

	ret := make([]byte, 0)
	base64.StdEncoding.Encode(ret, buf.Bytes())
	return bytes.NewReader(ret), nil
}
