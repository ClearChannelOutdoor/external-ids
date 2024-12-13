package externalIDs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IOFormatters(t *testing.T) {
	type testRun struct {
		name           string
		id             string
		fn             func(interface{}) string
		expectedResult string
	}

	runs := []testRun{
		{
			name:           "account",
			id:             "1234",
			fn:             FormatIOAccount,
			expectedResult: "io:account:1234",
		},
		{
			name:           "booking",
			id:             "1234",
			fn:             FormatIOBookingID,
			expectedResult: "io:booking:1234",
		},
		{
			name:           "creative",
			id:             "1234",
			fn:             FormatIOCreative,
			expectedResult: "io:creative:1234",
		},
		{
			name:           "customer",
			id:             "1234",
			fn:             FormatIOCustomer,
			expectedResult: "io:customer:1234",
		},
		{
			name:           "documents",
			id:             "1234",
			fn:             FormatIODocument,
			expectedResult: "io:document:1234",
		},
		{
			name:           "display",
			id:             "1234",
			fn:             FormatIODisplayID,
			expectedResult: "io:display:1234",
		},
		{
			name:           "employee",
			id:             "1234",
			fn:             FormatIOEmployeeID,
			expectedResult: "io:employee:1234",
		},
		{
			name:           "employee number",
			id:             "1234",
			fn:             FormatIOEmployeeNumber,
			expectedResult: "io:employeeNumber:1234",
		},
		{
			name:           "group booking",
			id:             "1234",
			fn:             FormatIOGroupBooking,
			expectedResult: "io:groupBooking:1234",
		},
		{
			name:           "order line type",
			id:             "1234",
			fn:             FormatIOOrderLineType,
			expectedResult: "io:lineType:1234",
		},
		{
			name:           "market",
			id:             "1234",
			fn:             FormatIOMarket,
			expectedResult: "io:market:1234",
		},
		{
			name:           "network",
			id:             "1234",
			fn:             FormatIONetworkID,
			expectedResult: "io:network:1234",
		},
		{
			name:           "networkCode",
			id:             "1234",
			fn:             FormatIONetworkCode,
			expectedResult: "io:networkCode:1234",
		},
		{
			name:           "order",
			id:             "1234",
			fn:             FormatIOOrderID,
			expectedResult: "io:order:1234",
		},
		{
			name:           "order line",
			id:             "1234",
			fn:             FormatIOOrderLine,
			expectedResult: "io:orderLine:1234",
		},
		{
			name:           "order market",
			id:             "1234",
			fn:             FormatIOOrderMarket,
			expectedResult: "io:orderMarket:1234",
		},
		{
			name:           "order number",
			id:             "1234",
			fn:             FormatIOOrderNumber,
			expectedResult: "io:orderNumber:1234",
		},
		{
			name:           "product map",
			id:             "1234",
			fn:             FormatIOProductMap,
			expectedResult: "io:productMap:1234",
		},
	}

	for _, test := range runs {
		t.Run(test.name, func(t *testing.T) {
			formattedId := test.fn(test.id)
			assert.Equal(t, test.expectedResult, formattedId)
		})
	}
}

func Test_QuattroFormatters(t *testing.T) {
	type testRun struct {
		name           string
		quattroDb      string
		id             string
		fn             func(string, interface{}) string
		expectedResult string
	}

	runs := []testRun{
		{
			name:           "booking",
			quattroDb:      "quattro_chicago",
			id:             "1234",
			fn:             FormatQuattroBookingID,
			expectedResult: "quattro_chicago:booking:1234",
		},
		{
			name:           "campaign",
			quattroDb:      "chicago",
			id:             "1234",
			fn:             FormatQuattroCampaign,
			expectedResult: "quattro_chicago:campaign:1234",
		},
		{
			name:           "campaign detail",
			quattroDb:      "chicago",
			id:             "1234",
			fn:             FormatQuattroCampaignDetail,
			expectedResult: "quattro_chicago:detail:1234",
		},
		{
			name:           "campaign segment",
			quattroDb:      "chicago",
			id:             "1234",
			fn:             FormatQuattroCampaignSegment,
			expectedResult: "quattro_chicago:segment:1234",
		},
		{
			name:           "digital booking",
			quattroDb:      "chicago",
			id:             "1234",
			fn:             FormatQuattroDigitalBookingID,
			expectedResult: "quattro_chicago:digitalBooking:1234",
		},
		{
			name:           "display",
			quattroDb:      "chicago",
			id:             "1234",
			fn:             FormatQuattroDisplayID,
			expectedResult: "quattro_chicago:display:1234",
		},
		{
			name:           "market",
			quattroDb:      "chicago",
			id:             "*",
			fn:             FormatQuattroMarketID,
			expectedResult: "quattro_chicago:market:*",
		},
		{
			name:           "network",
			quattroDb:      "chicago",
			id:             "1234",
			fn:             FormatQuattroNetworkID,
			expectedResult: "quattro_chicago:network:1234",
		},
	}

	for _, test := range runs {
		t.Run(test.name, func(t *testing.T) {
			formattedId := test.fn(test.quattroDb, test.id)
			assert.Equal(t, test.expectedResult, formattedId)
		})
	}
}

