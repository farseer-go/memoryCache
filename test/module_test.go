package test

import (
	"github.com/farseer-go/cacheMemory"
	"testing"
)

func TestModule_Shutdown(t *testing.T) {
	cacheMemory.Module{}.Shutdown()
}
