package externalIDs

import (
	"fmt"
	"strings"
)

const (
	ioBookingIDPrefix = "io:booking:"
	ioOrderIDPrefix   = "io:order:"
)

func FormatIOOrderID(orderID interface{}) string {
	return format(ioOrderIDPrefix, fmt.Sprint(orderID))
}

func FormatIOBookingID(bookingID interface{}) string {
	return format(ioBookingIDPrefix, fmt.Sprint(bookingID))
}

func format(prefix string, identifier string) string {
	if strings.HasPrefix(identifier, prefix) {
		return identifier
	}
	return fmt.Sprintf("%s%s", prefix, identifier)
}
