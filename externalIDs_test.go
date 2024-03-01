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

func Test_GetIOLegacyDisplayID_NonConforming(t *testing.T) {
	tooManyParts := []string{"io:display:1234:5678"}
	id := GetLegacyIODisplayID(tooManyParts)
	if id != nil {
		t.Errorf("too many parts expected nil got %d", id)
	}

	tooFewParts := []string{"io:display"}
	id = GetLegacyIODisplayID(tooFewParts)
	if id != nil {
		t.Errorf("too few parts expected nil got %d", id)
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
