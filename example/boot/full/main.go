// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.
package main

import (
	"context"
	_ "embed"
	"github.com/rookie-ninja/rk-entry/v2/entry"
	"github.com/swordkee/rk-grpc/v2/boot"
)

//go:embed boot.yaml
var boot []byte

func main() {
	// Bootstrap basic entries from boot config.
	rkentry.BootstrapBuiltInEntryFromYAML(boot)
	rkentry.BootstrapPluginEntryFromYAML(boot)

	// Bootstrap grpc entry from boot config
	res := rkgrpc.RegisterGrpcEntryYAML(boot)

	// Bootstrap gin entry
	res["greeter"].Bootstrap(context.Background())

	// Wait for shutdown signal
	rkentry.GlobalAppCtx.WaitForShutdownSig()

	// Interrupt gin entry
	res["greeter"].Interrupt(context.Background())
}
