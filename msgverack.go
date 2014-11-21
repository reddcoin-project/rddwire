// Copyright (c) 2013-2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rddwire

import (
	"io"
)

// MsgVerAck defines a Reddcoin verack message which is used for a peer to
// acknowledge a version message (MsgVersion) after it has used the information
// to negotiate parameters.  It implements the Message interface.
//
// This message has no payload.
type MsgVerAck struct{}

// BtcDecode decodes r using the Reddcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgVerAck) BtcDecode(r io.Reader, pver uint32) error {
	return nil
}

// BtcEncode encodes the receiver to w using the Reddcoin protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgVerAck) BtcEncode(w io.Writer, pver uint32) error {
	return nil
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgVerAck) Command() string {
	return CmdVerAck
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgVerAck) MaxPayloadLength(pver uint32) uint32 {
	return 0
}

// NewMsgVerAck returns a new Reddcoin verack message that conforms to the
// Message interface.
func NewMsgVerAck() *MsgVerAck {
	return &MsgVerAck{}
}
