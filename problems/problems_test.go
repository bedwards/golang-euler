package problems

import "testing"

func TestPickNext(t *testing.T) {
	p, err := New(1)
	if err != nil {
		t.Errorf("Could not construct problem: %#v", err)
	}
	cases := []struct {
		factors []int
		ceiling int
		want    int
	}{
		{[]int{3, 5}, 10, 23},
	}
	for _, c := range cases {
		_, got, err := p.Solve(c.factors, c.ceiling)
		if err != nil {
			t.Errorf("Solve(%#v, %#v) should not err: %#v", c.factors, c.ceiling, err)
		}
		if got != c.want {
			t.Errorf("Solve(%#v, %#v) == %#v, want %#v", c.factors, c.ceiling, got, c.want)
		}
	}
}
