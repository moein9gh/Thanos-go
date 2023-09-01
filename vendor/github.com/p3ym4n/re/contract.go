package re

type Error interface {
	Kind() Kind
	Message() string
	Internal() error

	RawMap() map[string]interface{}
	ProcessedMap() map[string]string

	Chain(op Op) Error
	ChainWithMeta(op Op, meta Meta) Error
}
