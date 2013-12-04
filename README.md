microbe
=======

[![Build Status](https://travis-ci.org/jpgneves/microbe.png)](https://travis-ci.org/jpgneves/microbe)

Yet another micro web-framework written in Go.

Microbe provides a pluggable URL routing engine and wrappers for HTTP request
and response handling built on top of Go's net/http handlers.

Routing
-------

Microbe provides a pluggable routing engine that currently provides two
basic implementations.

StaticRouter provides plain, no-frills "exact URL"-based matching, meaning
that there are no dynamic components in the URL.

MatchingRouter provides a trie-based URL routing mechanism that can take
parameters in the path:

'/foo/:param1/bar'

Requests
--------

Microbe requests wrap the Go-provided net/http objects and add functionality
to extract all parsed parameters (if the MatchingRouter was used).

The raw net/http Request and Response objects are still accessible plainly
from the corresponding Microbe objects.

Resources
---------

Data structures implementing the Microbe Resource interface can be routed to
using the Microbe router and their GET/POST (more coming, don't worry) methods
will be executed.
