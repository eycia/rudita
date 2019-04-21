package framwork

import (
	"time"

	"github.com/eycia/rudita/executor"
	"github.com/eycia/rudita/parser"
)

// like sql:
//  select group_by_aggr, group_by_field... from (
//        select expr1 as as1
//               expr2 as as2
//  ) where where1
//  group by group_by_field..., time(expr, by)
//  having
type SimpleSQLConfig struct {
	CaluExpr []string
	CaluAs   []string

	Where string

	GroupBy []string

	GroupbyAggrExpr []string
	GroupbyAggrAs   []string

	Having              string
	GroupbyTimeExpr     string
	GroupbyTimeDuration time.Duration

	ExprExecutor executor.Expr
}

type SimpleSQL struct {
	conf *SimpleSQLConfig
}

func NewSimpleSQL(config *SimpleSQLConfig) *SimpleSQL {
	return &SimpleSQL{config}
}

func (p *SimpleSQL) Start(feed chan<- parser.ValueGetter) {

}

func (p *SimpleSQL) Sum() <-chan executor.Point {
	return nil
}
