package z

import (
	"io"
	"reflect"
)

type RunnableInterface interface {
	Run() error
}

type NetReaderInterface interface {
	ReadString() (string, error)
}

type NetWriterInterface interface {
	io.WriteCloser
	WriteString(string) error
	Bytes() []byte
}

type (
	DictionaryTypeHandlerFn func(op DictionaryOp, name, alias string, value reflect.Value)
	DictionaryOp            int
)

const (
	HydrateOp DictionaryOp = iota + 1
	DehydrateOp
)

type DictionaryInterface interface {
	RegisterTypeHandler(typ any, handler DictionaryTypeHandlerFn)
	Hydrate(target interface{}) error
	Dehydrate(target interface{}) error
}

type MemoizerInterface interface {
	Run(args ...any) (v any, err error)
}
