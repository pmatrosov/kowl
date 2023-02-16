package msgpack

import (
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

func init() {
	msgpack.RegisterExtDecoder(21, "", offsetDateDecoder)
}

func offsetDateDecoder(d *msgpack.Decoder, v reflect.Value, extLen int) error {
	b := make([]byte, extLen)
	err := d.ReadFull(b)
	if err != nil {
		v.SetString("?error?")
		return err
	}

	v.SetString(readOffsetDate(b, 0))
	return nil
}

func readOffsetDate(b []byte, offset int) string {
	return readLocalDate(b, 0) + readOffset(b, 4)
}
