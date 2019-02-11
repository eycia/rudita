package parser

import (
	"fmt"
	"strings"
)

type DollarOption struct {
	Delimiter string
}

type Dollar struct {
	option *DollarOption
}

func (p *Dollar) Parse(line string) ValueGetter {
	values := MapValueGetter{}
	for i, fieldv := range strings.Split(line, p.option.Delimiter) {
		// TODO: speed up it
		values.SetString(fmt.Sprintf("$%d", i+1), fieldv)
	}
	return MapValueGetter(values)
}

func NewDollar(option *DollarOption) Parser {
	return &Dollar{
		option,
	}
}
