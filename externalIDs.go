package externalIDs

import (
	"fmt"
	"strings"
)

const (
	extCustomerOrderPrefix         = "customer:order:"
	geopathSegmentCodePrefix       = "geopath:segmentCode:"
	geopathTargetProfilePrefix     = "geopath:targetProfile:"
	ioAccountIdPrefix              = "io:account:"
	ioBookingIDPrefix              = "io:booking:"
	ioCreativePrefix               = "io:creative:"
	ioCustomerPrefix               = "io:customer:"
	ioDocumentPrefix               = "io:document:"
	ioDisplayPrefix                = "io:display:"
	ioEmployeePrefix               = "io:employee:"
	ioEmployeeNumberPrefix         = "io:employeeNumber:"
	ioGroupBookingPrefix           = "io:groupBooking:"
	ioOrderLineTypePrefix          = "io:lineType:"
	ioMarketPrefix                 = "io:market:"
	ioNetworkPrefix                = "io:network:"
	ioNetworkCodePrefix            = "io:networkCode:"
	ioOrderIDPrefix                = "io:order:"
	ioOrderLinePrefix              = "io:orderLine:"
	ioOrderMarketIDPrefix          = "io:orderMarket:"
	ioOrderNumberPrefix            = "io:orderNumber:"
	ioProductMapIDPrefix           = "io:productMap:"
	quattroDbPrefix                = "quattro_"
	quattroBookingPrefix           = ":booking:"
	quattroCampaignPrefix          = ":campaign:"
	quattroCampaignDetailPrefix    = ":campaignDetail:"
	quattroCampaignSegmentPrefix   = ":campaignSegment:"
	quattroDigitalBookingPrefix    = ":digitalbooking:"
	oldQuattroDigitalBookingPrefix = ":digitalBooking:"
	quattroDisplayID               = ":display:"
	quattroNetworkID               = ":network:"
	spotChartDisplayPrefix         = "spotchart:display:"
	spotChartScenePrefix           = "spotchart:scene:"
	spotChartSegmentPrefix         = "spotchart:segment:"
)

/* FORMATTERS */

func FormatIOAccount(id interface{}) string {
	return format(ioAccountIdPrefix, fmt.Sprint(id))
}

func FormatIOCreative(id interface{}) string {
	return format(ioCreativePrefix, fmt.Sprint(id))
}

func FormatIOCustomer(id interface{}) string {
	return format(ioCustomerPrefix, fmt.Sprint(id))
}

func FormatCustomerOrder(code interface{}) string {
	return format(extCustomerOrderPrefix, fmt.Sprint(code))
}

func FormatIODocument(id interface{}) string {
	return format(ioDocumentPrefix, fmt.Sprint(id))
}

func FormatSpotChartDisplayID(displayID interface{}) string {
	return format(spotChartDisplayPrefix, fmt.Sprint(displayID))
}

func FormatSpotChartSceneID(sceneID interface{}) string {
	return format(spotChartScenePrefix, fmt.Sprint(sceneID))
}

