package gober

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func init() {
	gob.Register(&param.LetsTalkRequest{})
}

func Encoder(value interface{}) ([]byte, error) {

	var network bytes.Buffer
	enc := gob.NewEncoder(&network)

	err := enc.Encode(&value)
	if err != nil {
		return nil, fmt.Errorf("gob encoder error : %w", err)
	}

	return network.Bytes(), nil
}

func Decoder(data []byte, param interface{}) (interface{}, error) {
	network := bytes.NewBuffer(data)

	dec := gob.NewDecoder(network)

	err := dec.Decode(&param)
	if err != nil {
		return nil, fmt.Errorf("gob decoder error : %w", err)
	}

	return param, nil
}
