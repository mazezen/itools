package iio

import (
	"encoding/binary"
	"io"
)

func (iio *ITcp) Encode(bytesBuffer io.Writer, content string) error {
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(iio.msgHeader)); err != nil {
		return err
	}
	clen := int32(len([]byte(content)))
	if err := binary.Write(bytesBuffer, binary.BigEndian, clen); err != nil {
		return err
	}

	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(content)); err != nil {
		return err
	}
	return nil
}
