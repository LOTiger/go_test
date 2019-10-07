package iteration

import (
	"testing"
	"fmt"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func ExampleRepeat()  {
	fmt.Printf(Repeat("a"))
	// Output:aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for index := 0; index < b.N; index++ {
		Repeat("a")
	}
}