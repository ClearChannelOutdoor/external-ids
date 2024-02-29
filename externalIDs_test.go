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

func Test_GetIOLegacyNumericEntityID(t *testing.T) {
	extId := []string{"io:display:1234"}
	id, ok := GetLegacyIOEntityNumericID(extId, "display")
	if !ok {
		t.Errorf("bad io display id format; expected: %s; got: %d", "1234", id)
	}
	if *id != 1234 {
		t.Errorf("bad io display id format; expected: %d; got: %d", 1234, *id)
	}
}

func Test_GetIOLegacyNumericEntityID_NotRequestedEntity(t *testing.T) {
	extId := []string{"io:display:1234"}
	id, ok := GetLegacyIOEntityNumericID(extId, "campaign")
	if ok {
		t.Errorf("expected not ok %s : %d", "1234", id)
	}
	if id != nil {
		t.Errorf("expected nil got %d", id)
	}
}

func Test_GetIOLegacyNumericEntityID_NonConforming(t *testing.T) {
	tooManyParts := []string{"io:display:1234:5678"}
	id, ok := GetLegacyIOEntityNumericID(tooManyParts, "display")
	if ok {
		t.Errorf("too many parts expected not ok %s : %d", "1234", id)
	}
	if id != nil {
		t.Errorf("too many parts expected nil got %d", id)
	}

	tooFewParts := []string{"io:display"}
	id, ok = GetLegacyIOEntityNumericID(tooFewParts, "display")
	if ok {
		t.Errorf("too few parts expected not ok %s : %d", "1234", id)
	}
	if id != nil {
		t.Errorf("too few parts expected nil got %d", id)
	}
}

// only returns the first matching id
func Test_GetIOLegacyNumericEntityID_MultipleIDs(t *testing.T) {
	extId := []string{"io:market:1234", "io:display:1234", "io:display:5678"}
	id, ok := GetLegacyIOEntityNumericID(extId, "display")
	if !ok {
		t.Errorf("expected ok %s : %d", "1234", id)
	}
	if *id != 1234 {
		t.Errorf("expected %d got %d", 1234, *id)
	}
}

// test non-numeric id string
func Test_GetIOLegacyNumericEntityID_NonNumericID(t *testing.T) {
	extId := []string{"io:display:abc"}
	id, ok := GetLegacyIOEntityNumericID(extId, "display")
	if ok {
		t.Errorf("expected not ok %s : %d", "abc", id)
	}
	if id != nil {
		t.Errorf("expected nil got %d", id)
	}
}
