package vocab

import (
	"bufio"
	"os"
)

// Provider is an interface for exposing a vocab
type Provider interface {
	Vocab() Dict
}

// ID is used to identify vocab items
type ID int32

// Int32 int32 representation of an ID
func (id ID) Int32() int32 {
	return int32(id)
}

// Dict is a container for tokens
// NOTE: python uses an OrderedDict, unsure of implications
type Dict struct {
	Token2Index map[string]ID
	Index2Token map[ID]string
}

// FromFile will read a newline delimited file into a Dict
func FromFile(path string) (Dict, error) {
	// TODO test
	f, err := os.Open(path)
	if err != nil {
		// TODO wrap w/ stdlib
		return Dict{}, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	voc := Dict{Token2Index: map[string]ID{}, Index2Token: map[ID]string{}}
	for scanner.Scan() {
		voc.Add(scanner.Text())
	}
	return voc, nil
}

// New iwll return a a covab dict from the given tokens, IDs will match index
func New(tokens []string) Dict {
	token2Index := make(map[string]ID, len(tokens))
	index2Token := make(map[ID]string, len(tokens))
	for i, t := range tokens {
		token2Index[t] = ID(i)
		index2Token[ID(i)] = t
	}
	return Dict{Token2Index: token2Index, Index2Token: index2Token}
}

// Add will add an item to the vocabulary, is not thread-safe
func (v Dict) Add(token string) {
	v.Token2Index[token] = ID(v.Size())
}

// GetID will return the ID of the token in the vocab. Will be negative if it doesn't exists
func (v Dict) GetID(token string) ID {
	id, ok := v.Token2Index[token]
	if !ok {
		return ID(-1)
	}
	return ID(id)
}

// GetToken will get a token by the ID, returns the empty string if ID does not exist
func (v Dict) GetToken(id ID) string {
	token, ok := v.Index2Token[id]
	if !ok {
		return ""
	}
	return token
}

/*

// HasID returns true if the vocab contains the token
func (v Dict) HasID(id ID) bool {
	for k, v := range v.tokens {
		if v =
	}
}

// HasToken returns true if the
func (v Dict) HasToken(token string) bool {

}
*/

// Size returns the size of the vocabulary
func (v Dict) Size() int {
	return len(v.Token2Index)
}

// LongestSubstring returns the longest token that is a substring of the token
func (v Dict) LongestSubstring(token string) string {
	// Greedt, optimize to trie if needed
	for i := len(token); i > 0; i-- {
		sub := token[:i]
		if _, ok := v.Token2Index[sub]; ok {
			return sub
		}
	}
	return ""
}

/*
func (v Dict) ConvertItems(items []string) []ID {
	ids := make([]ID, len(items))
	for i, m := range items {
		ids[i] = v.tokens[m]
	}
	return ids
}

func (v Dict) ConvertTokens(tokens []string) []ID {
	return v.ConvertItems(tokens)
}
*/
