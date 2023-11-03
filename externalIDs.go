package externalIDs

import (
	"fmt"
	"strings"
)

const (
	ioOrderIDPrefix = "io:order:"
)

func FormatIOOrderID(orderID interface{}) string {
	var parsedOrderID string
	value, isString := orderID.(string)
	if !isString {
		parsedOrderID = fmt.Sprintf("%v", orderID)
	} else {
		parsedOrderID = value
	}
	return format(ioOrderIDPrefix, parsedOrderID)
}

func format(prefix string, identifier string) string {
	if strings.HasPrefix(identifier, prefix) {
		return identifier
	}
	return fmt.Sprintf("%s%s", prefix, identifier)
}
