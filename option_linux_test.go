// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tcpopt_test

import (
	"syscall"
	"testing"

	"github.com/mikioh/tcpopt"
)

func TestMarshalAndParseOnLinux(t *testing.T) {
	for _, o := range []tcpopt.Option{
		&tcpopt.OriginalDst{Family: syscall.AF_INET},
		&tcpopt.OriginalDst{Family: syscall.AF_INET6},
	} {
		b, err := o.Marshal()
		if err != nil {
			t.Fatal(err)
		}
		if _, err := tcpopt.Parse(o.Level(), o.Name(), b); err != nil {
			t.Fatal(err)
		}
	}
}
