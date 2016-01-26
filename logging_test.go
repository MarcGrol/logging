package logging

import (
	"testing"

	"golang.org/x/net/context"
)

func TestSimple(t *testing.T) {
	c := context.TODO()
	log := New()
	log.Critical(c, "crit %s", "my")
	log.Error(c, "err %s", "my")
	log.Warning(c, "warn %s", "my")
	log.Info(c, "info %s", "my")
	log.Debug(c, "deb %s", "my")
}
