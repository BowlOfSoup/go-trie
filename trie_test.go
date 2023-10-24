package trie

import (
	"reflect"
	"testing"
)

var testMapOfInstruments = map[string]string{
	"piano":           "keyboard",
	"violin":          "string",
	"viola":           "string",
	"cello":           "string",
	"trumpet":         "brass",
	"trombone":        "brass",
	"tuba":            "brass",
	"flute":           "woodwind",
	"clarinet":        "woodwind",
	"oboe":            "woodwind",
	"bassoon":         "woodwind",
	"saxophone":       "woodwind",
	"drums":           "percussion",
	"timpani":         "percussion",
	"xylophone":       "percussion",
	"marimba":         "percussion",
	"triangle":        "percussion",
	"guitar":          "string",
	"bass guitar":     "string",
	"double bass":     "string",
	"harpsichord":     "keyboard",
	"organ":           "keyboard",
	"accordion":       "keyboard",
	"harmonica":       "free reed",
	"bagpipes":        "free reed",
	"sitar":           "string",
	"banjo":           "string",
	"mandolin":        "string",
	"ukulele":         "string",
	"theremin":        "electronic",
	"synthesizer":     "electronic",
	"electric guitar": "string",
	"bongo":           "percussion",
	"conga":           "percussion",
	"didgeridoo":      "brass",
	"clavichord":      "keyboard",
	"lute":            "string",
	"zither":          "string",
	"pan flute":       "woodwind",
	"piccolo":         "woodwind",
	"recorder":        "woodwind",
	"cor anglais":     "woodwind",
	"french horn":     "brass",
	"euphonium":       "brass",
	"cornet":          "brass",
	"bugle":           "brass",
	"tambourine":      "percussion",
	"castanets":       "percussion",
	"cabasa":          "percussion",
	"guiro":           "percussion",
}

// For asserting slices (slices are randomly sorted).
func slicesToMap(s []string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range s {
		m[v] = true
	}
	return m
}

