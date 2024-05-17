package externalIDs

import "testing"

func Test_FormatQuattroCampaign(t *testing.T) {
	key := "quattro_chicago"
	id := "1234"

	extId := FormatQuattroCampaign(key, id)
	if extId != "quattro_chicago:campaign:1234" {
		t.Errorf("bad quattro campaign format; expected: %s; got: %s", "quattro_chicago:campaign:1234", extId)
	}
}

func Test_FormatQuattroCampaignSegment(t *testing.T) {
	key := "quattro_chicago"
	id := "1234"

	segmentId := FormatQuattroCampaignSegment(key, id)
	if segmentId != "quattro_chicago:campaignSegment:1234" {
		t.Errorf("bad quattro campaignSegment format; expected: %s; got: %s", "quattro_chicago:campaignSegment:1234", segmentId)
	}
}

func Test_FormatQuattroCampaignDetail(t *testing.T) {
	key := "quattro_chicago"
	id := "1234"

	detailId := FormatQuattroCampaignDetail(key, id)
	if detailId != "quattro_chicago:campaignDetail:1234" {
		t.Errorf("bad quattro campaignSegment format; expected: %s; got: %s", "quattro_chicago:campaignDetail:1234", detailId)
	}
}

func Test_FormatQuattroDisplayID(t *testing.T) {
	key := "quattro_chicago"
	id := "1234"

	extId := FormatQuattroDisplayID(key, id)
	if extId != "quattro_chicago:display:1234" {
		t.Errorf("bad quattro display format; expected: %s; got: %s", "quattro_chicago:display:1234", extId)
	}
}

func Test_FormatQuattroNetworkID(t *testing.T) {
	key := "quattro_chicago"
	id := "1234"

	extId := FormatQuattroNetworkID(key, id)
	if extId != "quattro_chicago:network:1234" {
		t.Errorf("bad quattro display format; expected: %s; got: %s", "quattro_chicago:display:1234", extId)
	}
}

func Test_FormatQuattroBookingID(t *testing.T) {
	key := "quattro_chicago"
	id := "1234"

	extId := FormatQuattroBookingID(key, id)
	if extId != "quattro_chicago:booking:1234" {
		t.Errorf("bad quattro display format; expected: %s; got: %s", "quattro_chicago:booking:1234", extId)
	}
}

func Test_FormatQuattroDigitalBookingID(t *testing.T) {
	key := "quattro_chicago"
	id := "1234"

	extId := FormatQuattroDigitalBookingID(key, id)
	if extId != "quattro_chicago:digitalBooking:1234" {
		t.Errorf("bad quattro display format; expected: %s; got: %s", "quattro_chicago:digitalBooking:1234", extId)
	}
}

func Test_FormatIOEmployeeID(t *testing.T) {
	id := "1234"

	extId := FormatIOEmployeeID(id)
	if extId != "io:employee:1234" {
		t.Errorf("bad io employee id format; expected: %s; got: %s", "io:employee:1234", extId)
	}
}

func Test_GetIOLegacyDisplayID(t *testing.T) {
	extId := []string{"io:display:1234"}
	id := GetLegacyIODisplayID(extId)
	if id != "1234" {
		t.Errorf("bad io display id format; expected: %d; got: %s", 1234, id)
	}
}

func Test_GetIOLegacyDisplayID_NotRequestedEntity(t *testing.T) {
	extId := []string{"io:campaign:1234"}
	id := GetLegacyIODisplayID(extId)
	if id != "" {
		t.Errorf("expected nil got %s", id)
	}
}

// only returns the first matching id
func Test_GetIOLegacyDisplayID_MultipleIDs(t *testing.T) {
	extIds := []string{"io:market:1234", "io:display:1234", "io:display:5678"}
	id := GetLegacyIODisplayID(extIds)
	if id != "1234" {
		t.Errorf("expected %d got %s", 1234, id)
	}
}

func Test_GetIOEmployeeID(t *testing.T) {
	extIds := []string{"io:market:1234", "io:display:1234", "io:employee:5678"}
	id := GetLegacyIODisplayID(extIds)
	if id != "1234" {
		t.Errorf("expected %d got %s", 1234, id)
	}
}

// test getting site code
func Test_GetLegacySiteCode(t *testing.T) {
	extId := []string{"io:site:1234"}
	id := GetLegacySiteCode(extId)
	if id != "1234" {
		t.Errorf("bad site_code format; expected: %d; got: %s", 1234, id)
	}
}

func Test_GetQuattroBookingID(t *testing.T) {
	extIds := []string{"io:market:1234", "quattro_chicago:booking:2600", "io:display:5678"}
	id := GetQuattroBookingID("quattro_chicago", extIds)
	if id != "2600" {
		t.Errorf("expected %d got %s", 2600, id)
	}
}

func Test_GetQuattroDigitalBookingID(t *testing.T) {
	extIds := []string{"io:market:1234", "quattro_chicago:digitalBooking:2600", "io:display:5678"}
	id := GetQuattroDigitalBookingID("quattro_chicago", extIds)
	if id != "2600" {
		t.Errorf("expected %d got %s", 2600, id)
	}
}

func Test_GetQuattroBookingID_NoBookingID(t *testing.T) {
	extIds := []string{"io:market:1234", "io:booking:2600", "io:display:5678"}
	id := GetQuattroBookingID("quattro_chicago", extIds)
	if id != "" {
		t.Errorf("expected nil got %s", id)
	}
}

func Test_GetEmployeeNumber(t *testing.T) {
	extIds := []string{"io:employeeNumber:1234", "io:booking:2600", "io:display:5678"}
	id := GetEmployeeNumber(extIds)
	if id != "1234" {
		t.Errorf("expected 1234 got %s", id)
	}
}

func Test_GetOrderNumber(t *testing.T) {
	extIds := []string{"io:orderNumber:1234", "io:booking:2600", "io:display:5678"}
	id := GetOrderNumber(extIds)
	if id != "1234" {
		t.Errorf("expected 1234 got %s", id)
	}
}

func Test_GetQuattroCampaignID(t *testing.T) {
	extIds := []string{"quattro_chicago:campaign:1234", "io:booking:2600", "io:display:5678"}
	id := GetQuattroCampaignID("quattro_chicago", extIds)
	if id != "1234" {
		t.Errorf("expected 1234 got %s", id)
	}
}

func Test_GetQuattroCampaignSegmentID(t *testing.T) {
	extIds := []string{"quattro_chicago:campaignSegment:1234", "io:booking:2600", "io:display:5678"}
	id := GetQuattroCampaignSegmentID("quattro_chicago", extIds)
	if id != "1234" {
		t.Errorf("expected 1234 got %s", id)
	}
}

func Test_GetCustomerOrder(t *testing.T) {
	extIds := []string{"customer:order:1234", "io:booking:2600", "io:display:5678"}
	id := GetCustomerOrder(extIds)
	if id != "1234" {
		t.Errorf("expected 1234 got %s", id)
	}
}
