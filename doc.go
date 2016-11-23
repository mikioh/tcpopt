// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tcpopt implements encoding and decoding of TCP-level socket
// options.
//
// Example:
//
//	import (
//		"github.com/mikioh/tcp"
//		"github.com/mikioh/tcpopt"
//	)
//
//	tc, err := tcp.NewConn(c)
//	if err != nil {
//		// error handling
//	}
//	if err := tc.SetOption(tcpopt.KeepAlive(true)); err != nil {
//		// error handling
//	}
//	if err := tc.SetOption(tcpopt.KeepAliveProbeCount(3)); err != nil {
//		// error handling
//	}
package tcpopt
