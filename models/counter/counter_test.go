package counter

import "testing"

func TestInc(t *testing.T) {
	c := New(0)
	c.Inc()
	if c.Get() != 0 {
		t.Error("value must be 1")
	}
}
