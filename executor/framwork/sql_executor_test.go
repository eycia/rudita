package framwork

import (
	"testing"
	"time"

	"github.com/eycia/rudita/executor"
	"github.com/eycia/rudita/executor/expr"
	"github.com/eycia/utils/testutil"
)

func TestSimpleSQL(t *testing.T) {
	so, when := testutil.Use(t)

	luaExprExecutor := expr.NewLuaExpr(&expr.LuaExprOption{})

	exprs := []struct {
		expr executor.Expr
		name string
	}{{luaExprExecutor, "LuaExpr"}}

	for _, e := range exprs {
		when("SimpleSQL with %s expr", e.name)

		simpleSQL := NewSimpleSQL(&SimpleSQLConfig{
			CaluExpr:            []string{"1", "DEVICE_TYPE(user_agent)", "http_code, traffic"},
			CaluAs:              []string{"requests", "device", "http_code", "traffic"},
			Where:               `device == "MACOS" or device == "LINUX"`,
			GroupBy:             []string{"device", "http_code"},
			GroupbyAggrExpr:     []string{"SUM(requests)", "MAX(traffic)", "MIN(traffic)", "AVG(traffic)"},
			GroupbyAggrAs:       []string{"sum_req", "max_tr", "min_tr", "avg_tr"},
			Having:              `sum_req != 1`,
			GroupbyTimeExpr:     `PARSE_TIME(time, "200601021504")`,
			GroupbyTimeDuration: time.Minute * 5,
			ExprExecutor:        e.expr,
		})

		so.NotNil(simpleSQL, "NewSimpleSQL should success")
		when("when feed some log, Sum() should success")
		logs := `
MACOS 200 10000 201902122301
MACOS 200 20100 201902122301
LINUX 200 20200 201902122302
LINUX 200 30200 201902122302
LINUX 501 20200 201902122302
LINUX 501 30200 201902122302
LINUX 200 20200 201902122307
LINUX 200 30200 201902122308
LINUX 501 20200 201902122309
LINUX 501 30200 201902122310
LINUX 599 11111 201912052200
WINDO 200 20200 201902122303
WINDO 200 31000 201902122303
`
	}
}
