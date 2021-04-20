package main

import (
	"encoding/xml"
	"io"
	"mybatis_mapper_2_sql/ast"
	"strings"
)

func ParseXML(data string) (string, error) {
	r := strings.NewReader(data)
	d := xml.NewDecoder(r)
	n, err := parse(d, nil)
	if err != nil {
		return "", err
	}
	stmt, err := n.GetStmt(ast.NewContext())
	if err != nil {
		return "", err
	}
	return stmt, nil
}

func parse(d *xml.Decoder, start *xml.StartElement) (node ast.Node, err error) {
	if start != nil {
		node, err = scan(start)
		if err != nil {
			return nil, err
		}
	}

	for {
		t, err := d.Token()
		if err == io.EOF { // found end of element
			break
		}
		if err != nil {
			return nil, err
		}

		switch tt := t.(type) {
		case xml.StartElement:
			child, err := parse(d, &tt)
			if err != nil {
				return nil, err
			}
			if child == nil {
				continue
			}
			if node == nil {
				node = child
			} else {
				err := node.AddChildren(child)
				if err != nil {
					return nil, err
				}
			}
		case xml.EndElement:
			if start != nil && tt.Name == start.Name {
				return node, nil
			}
		case xml.CharData:
			s := string(tt)
			if strings.TrimSpace(s) == "" {
				continue
			}
			d := ast.NewData(tt)
			d.ScanData()
			if node != nil {
				node.AddChildren(d)
			}
		default:
			continue
		}
	}
	return node, nil
}

func scan(start *xml.StartElement) (ast.Node, error) {
	var node ast.Node
	switch start.Name.Local {
	case "mapper":
		node = ast.NewMapper()
	case "sql":
		node = ast.NewSqlNode()
	case "include":
		node = ast.NewIncludeNode()
	case "property":
		node = ast.NewPropertyNode()
	case "select", "update", "delete":
		node = ast.NewQueryNode()
	default:
		return node, nil
		//return node, fmt.Errorf("unknow xml %s", start.Name.Local)
	}
	node.Scan(start)
	return node, nil
}
