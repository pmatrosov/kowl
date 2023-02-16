package msgpack

import (
	"encoding/binary"
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
	"time"
)

func init() {
	msgpack.RegisterExtDecoder(17, "", instantDecoder)
}

func instantDecoder(d *msgpack.Decoder, v reflect.Value, extLen int) error {
	b := make([]byte, extLen)

	if err := d.ReadFull(b); err != nil {
		v.SetString("?error?")
		return err
	}

	v.SetString(readInstant(b))
	return nil
}

func readInstant(b []byte) string {
	seconds := binary.LittleEndian.Uint64(b)
	nanos := binary.LittleEndian.Uint32(b[8:])

	t := time.Unix(int64(seconds), int64(nanos)).UTC()

	return fmt.Sprintf("%04d/%02d/%02dT%02d:%02d:%02d.%06dZ", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
}
