package metadata

import (
	"archive/zip"
	"encoding/xml"
	"strings"
)

// OfficeCoreProperty stores information about core property products
type OfficeCoreProperty struct {
	XMLName        xml.Name `xml:"coreProperties"`
	Creator        string   `xml:"creator"`
	LastModifiedBy string   `"lastModifiedBy"`
}

// OfficeAppProperty stores information about office apps
type OfficeAppProperty struct {
	XMLName     xml.Name `xml:"Properties"`
	Application string   `xml:"Application"`
	Company     string   `xml:"Company"`
	Version     string   `xml:"AppVersion"`
}

var OfficeVersions = map[string]string{
	"16": "2016",
	"15": "2013",
	"14": "2010",
	"12": "2007",
	"11": "2003",
}

// GetMajorVersion returns the office version if it is in the map of OfficeVersions
func (a *OfficeAppProperty) GetMajorVersion() string {
	tokens := strings.Split(a.Version, ".") // Mutating a field of the struct
	if len(tokens) < 2 {
		return "Unknown"
	}
	v, ok := OfficeVersions[tokens[0]]
	if !ok {
		return "Unknown"
	}
	return v
}

// NewProperties returns populated copies of the properties from core and app.
func NewProperties(r *zip.Reader) (*OfficeCoreProperty, *OfficeAppProperty, error) {
	var coreProps OfficeCoreProperty
	var appProps OfficeAppProperty

	for _, f := range r.File {
		switch f.Name {
		case "docProps/core.xml":
			if err := process(f, &coreProps); err != nil {
				return nil, nil, err
			}
		case "docProps/app.xml":
			if err := process(f, &appProps); err != nil {
				return nil, nil, err
			}
		default:
			continue
		}
	}
	return &coreProps, &appProps, nil
}

// A prop. interface{} b/c the function doesn't care about what the structure to populate
// looks like.
func process(f *zip.File, prop interface{}) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()
	if err := xml.NewDecoder(rc).Decode(&prop); err != nil {
		return err
	}
	return nil
}
