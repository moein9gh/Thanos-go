package jsoner

import (
	"sync"

	jsoniter "github.com/json-iterator/go"
)

var j jsoniter.API
var once sync.Once

func get() jsoniter.API {
	once.Do(func() {
		j = jsoniter.ConfigCompatibleWithStandardLibrary
	})
	return j
}

func MarshalToString(v interface{}) (string, error) {
	return get().MarshalToString(v)
}

func Marshal(v interface{}) ([]byte, error) {
	return get().Marshal(v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return get().MarshalIndent(v, prefix, indent)
}

func UnmarshalFromString(str string, v interface{}) error {
	return get().UnmarshalFromString(str, v)
}

func Unmarshal(data []byte, v interface{}) error {
	return get().Unmarshal(data, v)
}
