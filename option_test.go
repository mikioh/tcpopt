// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tcpopt_test

import (
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/mikioh/tcpopt"
)

func TestOption(t *testing.T) {
	opts := []tcpopt.Option{
		tcpopt.NoDelay(true),
		tcpopt.SendBuffer(1<<16 - 1),
		tcpopt.ReceiveBuffer(1<<16 - 1),
		tcpopt.KeepAlive(true),
	}
	switch runtime.GOOS {
	case "openbsd":
	case "windows":
		opts = append(opts, tcpopt.KeepAliveIdleInterval(1*time.Hour))
		opts = append(opts, tcpopt.KeepAliveProbeInterval(10*time.Minute))
	default:
		opts = append(opts, tcpopt.KeepAliveIdleInterval(1*time.Hour))
		opts = append(opts, tcpopt.KeepAliveProbeInterval(10*time.Minute))
		opts = append(opts, tcpopt.KeepAliveProbeCount(3))
	}
	switch runtime.GOOS {
	case "netbsd", "windows":
	default:
		opts = append(opts, tcpopt.Cork(true))
	}
	switch runtime.GOOS {
	case "darwin", "linux":
		opts = append(opts, tcpopt.NotSentLowWMK(1))
	}
	for _, o := range opts {
		if o.Level() == 0 {
			t.Fatalf("got %#x; want non-zero value", o.Level())
		}
		if o.Name() == 0 {
			t.Fatalf("got %#x; want non-zero value", o.Name())
		}
		b, err := o.Marshal()
		if err != nil {
			t.Fatal(err)
		}
		oo, err := tcpopt.Parse(o.Level(), o.Name(), b)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(oo, o) {
			t.Fatalf("got %#v; want %#v", oo, o)
		}
	}
}

const (
	testOptLevel = 0xfffe
	testOptName  = 0xfffd
)

type testOption struct{}

func (*testOption) Level() int                        { return testOptLevel }
func (*testOption) Name() int                         { return testOptName }
func (*testOption) Marshal() ([]byte, error)          { return make([]byte, 16), nil }
func parseTestOption(_ []byte) (tcpopt.Option, error) { return &testOption{}, nil }

func TestParse(t *testing.T) {
	var b [16]byte
	tcpopt.Register(testOptLevel, testOptName, parseTestOption)
	o, err := tcpopt.Parse(testOptLevel, testOptName, b[:])
	if err != nil {
		t.Fatal(err)
	}
	tcpopt.Unregister(testOptLevel, testOptName)
	o, err = tcpopt.Parse(testOptLevel, testOptName, b[:])
	if err == nil || o != nil {
		t.Fatalf("got %v, %v; want nil, error", o, err)
	}
}
