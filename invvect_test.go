// Copyright (c) 2013-2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rddwire_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/reddcoin-project/rddwire"
	"github.com/davecgh/go-spew/spew"
)

// TestInvVectStringer tests the stringized output for inventory vector types.
func TestInvTypeStringer(t *testing.T) {
	tests := []struct {
		in   rddwire.InvType
		want string
	}{
		{rddwire.InvTypeError, "ERROR"},
		{rddwire.InvTypeTx, "MSG_TX"},
		{rddwire.InvTypeBlock, "MSG_BLOCK"},
		{0xffffffff, "Unknown InvType (4294967295)"},
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

// TestInvVect tests the InvVect API.
func TestInvVect(t *testing.T) {
	ivType := rddwire.InvTypeBlock
	hash := rddwire.ShaHash{}

	// Ensure we get the same payload and signature back out.
	iv := rddwire.NewInvVect(ivType, &hash)
	if iv.Type != ivType {
		t.Errorf("NewInvVect: wrong type - got %v, want %v",
			iv.Type, ivType)
	}
	if !iv.Hash.IsEqual(&hash) {
		t.Errorf("NewInvVect: wrong hash - got %v, want %v",
			spew.Sdump(iv.Hash), spew.Sdump(hash))
	}

}

// TestInvVectWire tests the InvVect wire encode and decode for various
// protocol versions and supported inventory vector types.
func TestInvVectWire(t *testing.T) {
	// Block 203707 hash.
	hashStr := "3264bc2ac36a60840790ba1d475d01367e7c723da941069e9dc"
	baseHash, err := rddwire.NewShaHashFromStr(hashStr)
	if err != nil {
		t.Errorf("NewShaHashFromStr: %v", err)
	}

	// errInvVect is an inventory vector with an error.
	errInvVect := rddwire.InvVect{
		Type: rddwire.InvTypeError,
		Hash: rddwire.ShaHash{},
	}

	// errInvVectEncoded is the wire encoded bytes of errInvVect.
	errInvVectEncoded := []byte{
		0x00, 0x00, 0x00, 0x00, // InvTypeError
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // No hash
	}

	// txInvVect is an inventory vector representing a transaction.
	txInvVect := rddwire.InvVect{
		Type: rddwire.InvTypeTx,
		Hash: *baseHash,
	}

	// txInvVectEncoded is the wire encoded bytes of txInvVect.
	txInvVectEncoded := []byte{
		0x01, 0x00, 0x00, 0x00, // InvTypeTx
		0xdc, 0xe9, 0x69, 0x10, 0x94, 0xda, 0x23, 0xc7,
		0xe7, 0x67, 0x13, 0xd0, 0x75, 0xd4, 0xa1, 0x0b,
		0x79, 0x40, 0x08, 0xa6, 0x36, 0xac, 0xc2, 0x4b,
		0x26, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Block 203707 hash
	}

	// blockInvVect is an inventory vector representing a block.
	blockInvVect := rddwire.InvVect{
		Type: rddwire.InvTypeBlock,
		Hash: *baseHash,
	}

	// blockInvVectEncoded is the wire encoded bytes of blockInvVect.
	blockInvVectEncoded := []byte{
		0x02, 0x00, 0x00, 0x00, // InvTypeBlock
		0xdc, 0xe9, 0x69, 0x10, 0x94, 0xda, 0x23, 0xc7,
		0xe7, 0x67, 0x13, 0xd0, 0x75, 0xd4, 0xa1, 0x0b,
		0x79, 0x40, 0x08, 0xa6, 0x36, 0xac, 0xc2, 0x4b,
		0x26, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Block 203707 hash
	}

	tests := []struct {
		in   rddwire.InvVect // NetAddress to encode
		out  rddwire.InvVect // Expected decoded NetAddress
		buf  []byte          // Wire encoding
		pver uint32          // Protocol version for wire encoding
	}{
		// Latest protocol version error inventory vector.
		{
			errInvVect,
			errInvVect,
			errInvVectEncoded,
			rddwire.ProtocolVersion,
		},

		// Latest protocol version tx inventory vector.
		{
			txInvVect,
			txInvVect,
			txInvVectEncoded,
			rddwire.ProtocolVersion,
		},

		// Latest protocol version block inventory vector.
		{
			blockInvVect,
			blockInvVect,
			blockInvVectEncoded,
			rddwire.ProtocolVersion,
		},

		// Protocol version BIP0035Version error inventory vector.
		{
			errInvVect,
			errInvVect,
			errInvVectEncoded,
			rddwire.BIP0035Version,
		},

		// Protocol version BIP0035Version tx inventory vector.
		{
			txInvVect,
			txInvVect,
			txInvVectEncoded,
			rddwire.BIP0035Version,
		},

		// Protocol version BIP0035Version block inventory vector.
		{
			blockInvVect,
			blockInvVect,
			blockInvVectEncoded,
			rddwire.BIP0035Version,
		},

		// Protocol version BIP0031Version error inventory vector.
		{
			errInvVect,
			errInvVect,
			errInvVectEncoded,
			rddwire.BIP0031Version,
		},

		// Protocol version BIP0031Version tx inventory vector.
		{
			txInvVect,
			txInvVect,
			txInvVectEncoded,
			rddwire.BIP0031Version,
		},

		// Protocol version BIP0031Version block inventory vector.
		{
			blockInvVect,
			blockInvVect,
			blockInvVectEncoded,
			rddwire.BIP0031Version,
		},

		// Protocol version NetAddressTimeVersion error inventory vector.
		{
			errInvVect,
			errInvVect,
			errInvVectEncoded,
			rddwire.NetAddressTimeVersion,
		},

		// Protocol version NetAddressTimeVersion tx inventory vector.
		{
			txInvVect,
			txInvVect,
			txInvVectEncoded,
			rddwire.NetAddressTimeVersion,
		},

		// Protocol version NetAddressTimeVersion block inventory vector.
		{
			blockInvVect,
			blockInvVect,
			blockInvVectEncoded,
			rddwire.NetAddressTimeVersion,
		},

		// Protocol version MultipleAddressVersion error inventory vector.
		{
			errInvVect,
			errInvVect,
			errInvVectEncoded,
			rddwire.MultipleAddressVersion,
		},

		// Protocol version MultipleAddressVersion tx inventory vector.
		{
			txInvVect,
			txInvVect,
			txInvVectEncoded,
			rddwire.MultipleAddressVersion,
		},

		// Protocol version MultipleAddressVersion block inventory vector.
		{
			blockInvVect,
			blockInvVect,
			blockInvVectEncoded,
			rddwire.MultipleAddressVersion,
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		// Encode to wire format.
		var buf bytes.Buffer
		err := rddwire.TstWriteInvVect(&buf, test.pver, &test.in)
		if err != nil {
			t.Errorf("writeInvVect #%d error %v", i, err)
			continue
		}
		if !bytes.Equal(buf.Bytes(), test.buf) {
			t.Errorf("writeInvVect #%d\n got: %s want: %s", i,
				spew.Sdump(buf.Bytes()), spew.Sdump(test.buf))
			continue
		}

		// Decode the message from wire format.
		var iv rddwire.InvVect
		rbuf := bytes.NewReader(test.buf)
		err = rddwire.TstReadInvVect(rbuf, test.pver, &iv)
		if err != nil {
			t.Errorf("readInvVect #%d error %v", i, err)
			continue
		}
		if !reflect.DeepEqual(iv, test.out) {
			t.Errorf("readInvVect #%d\n got: %s want: %s", i,
				spew.Sdump(iv), spew.Sdump(test.out))
			continue
		}
	}
}
