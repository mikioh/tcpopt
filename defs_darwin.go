// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package tcpopt

/*
#include <sys/socket.h>

#include <netinet/tcp.h>
*/
import "C"

const (
	sysSOL_SOCKET = C.SOL_SOCKET

	sysSO_KEEPALIVE = C.SO_KEEPALIVE
	sysSO_SNDBUF    = C.SO_SNDBUF
	sysSO_RCVBUF    = C.SO_RCVBUF
	sysSO_ERROR     = C.SO_ERROR

	sysTCP_NODELAY       = C.TCP_NODELAY
	sysTCP_MAXSEG        = C.TCP_MAXSEG
	sysTCP_KEEPALIVE     = C.TCP_KEEPALIVE
	sysTCP_KEEPINTVL     = C.TCP_KEEPINTVL
	sysTCP_KEEPCNT       = C.TCP_KEEPCNT
	sysTCP_NOPUSH        = C.TCP_NOPUSH
	sysTCP_ENABLE_ECN    = C.TCP_ENABLE_ECN
	sysTCP_NOTSENT_LOWAT = C.TCP_NOTSENT_LOWAT
)
