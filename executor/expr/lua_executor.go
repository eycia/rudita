package expr

import (
	"fmt"

	"github.com/eycia/rudita/parser"
	"github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

const luaFunc = `
function _calu_expr(v)
	%s
end
`

type Lua struct{}

func (p *Lua) Expr(values parser.ValueGetter, expr string) interface{} {
	L := lua.NewState()
	defer L.Close()

	expandLuaScript := fmt.Sprintf(luaFunc, expr)
	logrus.Debug("expandLuaScript is:", expandLuaScript)

	table := L.NewTable()
	values.For(func(k, v string) {
		table.RawSetString(k, lua.LString(v))
	})

	if err := L.DoString(expandLuaScript); err != nil {
		panic(err)
	}
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("_calu_expr"),
		NRet:    1,
		Protect: true,
	}, table); err != nil {
		panic(err)
	}
	ret := L.Get(-1) // returned value
	L.Pop(1)         // remove received value
	return ret
}
