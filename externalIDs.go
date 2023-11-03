package externalIDs

import (
	"fmt"
	"strings"
)

const (
	ioOrderIDPrefix = "io:order:"
)

func FormatIOOrderID(orderID interface{}) string {
	return format(ioOrderIDPrefix, fmt.Sprint(orderID))
}

func format(prefix string, identifier string) string {
	if strings.HasPrefix(identifier, prefix) {
		return identifier
	}
	return fmt.Sprintf("%s%s", prefix, identifier)
}
