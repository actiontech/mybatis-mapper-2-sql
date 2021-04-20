package ast

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type IfNode struct {
	Expression string
}

type ChooseNode struct {
	When      []*WhenNode
	Otherwise *OtherwiseNode
}

type WhenNode struct {
	Expression string
}

type TrimNode struct {
}

type OtherwiseNode struct {
}

type CharData struct {
	tmp    bytes.Buffer
	reader *bytes.Reader
	Data   []Data
}

func NewData(data []byte) *CharData {
	return &CharData{
		tmp:    bytes.Buffer{},
		reader: bytes.NewReader(data),
		Data:   []Data{},
	}
}

func (d *CharData) Scan(start *xml.StartElement) error {
	return nil
}

func (d *CharData) AddChildren(ns ...Node) error {
	return nil
}

func (d *CharData) GetStmt(ctx *Context) (string, error) {
	buff := bytes.Buffer{}
	for _, child := range d.Data {
		switch dt := child.(type) {
		case Value:
			buff.WriteString(dt.String())
		case *Param:
			buff.WriteString(dt.String())
		case *Variable:
			variable, ok := ctx.GetVariable(dt.Name)
			if !ok {
				return "", fmt.Errorf("variable %s undifine", dt.Name)
			}
			buff.WriteString(variable)
		}
	}
	return buff.String(), nil
}

func (d *CharData) String() string {
	buff := bytes.Buffer{}
	for _, child := range d.Data {
		buff.WriteString(child.String())
	}
	return buff.String()
}

func (d *CharData) ScanData() error {
	for {
		var err error
		r, err := d.read()
		if err == io.EOF { // found end of element
			break
		}
		if err != nil {
			return err
		}

		switch r {
		case '#':
			s, err := d.read()
			if err == io.EOF { // found end of element
				break
			}
			if s == '{' {
				err := d.scanParam()
				if err != nil {
					return err
				}
			}
		case '$':
			s, err := d.read()
			if err == io.EOF { // found end of element
				break
			}
			if s == '{' {
				err := d.scanVariable()
				if err != nil {
					return err
				}
			}
		default:
			err := d.scanValue()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *CharData) read() (rune, error) {
	r, _, err := d.reader.ReadRune()
	if err != nil {
		return r, err
	}
	_, err = d.tmp.WriteRune(r)
	if err != nil {
		return r, err
	}
	return r, nil
}

func (d *CharData) unRead() {
	d.reader.Seek(-1, io.SeekCurrent)
}

func (d *CharData) clean() {
	d.tmp.Reset()
}

func (d *CharData) scanParam() error {
	d.clean()
	for {
		r, err := d.read()
		if err == io.EOF {
			return fmt.Errorf("data is invalid, not found \"}\" for param")
		}
		if err != nil {
			return err
		}
		if r == '}' {
			break
		}
	}
	data := strings.TrimSuffix(d.tmp.String(), "}")
	d.Data = append(d.Data, &Param{Name: data})
	d.clean()
	return nil
}

func (d *CharData) scanVariable() error {
	d.clean()
	for {
		r, err := d.read()
		if err == io.EOF {
			return fmt.Errorf("data is invalid, not found \"}\" for vaiable")
		}
		if err != nil {
			return err
		}
		if r == '}' {
			break
		}
	}
	data := strings.TrimSuffix(d.tmp.String(), "}")
	d.Data = append(d.Data, &Variable{Name: data})
	d.clean()
	return nil
}

func (d *CharData) scanValue() error {
	var first rune
	var second rune
	for {
		r, err := d.read()
		if err == io.EOF { // found end of element
			break
		}
		if err != nil {
			return err
		}
		if r == '#' || r == '$' {
			first = r
			s, err := d.read()
			if err == io.EOF { // found end of element
				break
			}
			second = s
			if s == '{' {
				d.unRead()
				d.unRead()
				break
			}
		}
	}
	data := strings.TrimSuffix(d.tmp.String(), string([]rune{first, second}))
	d.Data = append(d.Data, Value(data))
	d.clean()
	return nil
}

type Data interface {
	String() string
}

type Value string

func (v Value) String() string {
	return string(v)
}

type Param struct {
	Name string
}

func (p *Param) String() string {
	return "?"
}

type Variable struct {
	Name string
}

func (p *Variable) String() string {
	return "$"
}
