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
