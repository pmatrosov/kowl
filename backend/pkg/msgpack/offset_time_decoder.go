package msgpack

import (
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

func init() {
	msgpack.RegisterExtDecoder(24, "", offsetTimeDecoder)
}

func offsetTimeDecoder(d *msgpack.Decoder, v reflect.Value, extLen int) error {
	b := make([]byte, extLen)
	err := d.ReadFull(b)
	if err != nil {
		v.SetString("?error?")
		return err
	}

	v.SetString(readOffsetTime(b))
	return nil
}

func readOffsetTime(b []byte) string {
	return readLocalTime(b, 0) + readOffset(b, 8)
}
