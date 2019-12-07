package basic

import (
	"time"
)

// RawLog is raw log
// XXX: string?
type RawLog string

type ValueKind int

const (
	Null ValueKind = iota
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
	Interface(key string) interface{}

	SetString(key string, value string) ValueGetter
	SetInt(key string, value int64) ValueGetter
	SetFloat(key string, value float64) ValueGetter
	SetBool(key string, value bool) ValueGetter
	SetInterface(key string, value interface{}) ValueGetter

	For(func(field string, kind ValueKind, value interface{}))

	ForEach(func(field string, value interface{}))

	ForEachKind(
		fString func(field string, value string),
		fInt func(field string, value int64),
		fFloat func(field string, value float64),
		fBool func(field string, value bool),
	)
}

type Point interface {
	Time() time.Time

	TagKey() []string
	MetricKey() []string

	Get(key string) interface{}
	Set(key string, value interface{})
}

type Transform interface {
	Push(v Point)
	Pull() []Point
}

type Parser interface {
	//XXX: []byte?
	Parse(line RawLog) ValueGetter
}
