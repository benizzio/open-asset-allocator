package service

import "strings"

const (
	// maxAssetTextSearchResults is the maximum number of assets returned for a text search.
	//
	// Authored by: OpenCode
	maxAssetTextSearchResults = 20
)

// parseAssetTextSearch converts free-text asset search input into required search terms. Quoted
// text remains one phrase, while unmatched quotes are treated as delimiters that are removed.
//
// Authored by: OpenCode
func parseAssetTextSearch(textSearch string) []string {
	textSearch = removeUnmatchedAssetTextSearchQuote(textSearch)

	var searchTerms = make([]string, 0)
	var unquotedTerm strings.Builder
	var quotedTerm strings.Builder
	var isQuoted = false

	appendUnquotedTerms := func() {
		appendAssetTextSearchTerms(&searchTerms, unquotedTerm.String())
		unquotedTerm.Reset()
	}

	for _, character := range textSearch {
		switch {
		case character == '"' && !isQuoted:
			appendUnquotedTerms()
			isQuoted = true
		case character == '"' && isQuoted:
			appendAssetTextSearchPhrase(&searchTerms, quotedTerm.String())
			quotedTerm.Reset()
			isQuoted = false
		case isQuoted:
			quotedTerm.WriteRune(character)
		default:
			unquotedTerm.WriteRune(character)
		}
	}

	if isQuoted {
		appendAssetTextSearchTerms(&searchTerms, quotedTerm.String())
	} else {
		appendUnquotedTerms()
	}

	return searchTerms
}

// removeUnmatchedAssetTextSearchQuote removes the final quote when the input contains an odd
// number of quote characters, leaving all remaining quotes available for phrase parsing.
//
// Authored by: OpenCode
func removeUnmatchedAssetTextSearchQuote(textSearch string) string {
	if strings.Count(textSearch, `"`)%2 == 0 {
		return textSearch
	}

	var lastQuoteIndex = strings.LastIndex(textSearch, `"`)
	return textSearch[:lastQuoteIndex] + textSearch[lastQuoteIndex+1:]
}

// appendAssetTextSearchTerms appends nonblank whitespace-separated search terms to the result.
//
// Authored by: OpenCode
func appendAssetTextSearchTerms(searchTerms *[]string, text string) {
	*searchTerms = append(*searchTerms, strings.Fields(text)...)
}

// appendAssetTextSearchPhrase appends a normalized nonblank quoted phrase to the result.
//
// Authored by: OpenCode
func appendAssetTextSearchPhrase(searchTerms *[]string, text string) {
	var phrase = strings.Join(strings.Fields(text), " ")
	if phrase != "" {
		*searchTerms = append(*searchTerms, phrase)
	}
}
