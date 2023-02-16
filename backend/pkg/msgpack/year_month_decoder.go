package msgpack

import (
	"encoding/binary"
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

func init() {
	msgpack.RegisterExtDecoder(25, "", yearMonthDecoder)
}

func yearMonthDecoder(d *msgpack.Decoder, v reflect.Value, extLen int) error {
	b := make([]byte, extLen)
	err := d.ReadFull(b)
	if err != nil {
		v.SetString("?error?")
		return err
	}

	v.SetString(readYearMonth(b, 0))
	return nil
}

func readYearMonth(b []byte, offset int) string {
	year := binary.LittleEndian.Uint16(b[offset:])
	month := b[offset+2]
	return fmt.Sprintf("%04d/%02d", year, month)
}
