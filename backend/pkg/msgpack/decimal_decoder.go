package msgpack

import (
	"encoding/binary"
	"github.com/vmihailenco/msgpack/v5"
	"math/big"
	"reflect"
	"strings"
)

func init() {
	msgpack.RegisterExtDecoder(16, "", decimalDecoder)
}

func decimalDecoder(d *msgpack.Decoder, v reflect.Value, extLen int) error {
	b := make([]byte, extLen)
	err := d.ReadFull(b)
	if err != nil {
		v.SetString("?error?")
		return err
	}

	v.SetString(readDecimal(b, extLen))
	return nil
}

func readDecimal(b []byte, extLen int) string {
	sign := b[0] > 0
	scale := int(b[1])

	x := make([]byte, 12) // 0..3 4..7 8..11

	if 2+4 <= extLen {
		binary.BigEndian.PutUint32(x[8:], binary.LittleEndian.Uint32(b[2:]))
	}

	if 2+4+4 <= extLen {
		binary.BigEndian.PutUint32(x[4:], binary.LittleEndian.Uint32(b[6:]))
	}

	if 2+4+4+4 <= extLen {
		binary.BigEndian.PutUint32(x[0:], binary.LittleEndian.Uint32(b[10:]))
	}

	d := new(big.Int)
	d.SetBytes(x)
	s := d.String()
	if scale > 0 {
		index := len(s) - scale
		if index > 0 {
			s = s[:index] + "." + s[index:]
		} else {
			s = "0." + strings.Repeat("0", -index) + s
		}
	}
	if sign {
		s = "-" + s
	}

	return s
}
