package main

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	
}


func TestFail(t *testing.T) {
	t.Errorf("this is supposed to fail")
}
