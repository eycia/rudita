package aggr

import (
	"fmt"
	"time"

	"github.com/eycia/rudita/basic"
)

type aggr struct {
	timeKey   string
	tagKey    []string
	metricKey []string

	timeRound time.Duration

	apply map[string] /*metric*/ func(sum interface{}, new basic.Point) interface{}

	cache map[string]basic.Point
}

func (p *aggr) Pull() []basic.Point {
	result := make([]basic.Point, 0, len(p.cache))
	for _, point := range p.cache {
		result = append(result, point)
	}
	return result
}

func (p *aggr) Push(value basic.Point) {
	roundedTime := value.Time().Round(p.timeRound)

	tagvs := make([]interface{}, len(p.tagKey))
	for i, tag := range p.tagKey {
		tagvs[i] = value.Get(tag)
	}
	tagCacheKey := fmt.Sprintln(roundedTime, tagvs) // TODO, tagvs...

	if p.cache[tagCacheKey] == nil {
		point := basic.NewEmptyPoint(p.timeKey, p.tagKey, p.metricKey)

		for i, tag := range p.tagKey {
			point.Set(tag, tagvs[i])
		}

		point.Set(
			p.timeKey,
			roundedTime,
		)
	}

	point := p.cache[tagCacheKey]

	for _, metric := range p.metricKey {
		point.Set(
			metric,
			p.apply[metric](point.Get(metric), value),
		)
	}
}