func Test_SalesforceFormatters(t *testing.T) {
	type testRun struct {
		name           string
		id             string
		fn             func(interface{}) string
		expectedResult string
	}

	runs := []testRun{
		{
			name:           "sf display",
			id:             "1234",
			fn:             FormatSFDisplayID,
			expectedResult: "salesforce:display:1234",
		},
		{
			name:           "sf display rate",
			id:             "1234",
			fn:             FormatSFDisplayRateID,
			expectedResult: "salesforce:displayRate:1234",
		},
		{
			name:           "sf bucket",
			id:             "1234",
			fn:             FormatSFBucketID,
			expectedResult: "salesforce:bucket:1234",
		},
		{
			name:           "sf matrix",
			id:             "1234",
			fn:             FormatSFMatrixID,
			expectedResult: "salesforce:matrix:1234",
		},
		{
			name:           "sf matrix line",
			id:             "1234",
			fn:             FormatSFMatrixLineID,
			expectedResult: "salesforce:matrixLine:1234",
		},
		{
			name:           "sf opportunity",
			id:             "1234",
			fn:             FormatSFOpportunityID,
			expectedResult: "salesforce:opportunity:1234",
		},
		{
			name:           "sf panel line item",
			id:             "1234",
			fn:             FormatSFPanelLineItemID,
			expectedResult: "salesforce:panelLineItem:1234",
		},
		{
			name:           "sf plan",
			id:             "1234",
			fn:             FormatSFPlanID,
			expectedResult: "salesforce:plan:1234",
		},
	}

	for _, test := range runs {
		t.Run(test.name, func(t *testing.T) {
			formattedId := test.fn(test.id)
			assert.Equal(t, test.expectedResult, formattedId)
		})
	}
}

func Test_SpotchartFormatters(t *testing.T) {
	type testRun struct {
		name           string
		id             string
		fn             func(interface{}) string
		expectedResult string
	}

	runs := []testRun{
		{
			name:           "spotchart display",
			id:             "1234",
			fn:             FormatSpotChartDisplayID,
			expectedResult: "spotchart:display:1234",
		},
		{
			name:           "spotchart scene",
			id:             "1234",
			fn:             FormatSpotChartSceneID,
			expectedResult: "spotchart:scene:1234",
		},
		{
			name:           "spotchart segment",
			id:             "1234",
			fn:             FormatSpotChartSegmentID,
			expectedResult: "spotchart:segment:1234",
		},
	}

	for _, test := range runs {
		t.Run(test.name, func(t *testing.T) {
			formattedId := test.fn(test.id)
			assert.Equal(t, test.expectedResult, formattedId)
		})
	}
}

func Test_MiscFormatters(t *testing.T) {
	type testRun struct {
		name           string
		id             string
		fn             func(interface{}) string
		expectedResult string
	}

	runs := []testRun{
		{
			name:           "customer order",
			id:             "1234",
			fn:             FormatCustomerOrder,
			expectedResult: "customer:order:1234",
		},
	}

	for _, test := range runs {
		t.Run(test.name, func(t *testing.T) {
			formattedId := test.fn(test.id)
			assert.Equal(t, test.expectedResult, formattedId)
		})
	}
}

