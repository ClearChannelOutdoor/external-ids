package externalIDs

import (
	"fmt"
	"strings"
)

const (
	extCustomerOrderPrefix = "customer:order:"

	// Geopath Formats
	geopathSegmentCodePrefix   = "geopath:segmentCode:"
	geopathTargetProfilePrefix = "geopath:targetProfile:"

	// IO Formats
	ioAccountIdPrefix      = "io:account:"
	ioBookingIDPrefix      = "io:booking:"
	ioCreativePrefix       = "io:creative:"
	ioCustomerPrefix       = "io:customer:"
	ioDocumentPrefix       = "io:document:"
	ioDisplayPrefix        = "io:display:"
	ioEmployeePrefix       = "io:employee:"
	ioEmployeeNumberPrefix = "io:employeeNumber:"
	ioGroupBookingPrefix   = "io:groupBooking:"
	ioOrderLineTypePrefix  = "io:lineType:"
	ioMarketPrefix         = "io:market:"
	ioNetworkPrefix        = "io:network:"
	ioNetworkCodePrefix    = "io:networkCode:"
	ioOrderIDPrefix        = "io:order:"
	ioOrderLinePrefix      = "io:orderLine:"
	ioOrderMarketIDPrefix  = "io:orderMarket:"
	ioOrderNumberPrefix    = "io:orderNumber:"
	ioProductMapIDPrefix   = "io:productMap:"

	// Quattro Formats
	quattroDbLongFormPrefix         = "quattro_"
	quattroBookingPrefix            = ":booking:"
	quattroCampaignPrefix           = ":campaign:"
	quattroDetailPrefix             = ":detail:"
	quattroSegmentPrefix            = ":segment:"
	quattroDigitalBookingPrefix     = ":digitalBooking:"
	quattroDisplayID                = ":display:"
	quattroMarketID                 = ":market:"
	quattroNetworkID                = ":network:"
	oldQuattroCampaignDetailPrefix  = ":campaignDetail:"
	oldQuattroCampaignSegmentPrefix = ":campaignSegment:"

	// Salesforce Formats
	sfDisplayPrefix       = "salesforce:display:"
	sfDisplayRatePrefix   = "salesforce:displayRate:"
	sfBucketPrefix        = "salesforce:bucket:"
	sfBucketNamePrefix    = "salesforce:bucketName:"
	sfMatrixPrefix        = "salesforce:matrix:"
	sfMatrixLinePrefix    = "salesforce:matrixLine:"
	sfNetworkPrefix       = "salesforce:network:"
	sfOpportunityPrefix   = "salesforce:opportunity:"
	sfPanelLineItemPrefix = "salesforce:panelLineItem:"
	sfPlanPrefix          = "salesforce:plan:"
	sfPlanNamePrefix      = "salesforce:planName:"

	// Spotchart Formats
	spotChartDisplayPrefix = "spotchart:display:"
	spotChartScenePrefix   = "spotchart:scene:"
	spotChartSegmentPrefix = "spotchart:segment:"
)

/* FORMATTERS */

// Geopath Formatters
func FormatGeopathTargetProfile(code interface{}) string {
	return format(geopathTargetProfilePrefix, fmt.Sprint(code))
}

func FormatGeopathSegmentCode(code interface{}) string {
	return format(geopathSegmentCodePrefix, fmt.Sprint(code))
}

// IO Formatters
func FormatIOAccount(id interface{}) string {
	return format(ioAccountIdPrefix, fmt.Sprint(id))
}

