package notuse

import (
	"github.com/eycia/rudita/executor/map/parser"
)

type Expr interface {
	Exec(values parser.ValueGetter, expr string) (interface{}, error)
}


