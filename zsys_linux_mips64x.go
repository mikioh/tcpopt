// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs defs_linux.go

// +build mips64 mips64le
// +build linux

package tcpopt

const (
	sysSOL_SOCKET = 0x1

	sysSO_KEEPALIVE         = 0x8
	sysSO_SNDBUF            = 0x1001
	sysSO_RCVBUF            = 0x1002
	sysSO_ERROR             = 0x1007
	sysSO_ORIGINAL_DST      = 0x50
	sysIP6T_SO_ORIGINAL_DST = 0x50

	sysTCP_NODELAY       = 0x1
	sysTCP_MAXSEG        = 0x2
	sysTCP_KEEPIDLE      = 0x4
	sysTCP_KEEPINTVL     = 0x5
	sysTCP_KEEPCNT       = 0x6
	sysTCP_CORK          = 0x3
	sysTCP_NOTSENT_LOWAT = 0x19
)

type sysSockaddrStorage struct {
	Family        uint16
	Pad_cgo_0     [6]byte
	X__ss_align   uint64
	X__ss_padding [112]int8
}

type sysSockaddr struct {
	Family uint16
	Data   [14]int8
}

type sysSockaddrInet struct {
	Family uint16
	Port   uint16
	Addr   [4]byte /* in_addr */
	X__pad [8]uint8
}

type sysSockaddrInet6 struct {
	Family   uint16
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte /* in6_addr */
	Scope_id uint32
}

const (
	sizeofSockaddrStorage = 0x80
	sizeofSockaddr        = 0x10
	sizeofSockaddrInet    = 0x10
	sizeofSockaddrInet6   = 0x1c
)
