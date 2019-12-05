package impl

import (
	"fmt"

	"github.com/eycia/rudita/executor"
	"github.com/eycia/rudita/parser"
)

type simplePoint struct {
	timeKey string
	tags    map[string]interface{}
	values  map[string]interface{}
}

type reduceExecutor struct {
	time  string
	tags  map[string]int // tag key to tag index
	apply map[string] /*metric*/ func(sum interface{}, new interface{}) interface{}

	cache map[string]*simplePoint
}

func (p *reduceExecutor) Start(feed chan<- parser.ValueGetter) {
	for value := range feed {
		tagvs := make([]interface{}, len(p.tags))
		for tag, i := range p.tags {
			tagvs[i] = value.Interface(tag)
		}
		tagCacheKey := fmt.Sprintln(tagvs...)

		if p.cache[tagCacheKey] == nil {
			revTags := map[string]interface{}{}
			for tag, i := range p.tags {
				revTags[tag] = tagvs[i]
			}

			p.cache[tagCacheKey] = &simplePoint{
				timeKey: p.time,
				tags:    revTags,
				values:  map[string]interface{}{},
			}
		}

		point := p.cache[tagCacheKey]

		value.For(func(field string, kind parser.Kind, value interface{}) {
			if _, ok := p.tags[field]; ok { // is tag
				return
			}
			if f, ok := p.apply[field]; ok {
				point.values[field] = f(point.values[field], value)
			}
		})
	}
}

func (p *reduceExecutor) Sum() <-chan executor.Point {
	panic("implement me")
}
