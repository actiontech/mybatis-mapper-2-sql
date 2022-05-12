package ast

import "testing"

func TestReplaceWhere(t *testing.T) {

	cases := []struct {
		input string
		expect string
	}{
		{
			"123",
			"123",
		},
		{
			"abc",
			"abc",
		},
		{
			"where123",
			"where123",
		},
		{
			"whereWhere",
			"whereWhere",
		},
		{
			"where####",
			"where####",
		},
		{
			"where 123",
			"AND 123",
		},
		{
			"where abc",
			"AND abc",
		},
		{
			"WHERE abc",
			"AND abc",
		},
		{
			"WHERE	abc",
			"AND\tabc",
		},
		{
			"WHERE\nabc",
			"AND\nabc",
		},
	}
	for _,c :=range cases {
		if actual:= replaceWhere(c.input); actual != c.expect {
			t.Errorf("test replaceWhere failed, input: %s, expect: %s, actual: %s\n", c.input, c.expect, actual)
		}
	}

}
