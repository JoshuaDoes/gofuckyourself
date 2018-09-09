package swearfilter

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// SwearFilter contains settings for the swear filter
type SwearFilter struct {
	Enabled bool //Whether or not the swear filter is enabled

	//Options to tell the swear filter how to operate
	DisableNormalize                bool //Disables normalization of alphabetic characters if set to true (ex: à -> a)
	DisableSpacedTab                bool //Disables converting tabs to singular spaces (ex: [tab][tab] -> [space][space])
	DisableMultiWhitespaceStripping bool //Disables stripping down multiple whitespaces (ex: hello[space][space]world -> hello[space]world)
	DisableZeroWidthStripping       bool //Disables stripping zero-width spaces
	DisableSpacedBypass             bool //Disables testing for spaced bypasses (if hell is in filter, look for occurrances of h and detect only alphabetic characters that follow; ex: h[space]e[space]l[space]l[space] -> hell)

	BlacklistedWords []string //A list of words to blacklist
}

// New returns an initialized SwearFilter struct to check messages against
func New(disableNormalize, disableSpacedTab, disableMutliWhitespaceStripping, disableSpacedBypass bool, words ...string) (filter *SwearFilter) {
	filter = &SwearFilter{
		Enabled:          true,
		BlacklistedWords: words,
	}
	return
}

// Check checks if a message contains blacklisted words and returns whether or not it does, a list of blacklisted words if so, and any errors returned by external packages
func (filter *SwearFilter) Check(message string) (bool, []string, error) {
	if len(filter.BlacklistedWords) <= 0 {
		return false, nil, nil
	}

	fixedMessage := message

	if !filter.DisableNormalize {
		bytes := make([]byte, len(fixedMessage))
		normalize := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
			return unicode.Is(unicode.Mn, r)
		}), norm.NFC)
		_, _, err := normalize.Transform(bytes, []byte(fixedMessage), true)
		if err != nil {
			return false, nil, err
		}
		fixedMessage = string(bytes)
	}

	if !filter.DisableSpacedTab {
		fixedMessage = strings.Replace(fixedMessage, "\t", " ", -1)
	}

	if !filter.DisableZeroWidthStripping {
		fixedMessage = strings.Replace(fixedMessage, "\u200b", "", -1)
	}

	if !filter.DisableMultiWhitespaceStripping {
		regexLeadCloseWhitepace := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
		regexInsideWhitespace := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
		fixedMessage = regexLeadCloseWhitepace.ReplaceAllString(fixedMessage, "")
		fixedMessage = regexInsideWhitespace.ReplaceAllString(fixedMessage, "")
	}

	detectedSwears := make([]string, 0)
	for _, swear := range filter.BlacklistedWords {
		if strings.Contains(fixedMessage, swear) {
			detectedSwears = append(detectedSwears, swear)
		}
	}

	if len(detectedSwears) > 0 {
		return true, detectedSwears, nil
	}
	return false, nil, nil
}
