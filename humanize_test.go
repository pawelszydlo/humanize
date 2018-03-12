package humanize

import "testing"

func TestNew_Correct(t *testing.T) {
	var humanizer *Humanizer
	var err error
	humanizer, err = New("en")
	if humanizer == nil {
		t.Error("Humanizer creation failed.")
	}
	if err != nil {
		t.Errorf("Humanizer creation failed with error: %s", err)
	}
}

func TestNew_Wrong(t *testing.T) {
	var humanizer *Humanizer
	var err error
	humanizer, err = New("xyz")
	if humanizer != nil || err == nil {
		t.Error("Humanizer creation succeeded where it should have failed.")
	}
}
