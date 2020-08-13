package svgparser

import (
	"encoding/xml"
	"fmt"
	"io"
)

// Render renders element to SVG
func (e *Element) Render(w io.Writer) error {
	encoder := xml.NewEncoder(w)

	if err := e.Encode(encoder); err != nil {
		return fmt.Errorf("Could not render element: %s", err)
	}

	return encoder.Flush()
}

// Serialize serializes element
func (e *Element) Serialize() xml.StartElement {
	var attributes []xml.Attr
	for name, value := range e.Attributes {
		attr := xml.Attr{
			Name:  xml.Name{Local: name},
			Value: value,
		}
		attributes = append(attributes, attr)
	}

	return xml.StartElement{
		Name: xml.Name{Local: e.Name},
		Attr: attributes,
	}
}

// Encode encodes the element
func (e *Element) Encode(encoder *xml.Encoder) error {
	start := e.Serialize()

	if err := encoder.EncodeToken(start); err != nil {
		return err
	}
	end := start.End()

	for _, child := range e.Children {
		if err := child.Encode(encoder); err != nil {
			return err
		}
	}
	var content xml.Token

	content = xml.CharData(e.Content)

	encoder.EncodeToken(content)
	return encoder.EncodeToken(end)
}
