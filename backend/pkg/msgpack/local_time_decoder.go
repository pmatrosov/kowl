package msgpack

import (
	"encoding/binary"
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

func init() {
	msgpack.RegisterExtDecoder(20, "", localTimeDecoder)
}

func localTimeDecoder(d *msgpack.Decoder, v reflect.Value, extLen int) error {
	b := make([]byte, extLen)

	if err := d.ReadFull(b); err != nil {
		v.SetString("?error?")
		return err
	}

	v.SetString(readLocalTime(b, 0))
	return nil
}

func readLocalTime(b []byte, offset int) string {
	nanos := binary.LittleEndian.Uint64(b[offset:])
	seconds := nanos / 1_000_000_000
	minutes := seconds / 60
	hours := minutes / 60
	minutes = minutes % 60
	seconds = seconds % 60
	nanos = nanos % 1_000_000_000
	return fmt.Sprintf("%02d:%02d:%02d.%06d", hours, minutes, seconds, nanos)
}
