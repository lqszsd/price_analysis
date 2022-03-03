package corns

import (
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {
	RegisterCrons()
	time.Sleep(time.Second * 100)
}
