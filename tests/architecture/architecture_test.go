package architecture

import (
	"runtime"
	"testing"
)

func TestArchitecture(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Not working on amd64 architecture.")
	}

	t.Fail()
}
