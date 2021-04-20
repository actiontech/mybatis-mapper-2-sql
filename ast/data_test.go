package ast

import (
	"testing"
)

func TestScan(t *testing.T) {
	d := NewData([]byte("asd #{include_target}${variable}asdasd"))
	err := d.ScanData()
	if err != nil {
		t.Errorf("parse error: %v", err)
		return
	}
	if len(d.Nodes) != 4 {
		t.Errorf("data length is 4, actual is %d", len(d.Nodes))
	}
}
