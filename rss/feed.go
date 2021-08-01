package rss

import (
	"encoding/json"
	"encoding/xml"
	"time"

	ext "github.com/mmcdole/gofeed/extensions"
)

// Feed is an RSS Feed
type Feed struct {
	RootName  string     `json:"-" xml:"-"`
	RootAttrs []xml.Attr `json:"-" xml:"-"`

	Title               string                   `json:"title,omitempty"                xml:"title,omitempty"`
	Link                string                   `json:"link,omitempty"                 xml:"link,omitempty"`
	Description         string                   `json:"description,omitempty"          xml:"description,omitempty"`
	Language            string                   `json:"language,omitempty"             xml:"language,omitempty"`
	Copyright           string                   `json:"copyright,omitempty"            xml:"copyright,omitempty"`
	ManagingEditor      string                   `json:"managingEditor,omitempty"       xml:"managingEditor,omitempty"`
	WebMaster           string                   `json:"webMaster,omitempty"            xml:"webMaster,omitempty"`
	PubDate             string                   `json:"pubDate,omitempty"              xml:"pubDate,omitempty"`
	PubDateParsed       *time.Time               `json:"pubDateParsed,omitempty"`
	LastBuildDate       string                   `json:"lastBuildDate,omitempty"        xml:"lastBuildDate,omitempty"`
	LastBuildDateParsed *time.Time               `json:"lastBuildDateParsed,omitempty"`
	Categories          []*Category              `json:"categories,omitempty"           xml:"categories,omitempty"`
	Generator           string                   `json:"generator,omitempty"            xml:"generator,omitempty"`
	Docs                string                   `json:"docs,omitempty"                 xml:"docs,omitempty"`
	TTL                 string                   `json:"ttl,omitempty"                  xml:"ttl,omitempty"`
	Image               *Image                   `json:"image,omitempty"                xml:"image,omitempty"`
	Rating              string                   `json:"rating,omitempty"               xml:"rating,omitempty"`
	SkipHours           []string                 `json:"skipHours,omitempty"            xml:"skipHours,omitempty"`
	SkipDays            []string                 `json:"skipDays,omitempty"             xml:"skipDays,omitempty"`
	Cloud               *Cloud                   `json:"cloud,omitempty"`
	TextInput           *TextInput               `json:"textInput,omitempty"`
	DublinCoreExt       *ext.DublinCoreExtension `json:"dcExt,omitempty"`
	ITunesExt           *ext.ITunesFeedExtension `json:"itunesExt,omitempty"`
	Extensions          ext.Extensions           `json:"extensions,omitempty"           xml:"-"`
	Items               []*Item                  `json:"items"`
	Version             string                   `json:"version"`
}

// Marshal the serialized xml for the parsed rss feed
func (f Feed) Marshal() ([]byte, error) {
	output, err := xml.Marshal(f)
	if err != nil {
		return []byte{}, err
	}
	return []byte(xml.Header + string(output)), nil
}

// MarshalIndent the serialized xml for the parsed rss feed
func (f Feed) MarshalIndent(prefix, indent string) ([]byte, error) {
	output, err := xml.MarshalIndent(f, prefix, indent)
	if err != nil {
		return []byte{}, err
	}
	return []byte(xml.Header + string(output)), nil
}

// MarshalXML is a custom xml marshaller function as the xml created from a feed has a unique
// format we must capture
func (f Feed) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{
		Name: xml.Name{Local: f.RootName},
		Attr: f.RootAttrs,
	})
	e.EncodeToken(xml.StartElement{
		Name: xml.Name{Local: "channel"},
	})

	encode(e, "title", f.Title)
	encode(e, "link", f.Link)
	encode(e, "description", f.Description)
	encode(e, "language", f.Language)
	encode(e, "copyright", f.Copyright)
	encode(e, "managingEditor", f.ManagingEditor)
	encode(e, "webMaster", f.WebMaster)
	encode(e, "pubDate", f.PubDate)
	encode(e, "lastBuildDate", f.LastBuildDate)
	encode(e, "category", f.Categories)
	encode(e, "generator", f.Generator)
	encode(e, "docs", f.Docs)
	encode(e, "ttl", f.TTL)
	encode(e, "image", f.Image)
	encode(e, "rating", f.Rating)
	encodeStringArray(e, "skipHours", "hour", f.SkipHours)
	encodeStringArray(e, "skipDays", "day", f.SkipDays)
	encode(e, "textinput", f.TextInput)
	encode(e, "cloud", f.Cloud)

	if f.ITunesExt != nil {
		f.ITunesExt.Encode(e)
		encode(e, "itunes:title", f.Title)
	}

	if f.DublinCoreExt != nil {
		f.DublinCoreExt.Encode(e)
	}

	encode(e, "items", f.Items)

	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "channel"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: f.RootName}})
	return nil
}

func (f Feed) String() string {
	json, _ := json.MarshalIndent(f, "", "    ")
	return string(json)
}

