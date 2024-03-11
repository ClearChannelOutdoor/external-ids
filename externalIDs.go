package externalIDs

import (
	"fmt"
	"strconv"
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

	quattroDbPrefix              = "quattro_"
	quattroCampaignPrefix        = ":campaign:"
	quattroCampaignSegmentPrefix = ":campaignSegment:"
	quattroCampaignDetailPrefix  = ":campaignDetail:"
	quattroBookingPrefix         = ":booking:"
	quattroDigitalBookingPrefix  = ":digitalBooking:"
	quattroDisplayID             = ":display:"
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
	return formatQuattroKey(marketCode, quattroCampaignPrefix, fmt.Sprint(campaignId))
}

func FormatQuattroCampaignSegment(marketCode string, segmentId interface{}) string {
	return formatQuattroKey(marketCode, quattroCampaignSegmentPrefix, fmt.Sprint(segmentId))
}

func FormatQuattroCampaignDetail(marketCode string, detailId interface{}) string {
	return formatQuattroKey(marketCode, quattroCampaignDetailPrefix, fmt.Sprint(detailId))
}

func FormatQuattroBookingID(sourceDbCode string, bookingID interface{}) string {
	return formatQuattroKey(sourceDbCode, quattroBookingPrefix, fmt.Sprint(bookingID))
}

func FormatQuattroDigitalBookingID(sourceDbCode string, bookingID interface{}) string {
	return formatQuattroKey(sourceDbCode, quattroDigitalBookingPrefix, fmt.Sprint(bookingID))
}

func FormatQuattroDisplayID(sourceDbCode string, panelID interface{}) string {
	return formatQuattroKey(sourceDbCode, quattroDisplayID, fmt.Sprint(panelID))
}

func GetQuattroBookingID(sourceDbCode string, externalIDs []string) *int {
	prefix := formatQuattroKey(sourceDbCode, quattroBookingPrefix, "")
	for _, externalID := range externalIDs {
		if !strings.HasPrefix(externalID, prefix) {
			// skip this iteration
			continue
		}

		externalIDParts := strings.Split(externalID, ":")
		intID, err := strconv.Atoi(externalIDParts[len(externalIDParts)-1])
		if err != nil {
			continue
		}

		return &intID
	}
	return nil
}

func GetLegacySiteCode(externalIDs []string) *int {
	return getLegacyIOEntityNumericID(externalIDs, "site")
}

func GetLegacyIODisplayID(externalIDs []string) *int {
	return getLegacyIOEntityNumericID(externalIDs, "display")
}

func getLegacyIOEntityNumericID(externalIDs []string, entity string) *int {
	for _, externalID := range externalIDs {
		// if externalID does not begin with io:{entity}
		if !strings.HasPrefix(externalID, fmt.Sprintf("io:%s:", entity)) {
			// skip this iteration
			continue
		}

		//split the externalID into a 3-part array on the : character
		externalIDParts := strings.Split(externalID, ":")
		if len(externalIDParts) != 3 {
			continue
		}

		// if entity part doesn't match requested type
		if externalIDParts[1] != entity {
			continue
		}

		// convert externalIDParts[2] to int
		intID, err := strconv.Atoi(externalIDParts[2])
		if err != nil {
			continue
		}

		return &intID
	}
	return nil
}

func formatQuattroKey(marketCode string, entity string, id string) string {
	quattroKey := format(quattroDbPrefix, marketCode)
	prefix := format(quattroKey, entity)
	return format(prefix, id)
}

func format(prefix string, identifier string) string {
	if strings.HasPrefix(identifier, prefix) {
		return identifier
	}
	return fmt.Sprintf("%s%s", prefix, identifier)
}
