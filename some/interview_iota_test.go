package some

import "testing"

func TestIota(t *testing.T) {
	const (
		a = iota
		b = iota
	)
	const (
		c = "c"
		d = iota
		e = iota
	)
	t.Log(a, b, c, d, e)
}
