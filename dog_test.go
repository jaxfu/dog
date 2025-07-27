package dog

import (
	"encoding/json"
	"testing"
)

const (
	TESTFILEPATH string = "test.txt"
)

func TestDog(t *testing.T) {
	lines, err := Get(TESTFILEPATH, DogOptions{})
	if err != nil {
		t.Error(err)
	}

	j, _ := json.MarshalIndent(lines, "", "    ")
	t.Log(string(j))
}
