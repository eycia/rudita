package executor

import (
	"time"

	"github.com/eycia/rudita/parser"
)

type Expr interface {
	Exec(values parser.ValueGetter, expr string) (interface{}, error)
}

type Point interface {
	Time() time.Time

	Value(field string) float64
	Tag(tagk string) string

	ValueForEach(func(field string, value float64))
	TagForEach(func(tagk string, tagv string))
}

// Stop by close feed
type Executor interface {
	Start(feed chan<- parser.ValueGetter)
	Sum() <-chan Point
}
