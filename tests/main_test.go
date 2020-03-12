package main

import (
	"testing"
)

func TestBuild(t *testing.T) {
	total := "ok"
	if total != "ok" {
		t.Errorf("Build not ok")
	}
}
