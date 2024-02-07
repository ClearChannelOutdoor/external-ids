package externalIDs

import (
	"fmt"
	"strings"
)

const (
	extAgencyCodePrefi         = "agency:code:"
	geopathSegmentCodePrefix   = "geopath:segmentCode:"
	geopathTargetProfilePrefix = "geopath:targetProfile:"
	ioAccountIdPrefix          = "io:account:"
	ioBookingIDPrefix          = "io:booking:"
	ioCustomerPrefix           = "io:customer:"
	ioEmployeeeNumberPrefix    = "io:employee:"
	ioMarketPrefix             = "io:market:"
	ioOrderIDPrefix            = "io:order:"
	ioOrderNumberPrefix        = "io:orderNumber:"
	ioOrderLinePrefix          = "io:orderLine:"
	ioOrderLineTypePrefix      = "io:orderLineType:"
	ioOrderMarketIDPrefix      = "io:orderMarket:"

	quattroDbPrefix       = "quattro_:"
	quattroCampaignPrefix = "campaign:"
	quattroPanelID        = ":display:"
)

func FormatIOAccount(id interface{}) string {
	return format(ioAccountIdPrefix, fmt.Sprint(id))
}

func FormatIOCustomer(id interface{}) string {
	return format(ioCustomerPrefix, fmt.Sprint(id))
}

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

func FormatIOOrderNumber(number interface{}) string {
	return format(ioOrderNumberPrefix, fmt.Sprint(number))
}

func FormatIOOrderLine(id interface{}) string {
	return format(ioOrderLinePrefix, fmt.Sprint(id))
}

func FormatIOOrderLineType(id interface{}) string {
	return format(ioOrderLineTypePrefix, fmt.Sprint(id))
}

func FormatIOOrderMarket(id interface{}) string {
	return format(ioOrderMarketIDPrefix, fmt.Sprint(id))
}

func FormatIOBookingID(bookingID interface{}) string {
	return format(ioBookingIDPrefix, fmt.Sprint(bookingID))
}

func FormatIOEmployeeNumber(number interface{}) string {
	return format(ioEmployeeeNumberPrefix, fmt.Sprint(number))
}

func FormatIOMarket(id interface{}) string {
	return format(ioMarketPrefix, fmt.Sprint(id))
}

func FormatQuattroCampaign(marketCode string, campaignId interface{}) string {
	key := formatQuattroKey(quattroCampaignPrefix, marketCode)
	return format(key, fmt.Sprint(campaignId))
}

func FormatQuattroPanelID(sourceDB string, panelID interface{}) string {
	key := fmt.Sprintf("%s%s", sourceDB, quattroPanelID)
	return format(key, fmt.Sprint(panelID))
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
