package cmd

import "testing"

func assertEquals(t *testing.T, result interface{}, expectedResult interface{}) {
	if result != expectedResult {
		t.Errorf("got %t, expected %t", result, expectedResult)
	}
}
