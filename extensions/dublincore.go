package ext

import "encoding/xml"

// DublinCoreExtension represents a feed extension
// for the Dublin Core specification.
type DublinCoreExtension struct {
	Title       []string `json:"title,omitempty"`
	Creator     []string `json:"creator,omitempty"`
	Author      []string `json:"author,omitempty"`
	Subject     []string `json:"subject,omitempty"`
	Description []string `json:"description,omitempty"`
	Publisher   []string `json:"publisher,omitempty"`
	Contributor []string `json:"contributor,omitempty"`
	Date        []string `json:"date,omitempty"`
	Type        []string `json:"type,omitempty"`
	Format      []string `json:"format,omitempty"`
	Identifier  []string `json:"identifier,omitempty"`
	Source      []string `json:"source,omitempty"`
	Language    []string `json:"language,omitempty"`
	Relation    []string `json:"relation,omitempty"`
	Coverage    []string `json:"coverage,omitempty"`
	Rights      []string `json:"rights,omitempty"`
}

// Encode will encode the dublincore extension in the provided xml encoder
func (dc DublinCoreExtension) Encode(e *xml.Encoder) error {
	encodeStringArray(e, "dc:title", dc.Title)
	encodeStringArray(e, "dc:creator", dc.Creator)
	encodeStringArray(e, "dc:author", dc.Author)
	encodeStringArray(e, "dc:subject", dc.Subject)
	encodeStringArray(e, "dc:description", dc.Description)
	encodeStringArray(e, "dc:publisher", dc.Publisher)
	encodeStringArray(e, "dc:contributor", dc.Contributor)
	encodeStringArray(e, "dc:date", dc.Date)
	encodeStringArray(e, "dc:type", dc.Type)
	encodeStringArray(e, "dc:format", dc.Format)
	encodeStringArray(e, "dc:identifier", dc.Identifier)
	encodeStringArray(e, "dc:source", dc.Source)
	encodeStringArray(e, "dc:language", dc.Language)
	encodeStringArray(e, "dc:relation", dc.Relation)
	encodeStringArray(e, "dc:coverage", dc.Coverage)
	encodeStringArray(e, "dc:rights", dc.Rights)

	return nil
}

// NewDublinCoreExtension creates a new DublinCoreExtension
// given the generic extension map for the "dc" prefix.
func NewDublinCoreExtension(extensions map[string][]Extension) *DublinCoreExtension {
	dc := &DublinCoreExtension{}
	dc.Title = parseTextArrayExtension("title", extensions)
	dc.Creator = parseTextArrayExtension("creator", extensions)
	dc.Author = parseTextArrayExtension("author", extensions)
	dc.Subject = parseTextArrayExtension("subject", extensions)
	dc.Description = parseTextArrayExtension("description", extensions)
	dc.Publisher = parseTextArrayExtension("publisher", extensions)
	dc.Contributor = parseTextArrayExtension("contributor", extensions)
	dc.Date = parseTextArrayExtension("date", extensions)
	dc.Type = parseTextArrayExtension("type", extensions)
	dc.Format = parseTextArrayExtension("format", extensions)
	dc.Identifier = parseTextArrayExtension("identifier", extensions)
	dc.Source = parseTextArrayExtension("source", extensions)
	dc.Language = parseTextArrayExtension("language", extensions)
	dc.Relation = parseTextArrayExtension("relation", extensions)
	dc.Coverage = parseTextArrayExtension("coverage", extensions)
	dc.Rights = parseTextArrayExtension("rights", extensions)
	return dc
}

func encodeStringArray(e *xml.Encoder, name string, val []string, attrs ...xml.Attr) error {
	if val == nil {
		return nil
	}

	for _, v := range val {
		encode(e, name, v)
	}
	return nil
}
