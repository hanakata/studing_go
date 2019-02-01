package trace

import "io"

type Tracer interface {
	//interfaceなので任意の型の引数を何個でも受け取る
	Trace(...interface{})
}

func New(w io.Writer) Tracer {
	return nil
}
