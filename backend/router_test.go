package main

import (
	"testing"
)

func TestFormatAsDate(t *testing.T) {
	if "2018-08-17" != formatAsDate("2018-08-17 14:35:45") {
		t.Error("Parse datetime error")
	}
	if "14:35:45" != formatAsTime("2018-08-17 14:35:45") {
		t.Error("Parse datetime error")
	}
}