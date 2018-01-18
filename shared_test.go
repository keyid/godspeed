package godspeed

import (
	"reflect"
	"testing"
)

func Test_uniqueTags(t *testing.T) {
	tests := []struct {
		i1, i2 []string
		o      []string
	}{
		{
			i1: []string{"these", "are", "some", "example", "some", "tags", "are"},
			o:  []string{"these", "are", "some", "example", "tags"},
		},
		{
			i1: []string{"a", "b", "c"},
			i2: []string{"d", "e", "a"},
			o:  []string{"a", "b", "c", "d", "e"},
		},
	}

	for _, tt := range tests {
		out := uniqueTags(tt.i1, tt.i2)
		if !reflect.DeepEqual(tt.o, out) {
			t.Errorf("uniqueTag(%#v, %#v) = %#v; want %#v", tt.i1, tt.i2, out, tt.o)
		}
	}
}

func BenchmarkUniqueTags(b *testing.B) {
	res := make([][]string, b.N)

	tags1 := []string{
		"this", "is", "some:tag", "with", "a", "value", "and",
		"I am", "just going to", "keep", "m", "a", "k", "i", "n", "g", " ",
		"tags so", "we", "have", "a", "decent", "number", "of", "tags", "to",
		"benchmark", ".", "I was", "hoping", "to be", "able", "to use", "more",
		"tags.", "Also,", "did I", "invent", "my", "own", "lorem",
		"i", "p", "s", "u", "m", "?",
	}

	tags2 := []string{
		"back", "again", "with", "some", "nonsense", "tags", "to", "pad", "how",
		"much data", "we", "are", "using", "to", "test", ".", "I", "have", "become",
		"much", "lazier", "since", "tags1", "so", "we", "will", "have", "a", "few",
		"l", "e", "s", "s", " ", "t", "a", "g", "s",
	}

	for i := 0; i < b.N; i++ {
		res[i] = uniqueTags(tags1, tags2)
	}
}
