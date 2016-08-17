// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !linux

package tcpopt

// Level implements the Level method of Option interface.
func (od *OriginalDst) Level() int { return 0 }

// Name implements the Name method of Option interface.
func (od *OriginalDst) Name() int { return 0 }

// Marshal implements the Marshal method of Option interface.
func (od *OriginalDst) Marshal() ([]byte, error) { return nil, errOpNoSupport }