func TestTrieLookup(t *testing.T) {
	trie := NewTrie()

	for instrument, category := range testMapOfInstruments {
		trie.Insert(instrument, category)
	}

	// Test Lookup
	tests := []struct {
		prefix string
		values []string
	}{
		// Single letters
		{"p", []string{"keyboard", "woodwind", "woodwind"}},
		{"v", []string{"string", "string"}},
		{"c", []string{"percussion", "woodwind", "woodwind", "keyboard", "percussion", "brass", "percussion", "string"}},
		{"t", []string{"percussion", "electronic", "percussion", "brass", "brass", "brass", "percussion"}},
		{"f", []string{"brass", "woodwind"}},
		{"s", []string{"string", "electronic", "woodwind"}},
		{"d", []string{"percussion", "string", "brass"}},
		{"g", []string{"percussion", "string"}},
		{"h", []string{"keyboard", "free reed"}},
		{"b", []string{"string", "string", "woodwind", "percussion", "free reed", "brass"}},

		// Double letters
		{"pi", []string{"keyboard", "woodwind"}},
		{"vi", []string{"string", "string"}},
		{"ce", []string{"string"}},
		{"tr", []string{"percussion", "brass", "brass"}},
		{"fl", []string{"woodwind"}},
		{"sa", []string{"woodwind"}},
		{"dr", []string{"percussion"}},
		{"gu", []string{"percussion", "string"}},
		{"sy", []string{"electronic"}},
		{"el", []string{"string"}},
		{"bo", []string{"percussion"}},
		{"zi", []string{"string"}},
		{"pa", []string{"woodwind"}},
		{"re", []string{"woodwind"}},
		{"co", []string{"percussion", "woodwind", "brass"}},
		{"ta", []string{"percussion"}},
		{"ca", []string{"percussion", "percussion"}},
		{"ha", []string{"keyboard", "free reed"}},
		{"ba", []string{"string", "string", "woodwind", "free reed"}},

		// More letters
		{"pia", []string{"keyboard"}},
		{"vio", []string{"string", "string"}},
		{"cel", []string{"string"}},
		{"tru", []string{"brass"}},
		{"flu", []string{"woodwind"}},
		{"sax", []string{"woodwind"}},
		{"dru", []string{"percussion"}},
		{"gui", []string{"percussion", "string"}},
		{"syn", []string{"electronic"}},
		{"ele", []string{"string"}},
		{"bong", []string{"percussion"}},
		{"zit", []string{"string"}},
		{"pan", []string{"woodwind"}},
		{"rec", []string{"woodwind"}},
		{"cor", []string{"woodwind", "brass"}},
		{"tam", []string{"percussion"}},
		{"cas", []string{"percussion"}},
		{"cab", []string{"percussion"}},
		{"harp", []string{"keyboard"}},
		{"bag", []string{"free reed"}},

		// Full words
		{"piano", []string{"keyboard"}},
		{"violin", []string{"string"}},
		{"cello", []string{"string"}},
		{"trumpet", []string{"brass"}},
		{"flute", []string{"woodwind"}},
		{"saxophone", []string{"woodwind"}},
		{"drums", []string{"percussion"}},
		{"guitar", []string{"string"}},
		{"synthesizer", []string{"electronic"}},
		{"electric guitar", []string{"string"}},
		{"bongo", []string{"percussion"}},
		{"zither", []string{"string"}},
		{"pan flute", []string{"woodwind"}},
		{"recorder", []string{"woodwind"}},
		{"cor anglais", []string{"woodwind"}},
		{"tambourine", []string{"percussion"}},
		{"castanets", []string{"percussion"}},
		{"cabasa", []string{"percussion"}},
		{"guiro", []string{"percussion"}},

		// Negative scenarios
		{"xyz", []string{}},          // Non-existent prefix
		{"harpoon", []string{}},      // Non-existent word
		{"pianist", []string{}},      // Prefix exists but full word doesn't
		{"guitars", []string{}},      // Plural form of an existing word
		{"trumpets", []string{}},     // Plural form of an existing word
		{"synthesizers", []string{}}, // Plural form of an existing word
		{"flutes", []string{}},       // Plural form of an existing word
		{"zz", []string{}},           // Non-existent double letter prefix
		{"qq", []string{}},           // Non-existent single letter
	}

	for _, test := range tests {
		got := trie.Lookup(test.prefix)
		if !reflect.DeepEqual(slicesToMap(got), slicesToMap(test.values)) {
			t.Errorf("For prefix %s, expected %v but got %v", test.prefix, test.values, got)
		}
	}
}

func TestLookupUnique(t *testing.T) {
	trie := NewTrie()

	for instrument, category := range testMapOfInstruments {
		trie.Insert(instrument, category)
	}

	// Test LookupUnique
	tests := []struct {
		prefix string
		values []string
	}{
		{"p", []string{"keyboard", "woodwind"}},
		{"v", []string{"string"}},
		{"c", []string{"percussion", "woodwind", "keyboard", "brass", "string"}},
		{"ba", []string{"string", "woodwind", "free reed"}},
		{"piano", []string{"keyboard"}},
		{"xyz", []string{}},
		{"harpoon", []string{}},
		{"pianist", []string{}},
	}

	for _, test := range tests {
		got := trie.LookupUnique(test.prefix)
		if !reflect.DeepEqual(slicesToMap(got), slicesToMap(test.values)) {
			t.Errorf("For prefix %s, expected %v but got %v", test.prefix, test.values, got)
		}
	}
}

func TestEmptyKey(t *testing.T) {
	trie := NewTrie()

	// Insert a value with an empty key
	trie.Insert("", "value")

	// Ensure the trie remains empty after the insertion
	if len(trie.root.children) != 0 {
		t.Errorf("Expected trie root to have no children after inserting an empty key, but got %d children", len(trie.root.children))
	}

	// Lookup the empty key and ensure it returns an empty result
	values := trie.Lookup("")
	if len(values) != 0 {
		t.Errorf("Expected lookup of an empty key to return an empty slice, but got: %v", values)
	}

	// LookupUnique for the empty key and ensure it returns an empty result
	uniqueValues := trie.LookupUnique("")
	if len(uniqueValues) != 0 {
		t.Errorf("Expected unique lookup of an empty key to return an empty slice, but got: %v", uniqueValues)
	}
}
