package parser

import (
	"fmt"
	"strings"

	"github.com/eycia/rudita/basic"
)

type DollarOption struct {
	Delimiter string
}

type Dollar struct {
	option *DollarOption
}

func (p *Dollar) Parse(line basic.RawLog) basic.ValueGetter {
	values := basic.MapValueGetter{}
	for i, fieldv := range strings.Split(string(line), p.option.Delimiter) {
		// TODO: speed up it
		values.SetString(fmt.Sprintf("$%d", i+1), fieldv)
	}
	return values
}

func NewDollar(option *DollarOption) basic.Parser {
	return &Dollar{
		option,
	}
}
