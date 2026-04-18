package iio

import (
	"encoding/binary"
	"errors"
	"io"
)

func (iio *ITcp) Decode(bytesBuffer io.Reader) (bodyBuf []byte, err error) {
	headBuf := make([]byte, len(iio.msgHeader))
	if _, err := io.ReadFull(bytesBuffer, headBuf); err != nil {
		return nil, err
	}

	if string(headBuf) != iio.msgHeader {
		return nil, errors.New("data buff head invalid")
	}

	lbuf := make([]byte, 4)
	if _, err := io.ReadFull(bytesBuffer, lbuf); err != nil {
		return nil, err
	}

	l := binary.BigEndian.Uint32(lbuf)
	body := make([]byte, l)
	if _, err := io.ReadFull(bytesBuffer, body); err != nil {
		return nil, err
	}
	return body, err
}
