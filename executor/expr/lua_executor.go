package expr

import (
	"fmt"

	"github.com/eycia/rudita/parser"
	"github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

const luaFunc = `
function _calu_expr()
	return %s
end
`

type LuaExpr struct{}

type LuaExprOption struct{}

func NewLuaExpr(option *LuaExprOption) *LuaExpr {
	return &LuaExpr{}
}

func (p *LuaExpr) Exec(values parser.ValueGetter, expr string) (interface{}, error) {
	L := lua.NewState()
	defer L.Close()

	expandLuaScript := fmt.Sprintf(luaFunc, expr)
	logrus.Debug("expandLuaScript is:", expandLuaScript)

	values.ForEachKind(
		func(k string, v string) {
			L.SetGlobal(k, lua.LString(v))
		}, func(k string, v int64) {
			L.SetGlobal(k, lua.LNumber(v))
		}, func(k string, v float64) {
			L.SetGlobal(k, lua.LNumber(v))
		}, func(k string, v bool) {
			L.SetGlobal(k, lua.LBool(v))
		},
	)

	if err := L.DoString(expandLuaScript); err != nil {
		panic(err)
	}

	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("_calu_expr"),
		NRet:    1,
		Protect: true,
	}); err != nil {
		return nil, err
	}
	ret := L.Get(-1) // returned value
	L.Pop(1)         // remove received value
	return ret, nil
}
