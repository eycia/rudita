package basic

import (
	"time"
)

type simplePoint struct {
	timeKey   string
	tagKey    []string
	metricKey []string

	store ValueGetter
}

func NewEmptyPoint(timeKey string, tagKey []string, metricKey []string) Point {
	return &simplePoint{
		timeKey:   timeKey,
		tagKey:    tagKey,
		metricKey: metricKey,
		store:     MapValueGetter{},
	}
}

func (p *simplePoint) Time() time.Time {
	return p.store.Interface(p.timeKey).(time.Time)
}

func (p *simplePoint) TagKey() []string {
	return p.tagKey
}

func (p *simplePoint) MetricKey() []string {
	return p.metricKey
}

func (p *simplePoint) Get(key string) interface{} {
	return p.store.Interface(key)
}

func (p *simplePoint) Set(key string, value interface{}) {
	p.store.SetInterface(key, value)
}
