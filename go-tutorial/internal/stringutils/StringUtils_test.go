package stringutils

import "testing"

func TestIsNumeric(t *testing.T) {

	//tests := []struct {
	//	input string
	//	want  bool
	//}{
	//	{"1345012", true},
	//	{"09b4t444", false},
	//}

	val := IsNumeric("1345012")
	if val == false {
		t.Errorf("get result %v", val)
	}
}
