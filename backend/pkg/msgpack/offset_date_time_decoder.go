package msgpack

import (
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

func init() {
	msgpack.RegisterExtDecoder(22, "", offsetDateTimeDecoder)
}

func offsetDateTimeDecoder(d *msgpack.Decoder, v reflect.Value, extLen int) error {
	b := make([]byte, extLen)
	err := d.ReadFull(b)
	if err != nil {
		v.SetString("?error?")
		return err
	}

	v.SetString(readOffsetDateTime(b, 0))
	return nil
}

func readOffsetDateTime(b []byte, offset int) string {
	return readLocalDateTime(b, offset) + readOffset(b, offset+12)
}