// Item is an RSS Item
type Item struct {
	XMLName xml.Name `xml:"item"`

	Title         string                   `json:"title,omitempty"          xml:"title,omitempty"`
	Link          string                   `json:"link,omitempty"           xml:"link,omitempty"`
	Description   string                   `json:"description,omitempty"    xml:"description,omitempty"`
	Content       string                   `json:"content,omitempty"        xml:"content,omitempty"`
	Author        string                   `json:"author,omitempty"         xml:"author,omitempty"`
	Categories    []*Category              `json:"categories,omitempty"`
	Comments      string                   `json:"comments,omitempty"       xml:"comments,omitempty"`
	Enclosure     *Enclosure               `json:"enclosure,omitempty"`
	GUID          *GUID                    `json:"guid,omitempty"`
	PubDate       string                   `json:"pubDate,omitempty"        xml:"pubDate,omitempty"`
	PubDateParsed *time.Time               `json:"pubDateParsed,omitempty"  xml:"-"`
	Source        *Source                  `json:"source,omitempty"`
	DublinCoreExt *ext.DublinCoreExtension `json:"dcExt,omitempty"`
	ITunesExt     *ext.ITunesItemExtension `json:"itunesExt,omitempty"`
	Extensions    ext.Extensions           `json:"extensions,omitempty"     xml:"-"`
	Custom        map[string]string        `json:"custom,omitempty"`
}

// MarshalXML is a custom xml marshaller for an item to allow for the itunes extension to
// be flattened into the item serialization.
func (i Item) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "item"}})

	encode(e, "title", i.Title)
	encode(e, "link", i.Link)
	encode(e, "description", i.Description)
	encode(e, "content", i.Content)
	encode(e, "author", i.Author)
	encode(e, "category", i.Categories)
	encode(e, "comments", i.Comments)
	encode(e, "enclosure", i.Enclosure)
	encode(e, "guid", i.GUID)
	encode(e, "pubDate", i.PubDate)
	encode(e, "source", i.Source)
	for k, v := range i.Custom {
		encode(e, k, v)
	}

	if i.ITunesExt != nil {
		i.ITunesExt.Encode(e)
		encode(e, "itunes:title", i.Title)
	}

	if i.DublinCoreExt != nil {
		i.DublinCoreExt.Encode(e)
	}

	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "item"}})
	return nil
}

// Image is an image that represents the feed
type Image struct {
	XMLName xml.Name `xml:"image"`

	URL         string `json:"url,omitempty"         xml:"url,omitempty"`
	Link        string `json:"link,omitempty"        xml:"link,omitempty"`
	Title       string `json:"title,omitempty"       xml:"title,omitempty"`
	Width       string `json:"width,omitempty"       xml:"width,omitempty"`
	Height      string `json:"height,omitempty"      xml:"height,omitempty"`
	Description string `json:"description,omitempty" xml:"description,omitempty"`
}

// Enclosure is a media object that is attached to
// the item
type Enclosure struct {
	XMLName xml.Name `xml:"enclosure"`

	URL    string `json:"url,omitempty"    xml:"url,attr"`
	Length string `json:"length,omitempty" xml:"length,attr"`
	Type   string `json:"type,omitempty"   xml:"type,attr"`
}

// GUID is a unique identifier for an item
type GUID struct {
	XMLName xml.Name `xml:"guid"`

	Value       string `json:"value,omitempty"       xml:",chardata"`
	IsPermalink string `json:"isPermaLink,omitempty" xml:"isPermaLink,attr"`
}

// Source contains feed information for another
// feed if a given item came from that feed
type Source struct {
	XMLName xml.Name `xml:"source"`

	Title string `json:"title,omitempty" xml:",chardata"`
	URL   string `json:"url,omitempty"   xml:"url,attr"`
}

// Category is category metadata for Feeds and Entries
type Category struct {
	XMLName xml.Name `xml:"category"`

	Domain string `json:"domain,omitempty" xml:"domain,attr,omitempty"`
	Value  string `json:"value,omitempty"  xml:",chardata"`
}

// TextInput specifies a text input box that
// can be displayed with the channel
type TextInput struct {
	XMLName xml.Name `xml:"textinput"`

	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Link        string `json:"link,omitempty"`
}

// Cloud allows processes to register with a
// cloud to be notified of updates to the channel,
// implementing a lightweight publish-subscribe protocol
// for RSS feeds
type Cloud struct {
	XMLName xml.Name `xml:"cloud"`

	Domain            string `json:"domain,omitempty"            xml:"domain,attr"`
	Port              string `json:"port,omitempty"              xml:"port,attr"`
	Path              string `json:"path,omitempty"              xml:"path,attr"`
	RegisterProcedure string `json:"registerProcedure,omitempty" xml:"registerProcedure,attr"`
	Protocol          string `json:"protocol,omitempty"          xml:"protocol,attr"`
}

func encode(e *xml.Encoder, name string, val interface{}, attrs ...xml.Attr) error {
	if val == nil {
		return nil
	}

	if sval, ok := val.(string); ok && len(sval) == 0 {
		return nil
	}

	return e.EncodeElement(val, xml.StartElement{Name: xml.Name{Local: name}, Attr: attrs})
}

func encodeStringArray(e *xml.Encoder, name, itemName string, val []string, attrs ...xml.Attr) error {
	if val == nil {
		return nil
	}

	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: name}})
	for _, v := range val {
		encode(e, itemName, v)
	}
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: name}})
	return nil
}
