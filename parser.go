package main

import (
	"encoding/xml"
	"io"
	"mybatis_mapper_2_sql/ast"
	"strings"
)

func Parse(data string) (string, error) {
	r := strings.NewReader(data)
	d := xml.NewDecoder(r)
	n, err := parse(d, nil)
	if err != nil {
		return "", err
	}
	return n.String(), nil
}

func parse(d *xml.Decoder, start *xml.StartElement) (node ast.Node, err error) {
	if start != nil {
		node, err = ast.Scan(start)
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

			//s := string(tt)
			//if strings.TrimSpace(s) == "" {
			//	continue
			//}
			//fmt.Println("=======data=======")
			//fmt.Println(s)
			//fmt.Println("==================")
		default:
			continue
		}
	}
	return node, nil
}
