package gopersian

import (
	"fmt"
	"testing"
)

func TestShape(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"سلام مهران", "\ufeb3\ufee0\ufe8e\ufee1 \ufee3\ufeec\ufeae\u0627\ufee5"},
		{"", ""},
	}

	for _, c := range cases {
		got := Shape(c.in)
		for _, r := range got {
			fmt.Printf("%+q", r)
		}
		if got != c.want {
			t.Errorf("invalid shape result, wanted: %s, got: %s", c.want, got)
		}
	}
}

// Reverse returns its argument string reversed rune-wise left to right.
func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Crowdbotics", "scitobdworC"},
		{"Hello, 世界", "界世 ,olleH"},
		{"نص عربي", "يبرع صن"},
		{"نَصٌ عَربِيٌّ", "ٌّيِبرَع ٌصَن"},
		{"سلام خوبی؟", "؟یبوخ مالس"},
		{"نَصٌ عَربِيٌّ!", "!ٌّيِبرَع ٌصَن"},
		{"نَصٌ example, عَربِيٌّ!", "!ٌّيِبرَع ,elpmaxe ٌصَن"},
		{"", ""},
	}

	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("invalid reverse result, input; %s,  wanted:%s, got: %s", c.in, c.want, got)
		}
	}
}
