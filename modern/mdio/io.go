package mdio

import (
	"bytes"
	"io"
	"io/ioutil"
)

func Tap(r io.ReadCloser, fn func(text string)) io.ReadCloser {
	buf, _ := ioutil.ReadAll(r)
	text := string(buf)

	fn(text)

	return ioutil.NopCloser(bytes.NewBuffer(buf))
}
