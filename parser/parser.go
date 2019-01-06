package parser

type Kind int

const (
	Null = iota
	Int
	Float
	String
	Time
)

// Log is always string first.
type ValueGetter interface {
	String(field string) string
	For(func(field, value string))
}

type Parser interface {
	//XXX: []byte?
	Parse(line string) ValueGetter
}

type MapValueGetter map[string]string

func (p MapValueGetter) String(field string) string {
	return p[field]
}

func (p MapValueGetter) For(f func(field, value string)) {
	for k, v := range p {
		f(k, v)
	}
}
