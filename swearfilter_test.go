package swearfilter

import (
	"testing"
)

func TestNew(t *testing.T) {
	filter := New(false, false, true, false, false, "fuck", "hell")
	if filter.DisableNormalize {
		t.Errorf("Filter option DisableNormalize was incorrect, got: %t, want: %t", filter.DisableNormalize, false)
	}
	if filter.DisableSpacedTab {
		t.Errorf("Filter option DisableSpacedTab was incorrect, got: %t, want: %t", filter.DisableSpacedTab, false)
	}
	if !filter.DisableMultiWhitespaceStripping {
		t.Errorf("Filter option DisableMultiWhitespaceStripping was incorrect, got: %t, want: %t", filter.DisableMultiWhitespaceStripping, true)
	}
	if filter.DisableZeroWidthStripping {
		t.Errorf("Filter option DisableZeroWidthStripping was incorrect, got: %t, want: %t", filter.DisableZeroWidthStripping, false)
	}
	if filter.DisableSpacedBypass {
		t.Errorf("Filter option DisableSpacedBypass was incorrect, got: %t, want: %t", filter.DisableSpacedBypass, false)
	}
	if len(filter.BlacklistedWords) != 2 {
		t.Errorf("Filter option BlacklistedWords was incorrect, got length: %d, want length: %d", len(filter.BlacklistedWords), 2)
	}
	if filter.BlacklistedWords[0] != "fuck" {
		t.Errorf("Filter option BlacklistedWords was incorrect, got first word: %s, want first word: %s", filter.BlacklistedWords[0], "fuck")
	}
	if filter.BlacklistedWords[1] != "hell" {
		t.Errorf("Filter option BlacklistedWords was incorrect, got second word: %s, want second word: %s", filter.BlacklistedWords[1], "hell")
	}
}

func TestCheck(t *testing.T) {
	filter := New(false, false, false, false, false, "fuck")
	messages := []string{"fucking", "fûçk", "asdf"}

	for i := 0; i < len(messages); i++ {
		tripped, trippers, err := filter.Check(messages[i])
		if err != nil {
			t.Errorf("Check failed due to external dependency: %v", err)
		}
		switch i {
		case 0, 1:
			if !tripped {
				t.Errorf("Check did not act as expected, got tripped: %t, want tripped: %t", tripped, true)
			}
			if len(trippers) != 1 {
				t.Errorf("Check did not act as expected, got trippers length: %d, want trippers length: %d", len(trippers), 1)
			}
			if trippers[0] != "fuck" {
				t.Errorf("Check did not act as expected, got first tripper: %s, want first tripper: %s", trippers[0], "fuck")
			}
		case 2:
			if tripped {
				t.Errorf("Check did not act as expected, got tripped: %t, want tripped: %t", tripped, false)
			}
			if len(trippers) != 0 {
				t.Errorf("Check did not act as expected, got trippers length: %d, want trippers length: %d", len(trippers), 0)
			}
		default:
			t.Errorf("Check test invalid, got test messages length: %d, want test messages length: %d", len(messages), 3)
		}
	}
}
