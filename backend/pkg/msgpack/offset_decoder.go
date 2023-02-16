package msgpack

import (
	"encoding/binary"
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

func init() {
	msgpack.RegisterExtDecoder(23, "", offsetDecoder)
}

func offsetDecoder(d *msgpack.Decoder, v reflect.Value, extLen int) error {
	b := make([]byte, extLen)

	err := d.ReadFull(b)
	if err != nil {
		v.SetString("?error?")
		return err
	}

	v.SetString(readOffset(b, 0))
	return nil
}

func readOffset(b []byte, offset int) string {

	var sign string
	if b[offset] > 0 {
		sign = "-"
	} else {
		sign = "+"
	}

	seconds := binary.LittleEndian.Uint16(b[offset+1:])
	minutes := seconds / 60
	hours := minutes / 60
	minutes = minutes % 60
	seconds = seconds % 60

	if seconds == 0 {
		if minutes == 0 {
			return fmt.Sprintf("%v%02d", sign, hours)
		} else {
			return fmt.Sprintf("%v%02d:%02d", sign, hours, minutes)
		}
	}
	return fmt.Sprintf("%v%02d:%02d:%02d", sign, hours, minutes, seconds)
}
