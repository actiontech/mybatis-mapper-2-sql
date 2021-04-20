package ast

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type Mapper struct {
	NameSpace  string
	SqlNodes   map[string]*SqlNode
	QueryNodes map[string]*QueryNode
}

func NewMapper() *Mapper {
	return &Mapper{
		SqlNodes:   map[string]*SqlNode{},
		QueryNodes: map[string]*QueryNode{},
	}
}

func (m *Mapper) AddChildren(ns ...Node) error {
	for _, n := range ns {
		switch nt := n.(type) {
		case *SqlNode:
			if _, ok := m.SqlNodes[nt.Id]; ok {
				return fmt.Errorf("sql id %s is repeat", nt.Id)
			}
			m.SqlNodes[nt.Id] = nt
		case *QueryNode:
			if _, ok := m.QueryNodes[nt.Id]; ok {
				return fmt.Errorf("%s id %s is repeat", nt.Type, nt.Id)
			}
			m.QueryNodes[nt.Id] = nt
		}
	}
	return nil
}

func (m *Mapper) Scan(start *xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "namespace" {
			m.NameSpace = attr.Value
		}
	}
	return nil
}

func (m *Mapper) String() string {
	buff := bytes.Buffer{}
	for _, child := range m.QueryNodes {
		buff.WriteString(child.String())
	}
	return buff.String()
}

func (m *Mapper) GetStmt(ctx *Context) (string, error) {
	buff := bytes.Buffer{}
	ctx.Sqls = m.SqlNodes
	for _, a := range m.QueryNodes {
		data, err := a.GetStmt(ctx)
		if err != nil {
			return "", err
		}
		buff.WriteString(data)
	}
	return buff.String(), nil
}
