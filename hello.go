package main

import (
	"./pkg/loglib"
	"./pkg/loglib/logruswrap"
)

func main() {
	//	l := loglib.Logger{Wrap: logruswrap.Wrap{}, KV: Pairs{}}
	l := loglib.Logger{KV: loglib.Pairs{"host": "mta.xt.local"}, Wrap: logruswrap.Wrap{}}

	l.Info("HELLO")
	l.With("process", "12").Info("HELLO ", "Zsolt")

	l2 := l.With("thread", "1666")
	l2.Info("HELLO ", "Zsolt")
}
