# external-ids
A library to manage and format externalIDs

### Purpose
There are a number of services at play that all have their own unique
identifiers for each concept. Rather than trying to amend our models with each
new service added, our microservices will instead have lists of "external IDs,"
that will be formatted in a specific way. The format will look like the
following:
```
system:type:12123
```

So for example, consider this external ID
```
io:order:123123
```

That identifies that `io` is the system of record the identifier originated
from, `order` is the concept being identified (as opposed to a `display` or
`network`), and the identifier itself is `123123`.

This library is intended to formalize which external identifiers are available,
and how they are to be formatted.

### Usage

Simply include the library and call the appropriate function.

```golang
package main

import (
	externalIDs "github.com/clearchanneloutdoor/external-ids"
)

func main() {
	orderID := 123123
	ioOrderID := externalIDs.FormatOrderId(orderID)
	// continue
}
```
