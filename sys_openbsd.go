// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tcpopt

var options = map[int]option{
	noDelay:   {ianaProtocolTCP, sysTCP_NODELAY, 0},
	bSend:     {sysSOL_SOCKET, sysSO_SNDBUF, 0},
	bReceive:  {sysSOL_SOCKET, sysSO_RCVBUF, 0},
	keepAlive: {sysSOL_SOCKET, sysSO_KEEPALIVE, 0},
}

var parsers = map[int64]func([]byte) (Option, error){
	ianaProtocolTCP<<32 | sysTCP_NODELAY: parseNoDelay,
	sysSOL_SOCKET<<32 | sysSO_SNDBUF:     parseSendBuffer,
	sysSOL_SOCKET<<32 | sysSO_RCVBUF:     parseReceiveBuffer,
	sysSOL_SOCKET<<32 | sysSO_KEEPALIVE:  parseKeepAlive,
}
