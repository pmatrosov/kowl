package msgpack

import (
	"github.com/google/uuid"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

func init() {
	msgpack.RegisterExtDecoder(30, "", guidDecoder)
}

func guidDecoder(d *msgpack.Decoder, v reflect.Value, extLen int) error {
	b := make([]byte, extLen)
	err := d.ReadFull(b)
	if err != nil {
		v.SetString("?error?")
		return err
	}

	v.SetString(readGuid(b))
	return nil
}

func readGuid(b []byte) string {
	g, err := uuid.FromBytes(b)
	if err != nil {
		return "?guid?"
	}
	return g.String()
}
