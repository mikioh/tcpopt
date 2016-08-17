// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tcpopt

import (
	"encoding/binary"
	"net"
	"syscall"
	"unsafe"
)

// Level implements the Level method of Option interface.
func (od *OriginalDst) Level() int {
	if od == nil {
		return 0
	}
	switch od.Family {
	case syscall.AF_INET:
		return ianaProtocolIP
	case syscall.AF_INET6:
		return ianaProtocolIPv6
	default:
		return 0
	}
}

// Name implements the Name method of Option interface.
func (od *OriginalDst) Name() int {
	if od == nil {
		return 0
	}
	switch od.Family {
	case syscall.AF_INET:
		return sysSO_ORIGINAL_DST
	case syscall.AF_INET6:
		return sysIP6T_SO_ORIGINAL_DST
	default:
		return 0
	}
}

// Marshal implements the Marshal method of tcpopt.Option interface.
func (od *OriginalDst) Marshal() ([]byte, error) {
	if od == nil {
		return nil, errInvalidOption
	}
	switch od.Family {
	case syscall.AF_INET:
		return make([]byte, sizeofSockaddrInet), nil
	case syscall.AF_INET6:
		return make([]byte, sizeofSockaddrInet6), nil
	default:
		return nil, &net.AddrError{Err: "invalid address family", Addr: od.IP.String()}
	}
}

func parseOriginalDst(b []byte) (Option, error) {
	switch len(b) {
	case sizeofSockaddrInet:
		od := new(OriginalDst)
		sa := (*sysSockaddrInet)(unsafe.Pointer(&b[0]))
		od.Family = syscall.AF_INET
		od.IP = net.IPv4(sa.Addr[0], sa.Addr[1], sa.Addr[2], sa.Addr[3])
		binary.BigEndian.PutUint16((*[2]byte)(unsafe.Pointer(&od.Port))[:], uint16(sa.Port))
		return od, nil
	case sizeofSockaddrInet6:
		od := new(OriginalDst)
		sa := (*sysSockaddrInet6)(unsafe.Pointer(&b[0]))
		od.Family = syscall.AF_INET6
		od.IP = make(net.IP, net.IPv6len)
		copy(od.IP, sa.Addr[:])
		binary.BigEndian.PutUint16((*[2]byte)(unsafe.Pointer(&od.Port))[:], uint16(sa.Port))
		od.ZoneID = int(sa.Scope_id)
		return od, nil
	default:
		return nil, errInvalidOption
	}
}