func Test_IOGetters(t *testing.T) {
	type testRun struct {
		name           string
		externalIDs    []string
		fn             func([]string) string
		expectedResult string
	}

	runs := []testRun{
		{
			name:           "account",
			externalIDs:    []string{"io:account:1234"},
			fn:             GetIOAccountID,
			expectedResult: "1234",
		},
		{
			name:           "booking",
			externalIDs:    []string{"io:booking:1234"},
			fn:             GetIOBookingID,
			expectedResult: "1234",
		},
		{
			name:           "creative",
			externalIDs:    []string{"io:creative:1234"},
			fn:             GetIOCreativeID,
			expectedResult: "1234",
		},
		{
			name:           "customer",
			externalIDs:    []string{"io:customer:1234"},
			expectedResult: "1234",
			fn:             GetIOCustomerID,
		},
		{
			name:           "document",
			externalIDs:    []string{"io:document:1234"},
			expectedResult: "1234",
			fn:             GetIODocumentID,
		},
		{
			name:           "display",
			externalIDs:    []string{"io:display:1234"},
			expectedResult: "1234",
			fn:             GetIODisplayID,
		},
		{
			name:           "employee",
			externalIDs:    []string{"io:employee:1234"},
			fn:             GetIOEmployeeID,
			expectedResult: "1234",
		},
		{
			name:           "employee number",
			externalIDs:    []string{"io:employeeNumber:1234"},
			fn:             GetEmployeeNumber,
			expectedResult: "1234",
		},
		{
			name:           "group booking",
			externalIDs:    []string{"io:groupBooking:1234"},
			fn:             GetIOGroupBookingID,
			expectedResult: "1234",
		},
		{
			name:           "order line type",
			externalIDs:    []string{"io:lineType:1234"},
			fn:             GetIOOrderLineTypeID,
			expectedResult: "1234",
		},
		{
			name:           "market",
			externalIDs:    []string{"io:market:1234"},
			fn:             GetLegacyMarketID,
			expectedResult: "1234",
		},
		{
			name:           "network",
			externalIDs:    []string{"io:network:1234"},
			fn:             GetIONetworkID,
			expectedResult: "1234",
		},
		{
			name:           "order",
			externalIDs:    []string{"io:order:1234"},
			fn:             GetIOOrderID,
			expectedResult: "1234",
		},
		{
			name:           "order line",
			externalIDs:    []string{"io:orderLine:1234"},
			fn:             GetIOOrderLineID,
			expectedResult: "1234",
		},
		{
			name:           "order market",
			externalIDs:    []string{"io:orderMarket:1234"},
			fn:             GetIOOrderMarketID,
			expectedResult: "1234",
		},
		{
			name:           "order number",
			externalIDs:    []string{"io:orderNumber:1234"},
			fn:             GetOrderNumber,
			expectedResult: "1234",
		},
		{
			name:           "product map",
			externalIDs:    []string{"io:productMap:1234"},
			fn:             GetIOProductMapID,
			expectedResult: "1234",
		},
	}

	for _, test := range runs {
		t.Run(test.name, func(t *testing.T) {
			actualId := test.fn(test.externalIDs)
			assert.Equal(t, test.expectedResult, actualId)
		})
	}
}

func Test_QuattroGetters(t *testing.T) {
	type testRun struct {
		name           string
		quattroDb      string
		externalIDs    []string
		fn             func(string, []string) string
		expectedResult string
	}

	runs := []testRun{
		{
			name:           "booking",
			quattroDb:      "quattro_chicago",
			externalIDs:    []string{"quattro_chicago:booking:1234"},
			fn:             GetQuattroBookingID,
			expectedResult: "1234",
		},
		{
			name:           "booking, wrong market",
			quattroDb:      "quattro_el_paso",
			externalIDs:    []string{"quattro_chicago:booking:1234"},
			fn:             GetQuattroBookingID,
			expectedResult: "",
		},
		{
			name:           "campaign",
			quattroDb:      "chicago",
			externalIDs:    []string{"quattro_chicago:campaign:1234"},
			fn:             GetQuattroCampaignID,
			expectedResult: "1234",
		},
		{
			name:           "campaign, wrong market",
			quattroDb:      "el_paso",
			externalIDs:    []string{"quattro_chicago:campaign:1234"},
			fn:             GetQuattroCampaignID,
			expectedResult: "",
		},
		{
			name:           "campaign segment",
			quattroDb:      "chicago",
			externalIDs:    []string{"quattro_chicago:segment:1234"},
			fn:             GetQuattroCampaignSegmentID,
			expectedResult: "1234",
		},
		{
			name:           "old campaign segment format",
			quattroDb:      "chicago",
			externalIDs:    []string{"quattro_chicago:campaignSegment:1234"},
			fn:             GetQuattroCampaignSegmentID,
			expectedResult: "1234",
		},
		{
			name:           "campaign segment, wrong market",
			quattroDb:      "chicago",
			externalIDs:    []string{"el_paso:campaignSegment:1234"},
			fn:             GetQuattroCampaignSegmentID,
			expectedResult: "",
		},
		{
			name:           "digital booking",
			quattroDb:      "chicago",
			externalIDs:    []string{"quattro_chicago:digitalBooking:1234"},
			fn:             GetQuattroDigitalBookingID,
			expectedResult: "1234",
		},
		{
			name:           "digital booking, wrong market",
			quattroDb:      "chicago",
			externalIDs:    []string{"el_paso:digitalBooking:1234"},
			fn:             GetQuattroDigitalBookingID,
			expectedResult: "",
		},
		{
			name:           "display",
			quattroDb:      "chicago",
			externalIDs:    []string{"quattro_chicago:display:1234"},
			fn:             GetQuattroDisplayID,
			expectedResult: "1234",
		},
		{
			name:           "display, wrong market",
			quattroDb:      "chicago",
			externalIDs:    []string{"el_paso:display:1234"},
			fn:             GetQuattroDisplayID,
			expectedResult: "",
		},
		{
			name:           "market",
			quattroDb:      "chicago",
			externalIDs:    []string{"quattro_chicago:market:1234"},
			fn:             GetQuattroMarketID,
			expectedResult: "1234",
		},
		{
			name:           "network",
			quattroDb:      "chicago",
			externalIDs:    []string{"quattro_chicago:network:1234"},
			fn:             GetQuattroNetworkID,
			expectedResult: "1234",
		},
		{
			name:           "network, wrong market",
			quattroDb:      "el_paso",
			externalIDs:    []string{"quattro_chicago:network:1234"},
			fn:             GetQuattroNetworkID,
			expectedResult: "",
		},
	}

	for _, test := range runs {
		t.Run(test.name, func(t *testing.T) {
			extId := test.fn(test.quattroDb, test.externalIDs)
			assert.Equal(t, test.expectedResult, extId)
		})
	}
}

