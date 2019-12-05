module github.com/eycia/rudita

require (
	github.com/eycia/utils v0.0.1
	github.com/gopherjs/gopherjs v0.0.0-20190430165422-3e4dfb77656c // indirect
	github.com/jtolds/gls v4.2.1+incompatible // indirect
	github.com/sirupsen/logrus v1.2.0
	github.com/smartystreets/assertions v0.0.0-20190116191733-b6c0e53d7304 // indirect
	github.com/smartystreets/goconvey v0.0.0-20181108003508-044398e4856c
	github.com/yuin/gopher-lua v0.0.0-20181231133414-1e6e6e1918e0
)

replace (
	golang.org/x/crypto v0.0.0-20180904163835-0709b304e793 => github.com/golang/crypto v0.0.0-20180904163835-0709b304e793
	golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33 => github.com/golang/sys v0.0.0-20180905080454-ebe1bf3edb33
)

replace github.com/eycia/utils v0.0.1 => ../utils

go 1.13
