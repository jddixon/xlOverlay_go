<h1 class="libTop">xlOverlay_go</h1>

The Go language Overlay functions for 
[XLattice.](https://jddixon.github.io/xlattice_go)

An Overlay is characterized by an address space, a transport protocol,
and possibly a set of rules for navigating the address space using
the protocol.

An Overlay may either be system-supported, like TCP/IP will
normally be, or it may explicitly depend upon an underlying
Overlay, in the way that HTTP, for example, is generally
implemented over TCP/IP.

If the Overlay is system-supported, traffic will be routed and
neighbors will be reached by making calls to operating system
primitives such as sockets.

In some Overlays there is a method which, given an EndPoint `x`, returns
another EndPoint `g`, a **gateway**, which can be used to route messages to
EndPoint `x`.

## Examples of Overlays

* the global Internet, `0/0`
* private IP address space, usually in `10/8` or `/16`s in  `192.168`
* **email-mediated**, where gateways convert messages into email and then use conventional email protocols (`SMTP`) to transport and store messages
* **data-keyed**, where machines exchange traffic through a content-keyed storage system
* **name-keyed**, where systems exchange traffic through a distributed system storing values by name (the key)

## Gateways

A Gateway is an 
[XLattice Node](https://jddixon.github.io/xlNode_go)
which can be used for routing traffic between Overlays.  The Gateway
can simply forward message with headers added or altered as necessary 
to route traffic, or the message content may be altered as necessitated
for example by protocol switching or encryption.

## Project Status

Minimal implementation.  Functions are largely stubbed.