func Test_SalesforceGetters(t *testing.T) {
	type testRun struct {
		name           string
		externalIDs    []string
		fn             func([]string) string
		expectedResult string
	}

	runs := []testRun{
		{
			name:           "sf display",
			externalIDs:    []string{"salesforce:display:1234"},
			expectedResult: "1234",
			fn:             GetSFDisplayID,
		},
		{
			name:           "sf display rate",
			externalIDs:    []string{"salesforce:displayRate:1234"},
			expectedResult: "1234",
			fn:             GetSFDisplayRateID,
		},
		{
			name:           "sf bucket",
			externalIDs:    []string{"salesforce:bucket:1234"},
			expectedResult: "1234",
			fn:             GetSFBucketID,
		},
		{
			name:           "sf matrix",
			externalIDs:    []string{"salesforce:matrix:1234"},
			expectedResult: "1234",
			fn:             GetSFMatrixID,
		},
		{
			name:           "sf matrix line",
			externalIDs:    []string{"salesforce:matrixLine:1234"},
			expectedResult: "1234",
			fn:             GetSFMatrixLineID,
		},
		{
			name:           "sf opportunity",
			externalIDs:    []string{"salesforce:opportunity:1234"},
			expectedResult: "1234",
			fn:             GetSFOpportunityID,
		},
		{
			name:           "sf panel line item",
			externalIDs:    []string{"salesforce:panelLineItem:1234"},
			expectedResult: "1234",
			fn:             GetSFPanelLineItemID,
		},
		{
			name:           "sf plan",
			externalIDs:    []string{"salesforce:plan:1234"},
			expectedResult: "1234",
			fn:             GetSFPlanID,
		},
	}

	for _, test := range runs {
		t.Run(test.name, func(t *testing.T) {
			extId := test.fn(test.externalIDs)
			assert.Equal(t, test.expectedResult, extId)
		})
	}
}

func Test_SpotchartGetters(t *testing.T) {
	type testRun struct {
		name           string
		externalIDs    []string
		fn             func([]string) string
		expectedResult string
	}

	runs := []testRun{
		{
			name:           "spotchart display",
			externalIDs:    []string{"spotchart:display:1234"},
			expectedResult: "1234",
			fn:             GetSpotChartDisplay,
		},
		{
			name:           "spotchart scene",
			externalIDs:    []string{"spotchart:scene:1234"},
			expectedResult: "1234",
			fn:             GetSpotChartScene,
		},
		{
			name:           "spotchart segment",
			externalIDs:    []string{"spotchart:segment:1234"},
			expectedResult: "1234",
			fn:             GetSpotChartSegment,
		},
	}

	for _, test := range runs {
		t.Run(test.name, func(t *testing.T) {
			extId := test.fn(test.externalIDs)
			assert.Equal(t, test.expectedResult, extId)
		})
	}
}

func Test_MiscGetters(t *testing.T) {
	type testRun struct {
		name           string
		externalIDs    []string
		fn             func([]string) string
		expectedResult string
	}

	runs := []testRun{
		{
			name:           "customer order",
			externalIDs:    []string{"customer:order:1234"},
			expectedResult: "1234",
			fn:             GetCustomerOrder,
		},
	}

	for _, test := range runs {
		t.Run(test.name, func(t *testing.T) {
			extId := test.fn(test.externalIDs)
			assert.Equal(t, test.expectedResult, extId)
		})
	}
}
