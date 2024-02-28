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

func Test_FormatQuattroDisplayID(t *testing.T) {
	mar := "CHI"
	id := "1234"

	extId := FormatQuattroDisplayID(mar, id)
	if extId != "quattro_CHI:display:1234" {
		t.Errorf("bad quattro display format; expected: %s; got: %s", "quattro_CHI:display:1234", extId)
	}
}
