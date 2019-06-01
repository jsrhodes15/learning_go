package iteration

import "fmt"
import "testing"


func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 9)
	expected := "aaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 9)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 7)
	fmt.Println(repeat)
	// Output: aaaaaaa
}