func FormatSpotChartSegmentID(segmentID interface{}) string {
	return format(spotChartSegmentPrefix, fmt.Sprint(segmentID))
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

func FormatIOProductMap(id interface{}) string {
	return format(ioProductMapIDPrefix, fmt.Sprint(id))
}

func FormatIOBookingID(bookingID interface{}) string {
	return format(ioBookingIDPrefix, fmt.Sprint(bookingID))
}

func FormatIODisplayID(number interface{}) string {
	return format(ioDisplayPrefix, fmt.Sprint(number))
}

func FormatIOEmployeeID(number interface{}) string {
	return format(ioEmployeePrefix, fmt.Sprint(number))
}

func FormatIOEmployeeNumber(number interface{}) string {
	return format(ioEmployeeNumberPrefix, fmt.Sprint(number))
}

func FormatIOGroupBooking(number interface{}) string {
	return format(ioGroupBookingPrefix, fmt.Sprint(number))
}

func FormatIOMarket(id interface{}) string {
	return format(ioMarketPrefix, fmt.Sprint(id))
}

func FormatIONetworkID(number interface{}) string {
	return format(ioNetworkPrefix, fmt.Sprint(number))
}

func FormatIONetworkCode(number interface{}) string {
	return format(ioNetworkCodePrefix, fmt.Sprint(number))
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

func FormatQuattroNetworkID(sourceDbCode string, digitalProductID interface{}) string {
	return formatQuattroKey(sourceDbCode, quattroNetworkID, fmt.Sprint(digitalProductID))
}

/* GETTERS */

func GetCustomerOrder(externalIDs []string) string {
	return parseExternalID(extCustomerOrderPrefix, externalIDs)
}

func GetEmployeeNumber(externalIDs []string) string {
	return parseExternalID(ioEmployeeNumberPrefix, externalIDs)
}

func GetIOAccountID(externalIDs []string) string {
	return parseExternalID(ioAccountIdPrefix, externalIDs)
}

func GetIOBookingID(externalIDs []string) string {
	return parseExternalID(ioBookingIDPrefix, externalIDs)
}

func GetIOCreativeID(externalIDs []string) string {
	return parseExternalID(ioCreativePrefix, externalIDs)
}

func GetIOCustomerID(externalIDs []string) string {
	return parseExternalID(ioCustomerPrefix, externalIDs)
}

func GetIODocumentID(externalIDs []string) string {
	return parseExternalID(ioDocumentPrefix, externalIDs)
}

func GetIODisplayID(externalIDs []string) string {
	return parseExternalID(ioDisplayPrefix, externalIDs)
}

func GetIOEmployeeID(externalIDs []string) string {
	return parseExternalID(ioEmployeePrefix, externalIDs)
}

func GetIOGroupBookingID(externalIDs []string) string {
	return parseExternalID(ioGroupBookingPrefix, externalIDs)
}

func GetIOOrderID(externalIDs []string) string {
	return parseExternalID(ioOrderIDPrefix, externalIDs)
}

func GetIOOrderLineID(externalIDs []string) string {
	return parseExternalID(ioOrderLinePrefix, externalIDs)
}

func GetIOOrderMarketID(externalIDs []string) string {
	return parseExternalID(ioOrderMarketIDPrefix, externalIDs)
}

func GetIOOrderLineTypeID(externalIDs []string) string {
	return parseExternalID(ioOrderLineTypePrefix, externalIDs)
}

func GetIONetworkID(externalIDs []string) string {
	return parseExternalID(ioNetworkPrefix, externalIDs)
}

func GetIONetworkCode(externalIDs []string) string {
	return parseExternalID(ioNetworkCodePrefix, externalIDs)
}

func GetIOProductMapID(externalIDs []string) string {
	return parseExternalID(ioProductMapIDPrefix, externalIDs)
}

func GetLegacyMarketID(externalIDs []string) string {
	return parseExternalID(ioMarketPrefix, externalIDs)
}

func GetLegacySiteCode(externalIDs []string) string {
	return getLegacyIOEntityNumericID(externalIDs, "site")
}

func GetLegacyIODisplayID(externalIDs []string) string {
	return getLegacyIOEntityNumericID(externalIDs, "display")
}

func GetOrderNumber(externalIDs []string) string {
	return parseExternalID(ioOrderNumberPrefix, externalIDs)
}

func GetQuattroBookingID(sourceDbCode string, externalIDs []string) string {
	prefix := formatQuattroKey(sourceDbCode, quattroBookingPrefix, "")
	return parseExternalID(prefix, externalIDs)
}

func GetQuattroCampaignID(marketCode string, externalIDs []string) string {
	prefix := formatQuattroKey(marketCode, quattroCampaignPrefix, "")
	return parseExternalID(prefix, externalIDs)
}

func GetQuattroCampaignSegmentID(marketCode string, externalIDs []string) string {
	prefix := formatQuattroKey(marketCode, quattroCampaignSegmentPrefix, "")
	return parseExternalID(prefix, externalIDs)
}

func GetQuattroDigitalBookingID(sourceDbCode string, externalIDs []string) string {
	prefix := formatQuattroKey(sourceDbCode, quattroDigitalBookingPrefix, "")
	extId := parseExternalID(prefix, externalIDs)
	if extId != "" {
		return extId
	}

	prefix = formatQuattroKey(sourceDbCode, oldQuattroDigitalBookingPrefix, "")
	return parseExternalID(prefix, externalIDs)
}

func GetQuattroDisplayID(sourceDbCode string, externalIDs []string) string {
	prefix := formatQuattroKey(sourceDbCode, quattroDisplayID, "")
	return parseExternalID(prefix, externalIDs)
}

func GetQuattroNetworkID(sourceDbCode string, externalIDs []string) string {
	prefix := formatQuattroKey(sourceDbCode, quattroNetworkID, "")
	return parseExternalID(prefix, externalIDs)
}

func GetSpotChartDisplay(externalIDs []string) string {
	return parseExternalID(spotChartDisplayPrefix, externalIDs)
}

func GetSpotChartScene(externalIDs []string) string {
	return parseExternalID(spotChartScenePrefix, externalIDs)
}

func GetSpotChartSegment(externalIDs []string) string {
	return parseExternalID(spotChartSegmentPrefix, externalIDs)
}

/* helpers */

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

func getLegacyIOEntityNumericID(externalIDs []string, entity string) string {
	prefix := fmt.Sprintf("io:%s:", entity)
	return parseExternalID(prefix, externalIDs)
}

func parseExternalID(prefix string, externalIDs []string) string {
	for _, externalID := range externalIDs {
		if strings.HasPrefix(externalID, prefix) {
			return strings.TrimPrefix(externalID, prefix)
		}
	}
	return ""
}
