package msgpack

import (
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

func init() {
	msgpack.RegisterExtDecoder(19, "", localDateTimeDecoder)
}

func localDateTimeDecoder(d *msgpack.Decoder, v reflect.Value, extLen int) error {
	b := make([]byte, extLen)

	if err := d.ReadFull(b); err != nil {
		v.SetString("?error?")
		return err
	}

	v.SetString(readLocalDateTime(b, 0))
	return nil
}

func readLocalDateTime(b []byte, offset int) string {
	return readLocalDate(b, offset) + "T" + readLocalTime(b, offset+4)
}
