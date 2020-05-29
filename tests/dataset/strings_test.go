package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (part: %s) - Indexes: expected (%d) <> found (%d)."

func TestIndex(t *testing.T) {
	tests := []struct {
		text     string
		part     string
		expected int
	}{
		{"In lorem ipsum", "In", 0},
		{"", "", 0},
		{"In lorem ipsum", "in", -1},
		{"In lorem ipsum", "lorem", 3},
	}

	for _, test := range tests {
		t.Logf("Dataset: %v", test)
		actual := strings.Index(test.text, test.part)

		if actual != test.expected {
			t.Errorf(msgIndex, test.text, test.part, test.expected, actual)
		}
	}
}
