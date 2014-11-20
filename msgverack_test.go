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

// TestVerAck tests the MsgVerAck API.
func TestVerAck(t *testing.T) {
	pver := rddwire.ProtocolVersion

	// Ensure the command is expected value.
	wantCmd := "verack"
	msg := rddwire.NewMsgVerAck()
	if cmd := msg.Command(); cmd != wantCmd {
		t.Errorf("NewMsgVerAck: wrong command - got %v want %v",
			cmd, wantCmd)
	}

	// Ensure max payload is expected value.
	wantPayload := uint32(0)
	maxPayload := msg.MaxPayloadLength(pver)
	if maxPayload != wantPayload {
		t.Errorf("MaxPayloadLength: wrong max payload length for "+
			"protocol version %d - got %v, want %v", pver,
			maxPayload, wantPayload)
	}

	return
}

// TestVerAckWire tests the MsgVerAck wire encode and decode for various
// protocol versions.
func TestVerAckWire(t *testing.T) {
	msgVerAck := rddwire.NewMsgVerAck()
	msgVerAckEncoded := []byte{}

	tests := []struct {
		in   *rddwire.MsgVerAck // Message to encode
		out  *rddwire.MsgVerAck // Expected decoded message
		buf  []byte             // Wire encoding
		pver uint32             // Protocol version for wire encoding
	}{
		// Latest protocol version.
		{
			msgVerAck,
			msgVerAck,
			msgVerAckEncoded,
			rddwire.ProtocolVersion,
		},

		// Protocol version BIP0035Version.
		{
			msgVerAck,
			msgVerAck,
			msgVerAckEncoded,
			rddwire.BIP0035Version,
		},

		// Protocol version BIP0031Version.
		{
			msgVerAck,
			msgVerAck,
			msgVerAckEncoded,
			rddwire.BIP0031Version,
		},

		// Protocol version NetAddressTimeVersion.
		{
			msgVerAck,
			msgVerAck,
			msgVerAckEncoded,
			rddwire.NetAddressTimeVersion,
		},

		// Protocol version MultipleAddressVersion.
		{
			msgVerAck,
			msgVerAck,
			msgVerAckEncoded,
			rddwire.MultipleAddressVersion,
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		// Encode the message to wire format.
		var buf bytes.Buffer
		err := test.in.BtcEncode(&buf, test.pver)
		if err != nil {
			t.Errorf("BtcEncode #%d error %v", i, err)
			continue
		}
		if !bytes.Equal(buf.Bytes(), test.buf) {
			t.Errorf("BtcEncode #%d\n got: %s want: %s", i,
				spew.Sdump(buf.Bytes()), spew.Sdump(test.buf))
			continue
		}

		// Decode the message from wire format.
		var msg rddwire.MsgVerAck
		rbuf := bytes.NewReader(test.buf)
		err = msg.BtcDecode(rbuf, test.pver)
		if err != nil {
			t.Errorf("BtcDecode #%d error %v", i, err)
			continue
		}
		if !reflect.DeepEqual(&msg, test.out) {
			t.Errorf("BtcDecode #%d\n got: %s want: %s", i,
				spew.Sdump(msg), spew.Sdump(test.out))
			continue
		}
	}
}
