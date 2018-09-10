package pam_test

import (
	"../src/pam"
	"reflect"
	"testing"
)

func TestCallback_001(t *testing.T) {
	c := pam.CallBackAdd(TestCallback_001)
	v := pam.CallBackGet(c)
	if reflect.TypeOf(v) != reflect.TypeOf(TestCallback_001) {
		t.Error("Received unexpected value")
	}
	pam.CallBackDelete(c)
}

func TestCallback_002(t *testing.T) {
	defer func() {
		recover()
	}()
	c := pam.CallBackAdd(TestCallback_002)
	pam.CallBackGet(c + 1)
	t.Error("Expected a panic")
}

func TestCallback_003(t *testing.T) {
	defer func() {
		recover()
	}()
	c := pam.CallBackAdd(TestCallback_003)
	pam.CallBackDelete(c + 1)
	t.Error("Expected a panic")
}
