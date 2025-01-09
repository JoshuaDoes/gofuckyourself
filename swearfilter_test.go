package swearfilter

import (
	"testing"
)

func TestNew(t *testing.T) {
	filter := NewSwearFilter(true, "fuck", "hell")
	if filter.DisableNormalize {
		t.Errorf("Filter option DisableNormalize was incorrect, got: %t, want: %t", filter.DisableNormalize, false)
	}
	if filter.DisableSpacedTab {
		t.Errorf("Filter option DisableSpacedTab was incorrect, got: %t, want: %t", filter.DisableSpacedTab, false)
	}
	if filter.DisableMultiWhitespaceStripping {
		t.Errorf("Filter option DisableMultiWhitespaceStripping was incorrect, got: %t, want: %t", filter.DisableMultiWhitespaceStripping, false)
	}
	if filter.DisableZeroWidthStripping {
		t.Errorf("Filter option DisableZeroWidthStripping was incorrect, got: %t, want: %t", filter.DisableZeroWidthStripping, false)
	}
	if !filter.EnableSpacedBypass {
		t.Errorf("Filter option EnableSpacedBypass was incorrect, got: %t, want: %t", filter.EnableSpacedBypass, true)
	}
	if filter.DisableLeetSpeak {
		t.Errorf("Filter option DisableLeetSpeak was incorrect, got: %t, want: %t", filter.EnableSpacedBypass, false)
	}
	if len(filter.BadWords) != 2 {
		t.Errorf("Filter option BadWords was incorrect, got length: %d, want length: %d", len(filter.BadWords), 2)
	}
}
func TestCheckAndAddDelete(t *testing.T) {
	filter := NewSwearFilter(true)
	var prev []string

	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		// Basic cases
		//{"basic match", "fucking", []string{"fuck"}},
		//{"unicode chars", "fûçk", []string{"fuck"}},
		//{"clean text", "asdf", nil},
		//{"spaced out", "f u c k", []string{"fuck"}},

		// Multi-char leet speak cases (from multiCharLeet map)
		//{"ph substitution", "phuck", []string{"fuck"}},

		// Single-char leet speak cases (from leetChars map)
		//{"basic a4", "b4st4rd", []string{"bastard"}},
		//{"at symbol", "@sshole", []string{"asshole"}},
		{"euro sign", `p€nis`, []string{"penis"}},
		// {"multi char w", "vvhor3", []string{"whore"}},
		//{"spaced chars", "s h ! t", []string{"shit"}},
		//{"multiple mappings", "fvv|<", []string{"fuck"}},
		//{"o variants", "b()()bs", []string{"boobs"}},
		//{"ph mapping", "ph@rt", []string{"fart"}},
		// {"double v w", `\/\/@nk3r`, []string{"wanker"}},
		{"pure numbers", "5417", []string{"salt"}},
		//{"mixed special", "@$$h0l3", []string{"asshole"}},
		//{"x variant", "><xx", []string{"xxx"}},
		//{"dollar sign", "$hit", []string{"shit"}},
		//{"number mix", "8!7ch", []string{"bitch"}},
		//{"brackets o", "c[]c|<", []string{"cock"}},
		//{"uu variant", "uuank", []string{"wank"}},
		//{"hash symbol", "#0rny", []string{"horny"}},
		//{"g variants", "6i6", []string{"gig"}},
		//{"plus sign", "+i+s", []string{"tits"}},
		//{"j and z", "j2j2", []string{"izi"}},
		//{"alt k test", "1<un7", []string{"kunt"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Remove previous words
			if len(prev) > 0 {
				filter.Delete(prev...)
			}

			// Add new words from expected
			if tt.expected != nil {
				filter.Add(tt.expected...)
			}

			// Store current words for next iteration
			prev = tt.expected

			// Verify current word list
			currentWords := filter.Words()
			t.Logf("Current filter words: %v", currentWords)

			// Run the check
			trippers, err := filter.Check(tt.input)
			if err != nil {
				t.Errorf("Check failed: %v", err)
			}

			// Verify results
			if (trippers == nil && tt.expected != nil) || (trippers != nil && tt.expected == nil) {
				t.Errorf("got trippers %v, want %v ", trippers, tt.expected)
				return
			}
			if len(trippers) != len(tt.expected) {
				t.Errorf("got trippers length %d, want %d", len(trippers), len(tt.expected))
				return
			}

			// Check each expected word is in trippers
			for _, expected := range tt.expected {
				found := false
				for _, got := range trippers {
					if got == expected {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("expected word %s not found in trippers %v", expected, trippers)
				}
			}
		})
	}
}
func TestCheck(t *testing.T) {
	//filter := NewSwearFilter(true, "fuck", "hell")
	filter := NewSwearFilter(true,
		"fuck",
		"hell",
		"asshole",
		"bastard",
		"bitch",
		"boobs",
		"cock",
		"fart",
		"fuck",
		"gig",
		"horny",
		"izi",
		"kk",
		"kunt",
		"penis",
		"salt",
		"shit",
		"tits",
		"wank",
		"wanker",
		"whore",
		"xxx",
	)
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		// Basic cases
		// {"basic match", "fucking", []string{"fuck"}},
		// {"unicode chars", "fûçk", []string{"fuck"}},
		// {"clean text", "asdf", []string{}},
		// {"spaced out", "f u c k", []string{"fuck"}},
		//
		// // Multi-char leet speak cases (from multiCharLeet map)
		// {"ph substitution", "phuck", []string{"fuck"}},
		//
		// // Single-char leet speak cases (from leetChars map)
		// {"basic a4", "b4st4rd", []string{"bastard"}},
		// {"at symbol", "@sshole", []string{"asshole"}},
		// {"euro sign", "p€nis", []string{"penis"}},
		// {"multi char w", "vvhor3", []string{"whore"}},
		// {"spaced chars", "s h ! t", []string{"shit"}},
		// {"multiple mappings", "fvv|<", []string{"fuck"}},
		// {"o variants", "b()()bs", []string{"boobs"}},
		// {"ph mapping", "ph@rt", []string{"fart"}},
		// {"double v w", "\\/\\//@nk3r", []string{"wanker"}},
		// {"pure numbers", "5417", []string{"salt"}},
		// {"mixed special", "@$$h0l3", []string{"asshole"}},
		// {"x variant", "><xx", []string{"xxx"}},
		// {"dollar sign", "$hit", []string{"shit"}},
		// {"number mix", "8!7ch", []string{"bitch"}},
		// {"brackets o", "c[]c|<", []string{"cock"}},
		// {"uu variant", "uuank", []string{"wank"}},
		// {"hash symbol", "#0rny", []string{"horny"}},
		// {"g variants", "696", []string{"gig"}},
		// {"plus sign", "+i+s", []string{"tits"}},
		// {"j and z", "j2j2", []string{"izi"}},
		// {"alt k test", "1<un7", []string{"kunt"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trippers, err := filter.Check(tt.input)
			if err != nil {
				t.Errorf("Check failed: %v", err)
			}
			if (trippers == nil && tt.expected != nil) || (trippers != nil && tt.expected == nil) {
				t.Errorf("got trippers %v , want %v ", trippers, tt.expected)
				return
			}
			if len(trippers) != len(tt.expected) {
				t.Errorf("got trippers length %d, want %d", len(trippers), len(tt.expected))
				return
			}
			// Check each expected word is in trippers
			for _, expected := range tt.expected {
				found := false
				for _, got := range trippers {
					if got == expected {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("expected word %s not found in trippers %v", expected, trippers)
				}
			}
		})
	}
}
