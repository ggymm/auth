package auth

import (
	"bytes"
	"encoding/gob"
)

func encode[T any](v *T) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := gob.NewEncoder(buf).Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decode[T any](data []byte, v *T) error {
	buf := bytes.NewBuffer(data)
	err := gob.NewDecoder(buf).Decode(v)
	if err != nil {
		return err
	}
	return nil
}
