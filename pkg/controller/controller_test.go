package controller

import "testing"

func TestController(t *testing.T) {

	err := Controller()
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
