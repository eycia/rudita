package expr

import (
	"github.com/eycia/rudita/executor"
	"github.com/eycia/rudita/parser"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLuaExpr_Exec(t *testing.T) {
	Convey("new a LuaExpr should ok", t, func() {
		lua := NewLuaExpr(nil)
		So(lua, ShouldNotBeNil)
		So(lua, ShouldImplement, (*executor.Expr)(nil))

		Convey("when ValueGetter is empty", func() {

			emptyValueGetter := parser.MapValueGetter{}
			So(emptyValueGetter, ShouldNotBeNil)

			Convey("LuaExpr should Exec() some expr correct", func() {
				cases := []struct {
					expr   string
					result interface{}
				}{
					{`1`, 1},
					{`1+2`, 3},
					{`"abc"`, "abc"},
				}

				for _, c := range cases {
					result, err := lua.Exec(emptyValueGetter, c.expr)
					So(err, ShouldBeNil)
					So(result, ShouldEqual, c.result)
				}
			})
		})

		Convey("when ValueGetter is not empty", func() {

			emptyValueGetter := parser.MapValueGetter{}.
				SetString("name", "eycia").
				SetBool("man", true).
				SetInt("age", 24).
				SetFloat("money", 0.1)
			So(emptyValueGetter, ShouldNotBeNil)

			Convey("LuaExpr should Exec() some expr correct", func() {
				cases := []struct {
					expr   string
					result interface{}
				}{
					{`1+age`, 25},
					{`name.." zhou"`, "eycia zhou"},
					{`money*10`, 1.0},
					{`man and name=="eycia"`, true},
					{`man and not name=="eycia"`, false},
				}

				for _, c := range cases {
					result, err := lua.Exec(emptyValueGetter, c.expr)
					So(err, ShouldBeNil)
					So(result, ShouldEqual, c.result)
				}
			})
		})
	})
}
