package calorieCounting

import "testing"

func TestStrToIntArr(t *testing.T) {
	stringSplice := []string{"1", "2", "3"}
	result := strToIntArr(stringSplice)
	if len(result) != len(stringSplice) {
		t.Errorf("FAILED: expected a result of length %d and got length %d", len(stringSplice), len(result))
	} else {
		t.Logf("PASSED: expected a result of length %d and got length %d", len(stringSplice), len(result))
	}
}
