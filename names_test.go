package names

import (
	"fmt"
	"strings"
	"testing"
)

// Test runs a demo + a dirty benchmark of the methods
func BenchmarkDockerNames(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(GetDockerName())
	}
}

func TestDockerNames(t *testing.T) {
	if !strings.Contains(GetDockerName(), "_") {
		t.Error("Didnt have _")
	}
}

func TestMultipleNames(t *testing.T) {
	for i := 0; i < 1000; i++ {
		fmt.Println(GetFirstName())
		fmt.Println(GetSurname())
		fmt.Println(GetName())
	}
}

func TestSentence(t *testing.T) {
	for i := 0; i < 1000; i++ {
		fmt.Println(MakeSentence2())
	}
}
