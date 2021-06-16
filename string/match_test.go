package string

import "testing"

func TestStringMatch(t *testing.T) {
	strs := "{}{{}}[{}]({[]})"
	valid := isValid(strs)
	if valid == false {
		t.Error("expect is true actual is ", valid)
	}

}
