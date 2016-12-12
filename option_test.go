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

func TestMarshalAndParse(t *testing.T) {
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
	switch runtime.GOOS {
	case "windows":
	default:
		opts = append(opts, tcpopt.MSS(4092))
		opts = append(opts, tcpopt.Error(42))
	}
	switch runtime.GOOS {
	case "darwin":
		opts = append(opts, tcpopt.ECN(true))
	}

	for _, o := range opts {
		if o.Level() <= 0 {
			t.Fatalf("got %#x; want greater than zero", o.Level())
		}
		if o.Name() <= 0 {
			t.Fatalf("got %#x; want greater than zero", o.Name())
		}
		b, err := o.Marshal()
		if err != nil {
			t.Fatal(err)
		}
		if runtime.GOOS == "windows" {
			continue
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
	testOptLevel = 0xfff1
	testOptName  = 0xfff2
)

type testOption struct{}

func (*testOption) Level() int                        { return testOptLevel }
func (*testOption) Name() int                         { return testOptName }
func (*testOption) Marshal() ([]byte, error)          { return make([]byte, 16), nil }
func parseTestOption(_ []byte) (tcpopt.Option, error) { return &testOption{}, nil }

func TestUserDefinedOptionParser(t *testing.T) {
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

func TestParseWithVariousBufferLengths(t *testing.T) {
	for _, o := range []tcpopt.Option{
		tcpopt.NoDelay(true),
		tcpopt.SendBuffer(1<<16 - 1),
		tcpopt.ReceiveBuffer(1<<16 - 1),
		tcpopt.KeepAlive(true),
		tcpopt.KeepAliveIdleInterval(1 * time.Hour),
		tcpopt.KeepAliveProbeInterval(10 * time.Minute),
		tcpopt.KeepAliveProbeCount(3),
		tcpopt.Error(42),
	} {
		for i := 0; i < 256; i++ {
			b := make([]byte, i)
			if _, err := tcpopt.Parse(o.Level(), o.Name(), b); err == nil {
				break
			}
		}
	}
}
