// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package tcpopt

// +godefs map struct_in_addr [4]byte /* in_addr */
// +godefs map struct_in6_addr [16]byte /* in6_addr */

/*
#include <sys/ioctl.h>
#include <sys/socket.h>

#include <linux/if.h>
#include <linux/in.h>
#include <linux/in6.h>
#include <linux/netfilter_ipv4.h>
#include <linux/netfilter_ipv6/ip6_tables.h>
#include <linux/tcp.h>
*/
import "C"

const (
	sysSOL_SOCKET = C.SOL_SOCKET

	sysSO_KEEPALIVE         = C.SO_KEEPALIVE
	sysSO_SNDBUF            = C.SO_SNDBUF
	sysSO_RCVBUF            = C.SO_RCVBUF
	sysSO_ERROR             = C.SO_ERROR
	sysSO_ORIGINAL_DST      = C.SO_ORIGINAL_DST
	sysIP6T_SO_ORIGINAL_DST = C.IP6T_SO_ORIGINAL_DST

	sysTCP_NODELAY       = C.TCP_NODELAY
	sysTCP_MAXSEG        = C.TCP_MAXSEG
	sysTCP_KEEPIDLE      = C.TCP_KEEPIDLE
	sysTCP_KEEPINTVL     = C.TCP_KEEPINTVL
	sysTCP_KEEPCNT       = C.TCP_KEEPCNT
	sysTCP_CORK          = C.TCP_CORK
	sysTCP_NOTSENT_LOWAT = C.TCP_NOTSENT_LOWAT
)

type sysSockaddrStorage C.struct_sockaddr_storage

type sysSockaddr C.struct_sockaddr

type sysSockaddrInet C.struct_sockaddr_in

type sysSockaddrInet6 C.struct_sockaddr_in6

const (
	sizeofSockaddrStorage = C.sizeof_struct_sockaddr_storage
	sizeofSockaddr        = C.sizeof_struct_sockaddr
	sizeofSockaddrInet    = C.sizeof_struct_sockaddr_in
	sizeofSockaddrInet6   = C.sizeof_struct_sockaddr_in6
)
