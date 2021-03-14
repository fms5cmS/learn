package some

import (
	"testing"
)

func TestFloat(t *testing.T) {
	doubleScore := func(source float32) (score float32) {
		defer func() {
			if score < 1 || score >= 100 {
				score = source
			}
		}()
		return source * 2
	}
	t.Log(doubleScore(0))
	t.Log(doubleScore(20.0))
	t.Log(doubleScore(50.0))
}
