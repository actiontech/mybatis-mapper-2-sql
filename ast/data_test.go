package ast

import (
	"testing"
)

func TestMyBatisScan(t *testing.T) {
	d := NewMyBatisData([]byte("asd #{include_target}${variable}asdasd"))
	err := d.ScanData()
	if err != nil {
		t.Errorf("parse error: %v", err)
		return
	}
	if len(d.Nodes) != 4 {
		t.Errorf("data length is 4, actual is %d", len(d.Nodes))
		return
	}
	actual := d.Nodes[0].String()
	if actual != "asd " {
		t.Errorf("expect data is \"asd \", actual is %s", actual)
		return
	}
	actual = d.Nodes[1].String()
	if actual != "?" {
		t.Errorf("expect data is \"?\", actual is %s", actual)
		return
	}
	actual = d.Nodes[2].String()
	if actual != "$" {
		t.Errorf("expect data is \"$\", actual is %s", actual)
		return
	}
	actual = d.Nodes[3].String()
	if actual != "asdasd" {
		t.Errorf("expect data is \"asdasd\", actual is %s", actual)
		return
	}
}

func TestIBatisScan(t *testing.T) {
	d := NewIBatisData([]byte("asd #include_target#$v$asdasd"))
	err := d.ScanData()
	if err != nil {
		t.Errorf("parse error: %v", err)
		return
	}
	if len(d.Nodes) != 4 {
		t.Errorf("data length is 3, actual is %d", len(d.Nodes))
		//return
	}
	actual := d.Nodes[0].String()
	if actual != "asd " {
		t.Errorf("expect data is \"asd \", actual is %s", actual)
		return
	}
	actual = d.Nodes[1].String()
	if actual != "?" {
		t.Errorf("expect data is \"?\", actual is %s", actual)
		return
	}
	actual = d.Nodes[2].String()
	if actual != "$" {
		t.Errorf("expect data is \"$\", actual is %s", actual)
		return
	}
	actual = d.Nodes[3].String()
	if actual != "asdasd" {
		t.Errorf("expect data is \"asdasd\", actual is %s", actual)
		return
	}
}
