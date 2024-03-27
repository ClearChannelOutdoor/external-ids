package externalIDs

import "testing"

func Test_FormatQuattroCampaign(t *testing.T) {
	mar := "CHI"
	id := "1234"

	extId := FormatQuattroCampaign(mar, id)
	if extId != "quattro_CHI:campaign:1234" {
		t.Errorf("bad quattro campaign format; expected: %s; got: %s", "quattro_CHI:campaign:1234", extId)
	}
}

func Test_FormatQuattroCampaignSegment(t *testing.T) {
	mar := "CHI"
	id := "1234"

	segmentId := FormatQuattroCampaignSegment(mar, id)
	if segmentId != "quattro_CHI:campaignSegment:1234" {
		t.Errorf("bad quattro campaignSegment format; expected: %s; got: %s", "quattro_CHI:campaignSegment:1234", segmentId)
	}
}

func Test_FormatQuattroCampaignDetail(t *testing.T) {
	mar := "CHI"
	id := "1234"

	detailId := FormatQuattroCampaignDetail(mar, id)
	if detailId != "quattro_CHI:campaignDetail:1234" {
		t.Errorf("bad quattro campaignSegment format; expected: %s; got: %s", "quattro_CHI:campaignDetail:1234", detailId)
	}
}

func Test_FormatQuattroDisplayID(t *testing.T) {
	mar := "CHI"
	id := "1234"

	extId := FormatQuattroDisplayID(mar, id)
	if extId != "quattro_CHI:display:1234" {
		t.Errorf("bad quattro display format; expected: %s; got: %s", "quattro_CHI:display:1234", extId)
	}
}

func Test_FormatQuattroNetworkID(t *testing.T) {
	mar := "CHI"
	id := "1234"

	extId := FormatQuattroNetworkID(mar, id)
	if extId != "quattro_CHI:network:1234" {
		t.Errorf("bad quattro display format; expected: %s; got: %s", "quattro_CHI:display:1234", extId)
	}
}

func Test_FormatQuattroBookingID(t *testing.T) {
	mar := "CHI"
	id := "1234"

	extId := FormatQuattroBookingID(mar, id)
	if extId != "quattro_CHI:booking:1234" {
		t.Errorf("bad quattro display format; expected: %s; got: %s", "quattro_CHI:booking:1234", extId)
	}
}

func Test_FormatQuattroDigitalBookingID(t *testing.T) {
	mar := "CHI"
	id := "1234"

	extId := FormatQuattroDigitalBookingID(mar, id)
	if extId != "quattro_CHI:digitalBooking:1234" {
		t.Errorf("bad quattro display format; expected: %s; got: %s", "quattro_CHI:digitalBooking:1234", extId)
	}
}

func Test_GetIOLegacyDisplayID(t *testing.T) {
	extId := []string{"io:display:1234"}
	id := GetLegacyIODisplayID(extId)
	if *id != 1234 {
		t.Errorf("bad io display id format; expected: %d; got: %d", 1234, *id)
	}
}

func Test_GetIOLegacyDisplayID_NotRequestedEntity(t *testing.T) {
	extId := []string{"io:campaign:1234"}
	id := GetLegacyIODisplayID(extId)
	if id != nil {
		t.Errorf("expected nil got %d", &id)
	}
}

// only returns the first matching id
func Test_GetIOLegacyDisplayID_MultipleIDs(t *testing.T) {
	extIds := []string{"io:market:1234", "io:display:1234", "io:display:5678"}
	id := GetLegacyIODisplayID(extIds)
	if *id != 1234 {
		t.Errorf("expected %d got %d", 1234, *id)
	}
}

// test non-numeric id string
func Test_GetIOLegacyDisplayID_NonNumericID(t *testing.T) {
	extIds := []string{"io:display:abc"}
	id := GetLegacyIODisplayID(extIds)
	if id != nil {
		t.Errorf("expected nil got %d", id)
	}
}

// test getting site code
func Test_GetLegacySiteCode(t *testing.T) {
	extId := []string{"io:site:1234"}
	id := GetLegacySiteCode(extId)
	if *id != 1234 {
		t.Errorf("bad site_code format; expected: %d; got: %d", 1234, *id)
	}
}

func Test_GetQuattroBookingID(t *testing.T) {
	extIds := []string{"io:market:1234", "quattro_CHI:booking:2600", "io:display:5678"}
	id := GetQuattroBookingID("CHI", extIds)
	if *id != 2600 {
		t.Errorf("expected %d got %d", 2600, *id)
	}
}

func Test_GetQuattroBookingID_NoBookingID(t *testing.T) {
	extIds := []string{"io:market:1234", "io:booking:2600", "io:display:5678"}
	id := GetQuattroBookingID("CHI", extIds)
	if id != nil {
		t.Errorf("expected nil got %d", *id)
	}
}

func Test_GetEmployeeNumber(t *testing.T) {
	extIds := []string{"io:employeeNumber:1234", "io:booking:2600", "io:display:5678"}
	id := GetEmployeeNumber(extIds)
	if *id != 1234 {
		t.Errorf("expected 1234 got %d", *id)
	}
}

func Test_GetOrderNumber(t *testing.T) {
	extIds := []string{"io:orderNumber:1234", "io:booking:2600", "io:display:5678"}
	id := GetOrderNumber(extIds)
	if *id != 1234 {
		t.Errorf("expected 1234 got %d", *id)
	}
}

func Test_GetQuattroCampaignID(t *testing.T) {
	extIds := []string{"quattro_CHI:campaign:1234", "io:booking:2600", "io:display:5678"}
	id := GetQuattroCampaignID("CHI", extIds)
	if *id != 1234 {
		t.Errorf("expected 1234 got %d", *id)
	}
}

func Test_GetExtAgencyCode(t *testing.T) {
	extIds := []string{"agency:code:1234", "io:booking:2600", "io:display:5678"}
	id := GetExtAgencyCode(extIds)
	if *id != 1234 {
		t.Errorf("expected 1234 got %d", *id)
	}
}
