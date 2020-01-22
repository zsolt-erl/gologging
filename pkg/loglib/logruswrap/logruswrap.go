package logruswrap

import (
	"fmt"

	"../../loglib"
	"github.com/sirupsen/logrus"
)

// Wrap implements a wrapper for logrus
type Wrap struct{}

// Info prints an info message including the key/value pairs
func (w Wrap) Info(KV loglib.Pairs, args ...interface{}) {
	for k, v := range KV {
		kvstr := fmt.Sprintf(" %v=%v", k, v)
		args = append(args, kvstr)
	}
	logrus.Info(args...)
}
