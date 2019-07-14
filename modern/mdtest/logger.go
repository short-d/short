package mdtest

type logger struct{}

func (logger) Info(info string) {}

func (logger) Error(err error) {}

func (logger) Crash(err error) {}

var FakeLogger = logger{}