func FormatIOBookingID(bookingID interface{}) string {
	return format(ioBookingIDPrefix, fmt.Sprint(bookingID))
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

// Quattro Formatters
func FormatQuattroCampaign(dbLongForm string, campaignId interface{}) string {
	return formatQuattroKey(dbLongForm, quattroCampaignPrefix, fmt.Sprint(campaignId))
}

func FormatQuattroCampaignSegment(dbLongForm string, segmentId interface{}) string {
	return formatQuattroKey(dbLongForm, quattroSegmentPrefix, fmt.Sprint(segmentId))
}

func FormatQuattroCampaignDetail(dbLongForm string, detailId interface{}) string {
	return formatQuattroKey(dbLongForm, quattroDetailPrefix, fmt.Sprint(detailId))
}

func FormatQuattroBookingID(dbLongForm string, bookingID interface{}) string {
	return formatQuattroKey(dbLongForm, quattroBookingPrefix, fmt.Sprint(bookingID))
}

func FormatQuattroDigitalBookingID(dbLongForm string, bookingID interface{}) string {
	return formatQuattroKey(dbLongForm, quattroDigitalBookingPrefix, fmt.Sprint(bookingID))
}

func FormatQuattroDisplayID(dbLongForm string, panelID interface{}) string {
	return formatQuattroKey(dbLongForm, quattroDisplayID, fmt.Sprint(panelID))
}

func FormatQuattroMarketID(dbLongForm string, marketID interface{}) string {
	return formatQuattroKey(dbLongForm, quattroMarketID, fmt.Sprint(marketID))
}

func FormatQuattroNetworkID(dbLongForm string, digitalProductID interface{}) string {
	return formatQuattroKey(dbLongForm, quattroNetworkID, fmt.Sprint(digitalProductID))
}

// SF Formatters
func FormatSFDisplayID(id interface{}) string {
	return format(sfDisplayPrefix, fmt.Sprint(id))
}

func FormatSFDisplayRateID(id interface{}) string {
	return format(sfDisplayRatePrefix, fmt.Sprint(id))
}

func FormatSFBucketID(id interface{}) string {
	return format(sfBucketPrefix, fmt.Sprint(id))
}

func FormatSFBucketName(id interface{}) string {
	return format(sfBucketNamePrefix, fmt.Sprint(id))
}

func FormatSFMatrixID(id interface{}) string {
	return format(sfMatrixPrefix, fmt.Sprint(id))
}

func FormatSFMatrixLineID(id interface{}) string {
	return format(sfMatrixLinePrefix, fmt.Sprint(id))
}

func FormatSFNetworkID(id interface{}) string {
	return format(sfNetworkPrefix, fmt.Sprint(id))
}

func FormatSFOpportunityID(oppID interface{}) string {
	return format(sfOpportunityPrefix, fmt.Sprint(oppID))
}

func FormatSFPanelLineItemID(id interface{}) string {
	return format(sfPanelLineItemPrefix, fmt.Sprint(id))
}

func FormatSFPlanID(id interface{}) string {
	return format(sfPlanPrefix, fmt.Sprint(id))
}

func FormatSFPlanNameID(id interface{}) string {
	return format(sfPlanNamePrefix, fmt.Sprint(id))
}

// Spotchart Formatters
func FormatSpotChartDisplayID(displayID interface{}) string {
	return format(spotChartDisplayPrefix, fmt.Sprint(displayID))
}

func FormatSpotChartSceneID(sceneID interface{}) string {
	return format(spotChartScenePrefix, fmt.Sprint(sceneID))
}

func FormatSpotChartSegmentID(segmentID interface{}) string {
	return format(spotChartSegmentPrefix, fmt.Sprint(segmentID))
}

/* GETTERS */

func GetCustomerOrder(externalIDs []string) string {
	return parseExternalID(extCustomerOrderPrefix, externalIDs)
}

// Geopath Getters

// IO Getters

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

func GetIONetworkID(externalIDs []string) string {
	return parseExternalID(ioNetworkPrefix, externalIDs)
}

func GetIONetworkCode(externalIDs []string) string {
	return parseExternalID(ioNetworkCodePrefix, externalIDs)
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

// Quattro Getters

func GetQuattroBookingID(dbLongForm string, externalIDs []string) string {
	prefix := formatQuattroKey(dbLongForm, quattroBookingPrefix, "")
	return parseExternalID(prefix, externalIDs)
}

func GetQuattroCampaignID(dbLongForm string, externalIDs []string) string {
	prefix := formatQuattroKey(dbLongForm, quattroCampaignPrefix, "")
	return parseExternalID(prefix, externalIDs)
}

func GetQuattroCampaignSegmentID(dbLongForm string, externalIDs []string) string {
	prefix := formatQuattroKey(dbLongForm, quattroSegmentPrefix, "")
	extId := parseExternalID(prefix, externalIDs)
	if extId != "" {
		return extId
	}

	oldPrefix := formatQuattroKey(dbLongForm, oldQuattroCampaignSegmentPrefix, "")
	return parseExternalID(oldPrefix, externalIDs)
}

func GetQuattroDigitalBookingID(dbLongForm string, externalIDs []string) string {
	prefix := formatQuattroKey(dbLongForm, quattroDigitalBookingPrefix, "")
	return parseExternalID(prefix, externalIDs)
}

func GetQuattroDisplayID(dbLongForm string, externalIDs []string) string {
	prefix := formatQuattroKey(dbLongForm, quattroDisplayID, "")
	return parseExternalID(prefix, externalIDs)
}

func GetQuattroMarketID(dbLongForm string, externalIDs []string) string {
	prefix := formatQuattroKey(dbLongForm, quattroMarketID, "")
	return parseExternalID(prefix, externalIDs)
}

func GetQuattroNetworkID(dbLongForm string, externalIDs []string) string {
	prefix := formatQuattroKey(dbLongForm, quattroNetworkID, "")
	return parseExternalID(prefix, externalIDs)
}

// SF Getters
func GetSFDisplayID(externalIDs []string) string {
	return parseExternalID(sfDisplayPrefix, externalIDs)
}

func GetSFDisplayRateID(externalIDs []string) string {
	return parseExternalID(sfDisplayRatePrefix, externalIDs)
}

func GetSFBucketID(externalIDs []string) string {
	return parseExternalID(sfBucketPrefix, externalIDs)
}

func GetSFBucketName(externalIDs []string) string {
	return parseExternalID(sfBucketNamePrefix, externalIDs)
}

func GetSFMatrixID(externalIDs []string) string {
	return parseExternalID(sfMatrixPrefix, externalIDs)
}

func GetSFMatrixLineID(externalIDs []string) string {
	return parseExternalID(sfMatrixLinePrefix, externalIDs)
}

func GetSFNetworkID(externalIDs []string) string {
	return parseExternalID(sfNetworkPrefix, externalIDs)
}

func GetSFOpportunityID(externalIDs []string) string {
	return parseExternalID(sfOpportunityPrefix, externalIDs)
}

func GetSFPanelLineItemID(externalIDs []string) string {
	return parseExternalID(sfPanelLineItemPrefix, externalIDs)
}

func GetSFPlanID(externalIDs []string) string {
	return parseExternalID(sfPlanPrefix, externalIDs)
}

func GetSFPlanNameID(externalIDs []string) string {
	return parseExternalID(sfPlanNamePrefix, externalIDs)
}

// SpotChart Getters
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

func formatQuattroKey(dbLongForm string, entity string, id string) string {
	quattroKey := format(quattroDbLongFormPrefix, dbLongForm)
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
