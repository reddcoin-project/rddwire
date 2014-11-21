rddwire
=======

[![Build Status](https://travis-ci.org/reddcoin-project/rddwire.png?branch=master)]
(https://travis-ci.org/reddcoin-project/rddwire) [![Coverage Status]
(https://coveralls.io/repos/reddcoin-project/rddwire/badge.png?branch=master)]
(https://coveralls.io/r/reddcoin-project/rddwire?branch=master)

Package rddwire implements the Reddcoin wire protocol.  A comprehensive suite of
tests with 100% test coverage is provided to ensure proper functionality.
Package rddwire is licensed under the liberal ISC license.

There is an associated blog post about the release of this package
[here](https://blog.conformal.com/btcwire-the-bitcoin-wire-protocol-package-from-btcd/).

This package is one of the core packages from rddd, an alternative full-node
implementation of Reddcoin which is under active development by Conformal.
Although it was primarily written for rddd, this package has intentionally been
designed so it can be used as a standalone package for any projects needing to
interface with Reddcoin peers at the wire protocol level.

## Documentation

[![GoDoc](https://godoc.org/github.com/reddcoin-project/rddwire?status.png)]
(http://godoc.org/github.com/reddcoin-project/rddwire)

Full `go doc` style documentation for the project can be viewed online without
installing this package by using the GoDoc site here:
http://godoc.org/github.com/reddcoin-project/rddwire

You can also view the documentation locally once the package is installed with
the `godoc` tool by running `godoc -http=":6060"` and pointing your browser to
http://localhost:6060/pkg/github.com/reddcoin-project/rddwire

## Installation

```bash
$ go get github.com/reddcoin-project/rddwire
```

## Reddcoin Message Overview

The Reddcoin protocol consists of exchanging messages between peers. Each message
is preceded by a header which identifies information about it such as which
Reddcoin network it is a part of, its type, how big it is, and a checksum to
verify validity. All encoding and decoding of message headers is handled by this
package.

To accomplish this, there is a generic interface for Reddcoin messages named
`Message` which allows messages of any type to be read, written, or passed
around through channels, functions, etc. In addition, concrete implementations
of most of the currently supported Reddcoin messages are provided. For these
supported messages, all of the details of marshalling and unmarshalling to and
from the wire using Reddcoin encoding are handled so the caller doesn't have to
concern themselves with the specifics.

## Reading Messages Example

In order to unmarshal Reddcoin messages from the wire, use the `ReadMessage`
function. It accepts any `io.Reader`, but typically this will be a `net.Conn`
to a remote node running a Reddcoin peer.  Example syntax is:

```Go
	// Use the most recent protocol version supported by the package and the
	// main Reddcoin network.
	pver := rddwire.ProtocolVersion
	rddnet := rddwire.MainNet

	// Reads and validates the next Reddcoin message from conn using the
	// protocol version pver and the Reddcoin network rddnet.  The returns
	// are a rddwire.Message, a []byte which contains the unmarshalled
	// raw payload, and a possible error.
	msg, rawPayload, err := rddwire.ReadMessage(conn, pver, rddnet)
	if err != nil {
		// Log and handle the error
	}
```

See the package documentation for details on determining the message type.

## Writing Messages Example

In order to marshal Reddcoin messages to the wire, use the `WriteMessage`
function. It accepts any `io.Writer`, but typically this will be a `net.Conn`
to a remote node running a Reddcoin peer. Example syntax to request addresses
from a remote peer is:

```Go
	// Use the most recent protocol version supported by the package and the
	// main Reddcoin network.
	pver := rddwire.ProtocolVersion
	rddnet := rddwire.MainNet

	// Create a new getaddr Reddcoin message.
	msg := rddwire.NewMsgGetAddr()

	// Writes a Reddcoin message msg to conn using the protocol version
	// pver, and the Reddcoin network rddnet.  The return is a possible
	// error.
	err := rddwire.WriteMessage(conn, msg, pver, rddnet)
	if err != nil {
		// Log and handle the error
	}
```

## GPG Verification Key

All official release tags are signed by Conformal so users can ensure the code
has not been tampered with and is coming from Conformal.  To verify the
signature perform the following:

- Download the public key from the Conformal website at
  https://opensource.conformal.com/GIT-GPG-KEY-conformal.txt

- Import the public key into your GPG keyring:
  ```bash
  gpg --import GIT-GPG-KEY-conformal.txt
  ```

- Verify the release tag with the following command where `TAG_NAME` is a
  placeholder for the specific tag:
  ```bash
  git tag -v TAG_NAME
  ```

## License

Package rddwire is licensed under the [copyfree](http://copyfree.org) ISC
License.
