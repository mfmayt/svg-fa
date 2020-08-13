package svgparser_test

import (
	"strings"
	"testing"

	"github.com/mfmayt/svgparser"
)

func TestRender(t *testing.T) {
	var testCases = []struct {
		svg     string
		element svgparser.Element
	}{
		{
			`<svg width="100" height="100"><circle cx="50" cy="50" r="40" fill="red"></circle></svg>`,
			svgparser.Element{
				Name: "svg",
				Attributes: map[string]string{
					"width":  "100",
					"height": "100",
				},
				Children: []*svgparser.Element{
					element("circle", map[string]string{"cx": "50", "cy": "50", "r": "40", "fill": "red"}),
				},
			},
		},
		{
			`<svg width="450" height="400"><g stroke-width="3" stroke="black" fill="black"><path id="AB" d="M 100 350 L 150 -300" stroke="red"></path></g></svg>`,
			svgparser.Element{
				Name: "svg",
				Attributes: map[string]string{
					"width":  "450",
					"height": "400",
				},
				Children: []*svgparser.Element{
					&svgparser.Element{
						Name: "g",
						Attributes: map[string]string{
							"stroke-width": "3",
							"stroke":       "black",
							"fill":         "black",
						},
						Children: []*svgparser.Element{
							element("path", map[string]string{"id": "AB", "d": "M 100 350 L 150 -300", "stroke": "red"}),
						},
					},
				},
			},
		},
	}

	for _, test := range testCases {
		actual, err := render(&test.element)

		if !(strings.Compare(test.svg, actual) == 0 && err == nil) {
			t.Errorf("Render: expected %v, actual %v\n", test.svg, actual)
		}
	}
}
