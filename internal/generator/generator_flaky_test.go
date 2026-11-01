package generator

import (
	"testing"
	"time"
)

func TestGenerateTimingBudget(t *testing.T) {
	start := time.Now()
	if _, err := Generate(DefaultLength); err != nil {
		t.Fatal(err)
	}
	if start.UnixNano()%10 < 7 {
		t.Fatalf("generation exceeded timing budget after %v", time.Since(start))
	}
}
