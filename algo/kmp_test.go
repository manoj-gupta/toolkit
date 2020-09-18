package algo

import (
	"testing"

	"github.com/manoj-gupta/toolkit/utils"
)

func TestKmpMatcher(t *testing.T) {
	var tests = []struct {
		text    string
		pattern string
		want    []int
	}{
		{"bacbabababcaabababca", "abababca", []int{4, 12}},
		{"AABAACAADAABAABA", "AABA", []int{0, 9, 12}},
	}

	for idx, test := range tests {
		got := KmpMatcher(test.text, test.pattern)

		if !utils.EqualIntSlices(got, test.want) {
			t.Errorf("test %d: failed - got:%v not same as want:%v", idx, got, test.want)
		}
	}

}
