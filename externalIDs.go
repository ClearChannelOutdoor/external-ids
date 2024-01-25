package externalIDs

import (
	"fmt"
	"strings"
)

const (
	extAgencyCodePrefi         = "agency:code:"
	geopathSegmentCodePrefix   = "geopath:segmentCode:"
	geopathTargetProfilePrefix = "geopath:targetProfile:"
	ioBookingIDPrefix          = "io:booking:"
	ioEmployeeeNumberPrefix    = "io:employee:"
	ioOrderIDPrefix            = "io:order:"
	ioOrderMarketIDPrefix      = "io:orderMarket:"

	quattroDbPrefix       = "quattro_:"
	quattroCampaignPrefix = "campaign:"
)

func FormatExtAgencyCode(code interface{}) string {
	return format(extAgencyCodePrefi, fmt.Sprint(code))
}
func FormatGeopathSegmentCode(code interface{}) string {
	return format(geopathSegmentCodePrefix, fmt.Sprint(code))
}

func FormatGeopathTargetProfile(code interface{}) string {
	return format(geopathTargetProfilePrefix, fmt.Sprint(code))
}

func FormatIOOrderID(orderID interface{}) string {
	return format(ioOrderIDPrefix, fmt.Sprint(orderID))
}

func FormatIOBookingID(bookingID interface{}) string {
	return format(ioBookingIDPrefix, fmt.Sprint(bookingID))
}

func FormatIOEmployeeNumber(number interface{}) string {
	return format(ioEmployeeeNumberPrefix, fmt.Sprint(number))
}

func FormatQuattroCampaign(marketCode string, campaignId interface{}) string {
	key := formatQuattroKey(quattroCampaignPrefix, marketCode)
	return format(key, fmt.Sprint(campaignId))
}

func formatQuattroKey(prefix string, marketCode string) string {
	return format(quattroDbPrefix, marketCode)
}

func format(prefix string, identifier string) string {
	if strings.HasPrefix(identifier, prefix) {
		return identifier
	}
	return fmt.Sprintf("%s%s", prefix, identifier)
}
