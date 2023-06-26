package main

import (
	"douyinAPI/internal"
	"github.com/pengcainiao/zero/core/sysx"
	"github.com/pengcainiao/zero/rest/httprouter"
)

func main() {
	sysx.SubSystem = "dy"
	httprouter.RecoveryMainAsync(func() {
		internal.Setup()
	})
}
