package xml

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Auxiliar functions to assist in tests.
var (
	osOpen = os.Open
)

// Parse parses a draw.io XML file and returns an MxFile struct.
func Parse(fileName string) (*MxFile, error) {
	xmlFile, err := osOpen(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer xmlFile.Close()

	var mxFile MxFile
	if err := xml.NewDecoder(xmlFile).Decode(&mxFile); err != nil {
		return nil, fmt.Errorf("error decoding XML: %w", err)
	}

	return &mxFile, nil
}
