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

func TestReorder(t *testing.T) {
	cases := []struct {
		Tag      string
		in, want string
	}{
		{"t1", "Hello, world", "Hello, world"},
		{"t2", "Hello, 世界", "Hello, 世界"},
		{"t5", "نص عربي", "يبرع صن"},
		// TODO: Fix other cases
		//{"t6", " خوبی؟ drink water 1234 چ خبر؟", "؟ربخ چ 1234 drink water ؟یبوخ "},
		//{"t7", "hi خوبی؟ drink water 1234 چ خبر؟", "؟ربخ چ 1234 drink water ؟یبوخ hi"},
		//{"t8", "نَصٌ عَربِيٌّ!", "!ٌّيِبرَع ٌصَن"},
		//{"t9", "نَصٌ example, عَربِيٌّ!", "!يٌّبِرعَ ,example صٌَن"},
		{"t10", "", ""},
	}

	for _, c := range cases {
		t.Run(c.Tag, func(t *testing.T) {
			got, err := Reorder(c.in)
			if err != nil {
				t.Error("got error", err)
				return
			}
			if got != c.want {
				t.Errorf("invalid reverse result, input; %s,  wanted:%s, got: %s", c.in, c.want, got)
			}
		})
	}
}
