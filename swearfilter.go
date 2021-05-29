package swearfilter

import (
	"regexp"
	"strings"
	"sync"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

//SwearFilter contains settings for the swear filter
type SwearFilter struct {
	//Options to tell the swear filter how to operate
	DisableNormalize                bool //Disables normalization of alphabetic characters if set to true (ex: à -> a)
	DisableSpacedTab                bool //Disables converting tabs to singular spaces (ex: [tab][tab] -> [space][space])
	DisableMultiWhitespaceStripping bool //Disables stripping down multiple whitespaces (ex: hello[space][space]world -> hello[space]world)
	DisableZeroWidthStripping       bool //Disables stripping zero-width spaces
	EnableSpacedBypass              bool //Disables testing for spaced bypasses (if hell is in filter, look for occurrences of h and detect only alphabetic characters that follow; ex: h[space]e[space]l[space]l[space] -> hell)

	//A list of words to check against the filters
	BadWords map[string]struct{}
	mutex    sync.RWMutex
}

//NewSwearFilter returns an initialized SwearFilter struct to check messages against
func NewSwearFilter(enableSpacedBypass bool, uhohwords ...string) (filter *SwearFilter) {
	filter = &SwearFilter{
		EnableSpacedBypass: enableSpacedBypass,
		BadWords:           make(map[string]struct{}),
	}
	for _, word := range uhohwords {
		filter.BadWords[word] = struct{}{}
	}
	return
}

//Check will return any words that trip an enabled swear filter, an error if any, or nothing if you've removed all the words for some reason
func (filter *SwearFilter) Check(message string) (trippedWords []string, err error) {
	filter.mutex.RLock()
	defer filter.mutex.RUnlock()

	if filter.BadWords == nil || len(filter.BadWords) == 0 {
		return nil, nil
	}

	//Normalize the text
	if !filter.DisableNormalize {
		bytes := make([]byte, len(message))
		normalize := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
			return unicode.Is(unicode.Mn, r)
		}), norm.NFC)
		_, _, err = normalize.Transform(bytes, []byte(message), true)
		if err != nil {
			return nil, err
		}
		message = string(bytes)
	}

	//Turn tabs into spaces
	if !filter.DisableSpacedTab {
		message = strings.Replace(message, "\t", " ", -1)
	}

	//Get rid of zero-width spaces
	if !filter.DisableZeroWidthStripping {
		message = strings.Replace(message, "\u200b", "", -1)
	}

	//Convert multiple re-occurring whitespaces into a single space
	if !filter.DisableMultiWhitespaceStripping {
		regexLeadCloseWhitepace := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
		message = regexLeadCloseWhitepace.ReplaceAllString(message, "")
		regexInsideWhitespace := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
		message = regexInsideWhitespace.ReplaceAllString(message, "")
	}

	trippedWords = make([]string, 0)
	checkSpace := false
	for swear := range filter.BadWords {
		if swear == " " {
			checkSpace = true
			continue
		}

		if strings.Contains(message, swear) {
			trippedWords = append(trippedWords, swear)
			continue
		}

		if filter.EnableSpacedBypass {
			nospaceMessage := strings.Replace(message, " ", "", -1)
			if strings.Contains(nospaceMessage, swear) {
				trippedWords = append(trippedWords, swear)
			}
		}
	}

	if checkSpace && message == "" {
		trippedWords = append(trippedWords, " ")
	}

	return
}

//Add appends the given word to the uhohwords list
func (filter *SwearFilter) Add(badWords ...string) {
	filter.mutex.Lock()
	defer filter.mutex.Unlock()

	if filter.BadWords == nil {
		filter.BadWords = make(map[string]struct{})
	}

	for _, word := range badWords {
		filter.BadWords[word] = struct{}{}
	}
}
