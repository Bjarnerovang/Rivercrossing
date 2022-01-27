package state

import "testing"

// Test function implemented in line with the Golang's testing tool
func TestViewState(t *testing.T) {
	wanted := "[kylling rev korn hs ---\\ \\__/ _________________/---]"
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
func TestPutInBoat(t *testing.T) {
	wanted := "rev"
	var state = PutInBoat()
	if state != wanted {
		t.Errorf("Feil, fikk #{state}, ønsket #{wanted}")
	}
}
