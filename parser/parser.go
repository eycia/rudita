package parser

type Kind int

const (
	Null Kind = iota
	String
	Int
	Float
	Bool
	//Time
)

type ValueGetter interface {
	String(field string) (string, bool)
	Int(field string) (int64, bool)
	Float(field string) (float64, bool)
	Bool(field string) (bool, bool)

	SetString(key string, value string) ValueGetter
	SetInt(key string, value int64) ValueGetter
	SetFloat(key string, value float64) ValueGetter
	SetBool(key string, value bool) ValueGetter

	For(func(field string, kind Kind, value interface{}))

	ForEachKind(
		fString func(field string, value string),
		fInt func(field string, value int64),
		fFloat func(field string, value float64),
		fBool func(field string, value bool),
	)
}
type Parser interface {
	//XXX: []byte?
	Parse(line string) ValueGetter
}

type MapValueGetter map[string]interface{}

func (p MapValueGetter) SetString(key string, value string) ValueGetter {
	p[key] = value
	return p
}

func (p MapValueGetter) SetInt(key string, value int64) ValueGetter {
	p[key] = value
	return p
}

func (p MapValueGetter) SetFloat(key string, value float64) ValueGetter {
	p[key] = value
	return p
}

func (p MapValueGetter) SetBool(key string, value bool) ValueGetter {
	p[key] = value
	return p
}

func (p MapValueGetter) String(field string) (string, bool) {
	s, ok := p[field].(string)
	return s, ok
}

func (p MapValueGetter) Int(field string) (int64, bool) {
	i, ok := p[field].(int64)
	return i, ok
}

func (p MapValueGetter) Float(field string) (float64, bool) {
	f, ok := p[field].(float64)
	return f, ok
}

func (p MapValueGetter) Bool(field string) (bool, bool) {
	b, ok := p[field].(bool)
	return b, ok
}

func (p MapValueGetter) For(f func(field string, kind Kind, value interface{})) {
	for k, v := range p {
		switch vv := v.(type) {
		case string:
			f(k, String, vv)
		case int64:
			f(k, Int, vv)
		case float64:
			f(k, Float, vv)
		case bool:
			f(k, Bool, vv)
		}
		// skip if not match
	}
}

func (p MapValueGetter) ForEachKind(
	fString func(field string, value string),
	fInt func(field string, value int64),
	fFloat func(field string, value float64),
	fBool func(field string, value bool),
) {
	for k, v := range p {
		switch vv := v.(type) {
		case string:
			fString(k, vv)
		case int64:
			fInt(k, vv)
		case float64:
			fFloat(k, vv)
		case bool:
			fBool(k, vv)
		}
		// skip if not match
	}
}
