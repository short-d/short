package mdtest

import "short/fw"

type logger struct{}

func (logger) Info(info string) {}

func (logger) Error(err error) {}

func (logger) Crash(err error) {}

var FakeLogger fw.Logger = logger{}
