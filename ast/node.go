package ast

import "encoding/xml"

type Node interface {
	Scan(start *xml.StartElement) error
	String() string
	AddChildren(ns ...Node) error
}

type ChildrenNode struct {
	children []Node
}

func NewNode() *ChildrenNode {
	return &ChildrenNode{
		children: []Node{},
	}
}

func (n *ChildrenNode) AddChildren(ns ...Node) error {
	n.children = append(n.children, ns...)
	return nil
}
