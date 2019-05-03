package golib

import (
	"testing"
)

func TestFindTransitdata(t *testing.T) {

	ts := FindTransitdata(500, 600)

	if ts == nil {
		t.Error("FindTransitdata error.")
	} else {
		t.Log("FindTransitdata passed.", ts)
	}

}
