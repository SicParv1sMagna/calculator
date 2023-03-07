package uniq

import (
	"testing"
)

func TestCountLines(t *testing.T) {
	line := []string{"I love music.", "I love music.", "I love music.", "", "I love music of Katrik.", "I love music of Katrik.", "Thanks.", "I love music of Katrik.", "I love music of Katrik."}
	num_of_coincedenece := CountLines(&line, false)
	correct := []int{3, 1, 2, 1, 2}
	for i := 0; i < len(num_of_coincedenece); i++ {
		if correct[i] != num_of_coincedenece[i] {
			t.Fatalf("Error: number of coincedence wrong")
		}
	}
}

func TestDoubleLines(t *testing.T) {
	line := []string{"I love music.", "I love music.", "I love music.", "", "I love music of Katrik.", "I love music of Katrik.", "Thanks.", "I love music of Katrik.", "I love music of Katrik."}
	DoubleLines(&line, false)
	correct := []string{"I love music.", "I love music of Katrik.", "I love music of Katrik."}
	for i := 0; i < len(line); i++ {
		if correct[i] != line[i] {
			t.Fatalf("Error: double lines are incorrect")
		}
	}
}

func TestRegisterLines(t *testing.T) {
	line := []string{"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", "", "I love MuSIC of Katrik.", "I love music of katrik.", "Thanks.", "I love music of katrIk.", "I love MuSIC of Katrik."}
	RegisterLines(&line)
	correct := []string{"I LOVE MUSIC.", "", "I love MuSIC of Katrik.", "Thanks.", "I love MuSIC of Katrik."}
	for i := 0; i < len(line); i++ {
		if correct[i] != line[i] {
			t.Fatalf("Error: register of the lines are incorrect")
		}
	}
}

func TestSingleLines(t *testing.T) {
	line := []string{"I love music.", "I love music.", "I love music.", "", "I love music of Katrik.", "I love music of Katrik.", "Thanks.", "I love music of Katrik.", "I love music of Katrik."}
	SingleLines(&line, false)
	correct := []string{"", "Thanks."}
	for i := 0; i < len(line); i++ {
		if correct[i] != line[i] {
			t.Fatalf("Error: single lines are incorrect")
		}
	}
}

func TestUncountChar(t *testing.T) {
	line := []string{"I love music.", "I love music.", "I love music.", "", "I love music of Katrik.", "I love music of Katrik.", "Thanks.", "I love music of Katrik.", "I love music of Katrik."}
	UncountChars(5, &line, false)
	correct := []string{"e music.", "I love music.", "", "I love music of Katrik.", "Thanks.", "I love music of Katrik."}
	for i := 0; i < len(line); i++ {
		if correct[i] != line[i] {
			t.Fatal("Error: uncount chars are incorrect")
		}
	}
}

func TestUncountFields(t *testing.T) {
	line := []string{"I love music.", "I love music.", "I love music.", "", "I love music of Katrik.", "I love music of Katrik.", "Thanks.", "I love music of Katrik.", "I love music of Katrik."}
	UncountFields(5, &line, false)
	correct := []string{"I love music of Katrik.", "Thanks.", "I love music of Katrik."}
	for i := 0; i < len(line); i++ {
		if correct[i] != line[i] {
			t.Fatal("Error: uncount fields are incorrect")
		}
	}
}

func TestUniq(t *testing.T) {
	line := []string{"I love music.", "I love music.", "I love music.", "", "I love music of Katrik.", "I love music of Katrik.", "Thanks.", "I love music of Katrik.", "I love music of Katrik."}
	UniqStrings(&line, false)
	correct := []string{"I love music.", "", "I love music of Katrik.", "Thanks.", "I love music of Katrik."}
	for i := 0; i < len(line); i++ {
		if correct[i] != line[i] {
			t.Fatal("Error: uniq string are incorrect")
		}
	}
}
