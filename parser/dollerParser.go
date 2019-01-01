package parser

import (
	"fmt"
	"strings"
)

type DollerOption struct {
	Delimiter string
}

type Doller struct {
	option *DollerOption
}

func (p *Doller) Parse(line string) ValueGetter {
	values := map[string]string{}
	for i, field := range strings.Split(line, p.option.Delimiter) {
		// TODO: speed up it
		values[fmt.Sprintf("$%d", i+1)] = field
	}
	return MapValueGetter(values)
}

func NewDoller(option *DollerOption) Parser {
	return &Doller{
		option,
	}
}
