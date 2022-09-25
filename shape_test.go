package gopersian

import (
	"testing"
)

func TestShape(t *testing.T) {
	cases := []struct {
		tag  string
		in   string
		want string
	}{
		{"t1", "", ""},
		{"t2", "سلام مهران", "\ufeb3\ufee0\ufe8e\ufee1 \ufee3\ufeec\ufeae\u0627\ufee5"},
		{"t3", "سلام مهران abc", "\ufeb3\ufee0\ufe8e\ufee1 \ufee3\ufeec\ufeae\u0627\ufee5 abc"},
	}

	for _, c := range cases {
		t.Run(c.tag, func(t *testing.T) {
			got := Shape(c.in)
			//for _, r := range got {
			//	fmt.Printf("%+q", r)
			//}
			if got != c.want {
				t.Errorf("invalid shape result, wanted: %s, got: %s", c.want, got)
			}
		})
	}
}

func TestReorder(t *testing.T) {
	tests := []struct {
		Tag      string
		in, want string
	}{
		{"t1", "Hello, world", "Hello, world"},
		{"t1", "خوبی", "یبوخ"},
		{"t6", "خوبی؟ drink water 1234 چ خبر؟", "؟ربخ چ drink water 1234 ؟یبوخ"},
		{"t10", "", ""},
	}

	for _, c := range tests {
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

func TestBidi(t *testing.T) {
	tests := []struct {
		Tag  string
		in   string
		want string
	}{
		{"t1", "سلام مهران", "\ufee5\u0627\ufeae\ufeec\ufee3 \ufee1\ufe8e\ufee0\ufeb3"},
		{"t2", "", ""},
	}

	for _, c := range tests {
		t.Run(c.Tag, func(t *testing.T) {
			got, err := Bidi(c.in)
			if err != nil {
				t.Error("got error", err)
			}
			if got != c.want {
				t.Errorf("invalid result, wanted: %s, got: %s", c.want, got)
			}
		})

	}
}
