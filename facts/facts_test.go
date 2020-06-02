package facts

import (
	"testing"
)

func TestRandomFact(t *testing.T) {
	for i := 0; i < 100; i++ {
		fact := RandomFact()
		t.Log(fact)
		if len(fact) == 0 {
			t.Error("Fact is empty")
		}
	}
}
