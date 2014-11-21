// Copyright (c) 2013-2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rddwire_test

import (
	"testing"

	"github.com/reddcoin-project/rddwire"
)

// TestServiceFlagStringer tests the stringized output for service flag types.
func TestServiceFlagStringer(t *testing.T) {
	tests := []struct {
		in   rddwire.ServiceFlag
		want string
	}{
		{0, "0x0"},
		{rddwire.SFNodeNetwork, "SFNodeNetwork"},
		{0xffffffff, "SFNodeNetwork|0xfffffffe"},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

// TestReddcoinNetStringer tests the stringized output for Reddcoin net types.
func TestReddcoinNetStringer(t *testing.T) {
	tests := []struct {
		in   rddwire.ReddcoinNet
		want string
	}{
		{rddwire.MainNet, "MainNet"},
		{rddwire.TestNet, "TestNet"},
		{rddwire.TestNet3, "TestNet3"},
		{rddwire.SimNet, "SimNet"},
		{0xffffffff, "Unknown ReddcoinNet (4294967295)"},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}
