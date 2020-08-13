package svgparser_test

import (
	"strings"

	"github.com/mfmayt/svgparser"
)

func element(name string, attrs map[string]string) *svgparser.Element {
	return &svgparser.Element{
		Name:       name,
		Attributes: attrs,
		Children:   []*svgparser.Element{},
	}
}

func parse(svg string, validate bool) (*svgparser.Element, error) {
	element, err := svgparser.Parse(strings.NewReader(svg), validate)
	return element, err
}

func render(*svgparser.Element) (string, error) {
	return "test", nil
}
