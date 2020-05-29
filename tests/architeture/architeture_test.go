package architeture

import (
	"runtime"
	"testing"
)

func TestArchiteture(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Not working on amd64 architeture.")
	}

	t.Fail()
}
