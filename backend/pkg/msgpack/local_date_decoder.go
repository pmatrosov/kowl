package msgpack

import (
	"encoding/binary"
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

func init() {
	msgpack.RegisterExtDecoder(18, "", localDateDecoder)
}

func localDateDecoder(d *msgpack.Decoder, v reflect.Value, extLen int) error {
	b := make([]byte, extLen)

	if err := d.ReadFull(b); err != nil {
		v.SetString("?error?")
		return err
	}

	v.SetString(readLocalDate(b, 0))
	return nil
}

func readLocalDate(b []byte, offset int) string {
	year := binary.LittleEndian.Uint16(b[offset:])
	month := b[offset+2]
	day := b[offset+3]
	return fmt.Sprintf("%04d/%02d/%02d", year, month, day)
}
