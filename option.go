// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tcpopt implements encoding and decoding of TCP-level socket
// options.
package tcpopt

import "time"

// An Option represents a socket option.
type Option interface {
	// Level returns the platform-specific socket option level.
	Level() int

	// Name returns the platform-specific socket option name.
	Name() int

	// Marshal returns the binary encoding of socket option.
	Marshal() ([]byte, error)
}

// NoDealy specifies the use of Nagle's algorithm.
type NoDelay bool

// Level implements the Level method of Option interface.
func (nd NoDelay) Level() int { return options[noDelay].level }

// Name implements the Name method of Option interface.
func (nd NoDelay) Name() int { return options[noDelay].name }

// SendBuffer specifies the size of send buffer.
type SendBuffer int

// Level implements the Level method of Option interface.
func (sb SendBuffer) Level() int { return options[bSend].level }

// Name implements the Name method of Option interface.
func (sb SendBuffer) Name() int { return options[bSend].name }

// ReceiveBuffer specifies the size of send buffer.
type ReceiveBuffer int

// Level implements the Level method of Option interface.
func (rb ReceiveBuffer) Level() int { return options[bReceive].level }

// Name implements the Name method of Option interface.
func (rb ReceiveBuffer) Name() int { return options[bReceive].name }

// KeepAlive specifies the use of keep alive.
type KeepAlive bool

// Level implements the Level method of Option interface.
func (ka KeepAlive) Level() int { return options[keepAlive].level }

// Name implements the Name method of Option interface.
func (ka KeepAlive) Name() int { return options[keepAlive].name }

// KeepAliveIdleInterval is the idle interval until the first probe is
// sent.
//
// OpenBSD doesn't support this option.
// See TCP_KEEPIDLE or TCP_KEEPALIVE for further information.
type KeepAliveIdleInterval time.Duration

// Level implements the Level method of Option interface.
func (ka KeepAliveIdleInterval) Level() int { return options[kaIdleInterval].level }

// Name implements the Name method of Option interface.
func (ka KeepAliveIdleInterval) Name() int { return options[kaIdleInterval].name }

// KeepAliveProbeInterval is the interval between keepalive probes.
//
// OpenBSD doesn't support this option.
// See TCP_KEEPINTVL for further information.
type KeepAliveProbeInterval time.Duration

// Level implements the Level method of Option interface.
func (ka KeepAliveProbeInterval) Level() int { return options[kaProbeInterval].level }

// Name implements the Name method of Option interface.
func (ka KeepAliveProbeInterval) Name() int { return options[kaProbeInterval].name }

// KeepAliveProbeCount is the number of keepalive probes should be
// repeated when the peer is not responding.
//
// OpenBSD and Windows don't support this option.
// See TCP_KEEPCNT for further information.
type KeepAliveProbeCount int

// Level implements the Level method of Option interface.
func (ka KeepAliveProbeCount) Level() int { return options[kaProbeCount].level }

// Name implements the Name method of Option interface.
func (ka KeepAliveProbeCount) Name() int { return options[kaProbeCount].name }
