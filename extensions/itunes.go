package ext

import "encoding/xml"

// ITunesFeedExtension is a set of extension
// fields for RSS feeds.
type ITunesFeedExtension struct {
	XMLName    xml.Name          `xml:"-"`
	Author     string            `json:"author,omitempty"     xml:"itunes:author,omitempty"`
	Block      string            `json:"block,omitempty"      xml:"itunes:block,omitempty"`
	Categories []*ITunesCategory `json:"categories,omitempty" xml:"itunes:category,omitempty"`
	Explicit   string            `json:"explicit,omitempty"   xml:"itunes:explicit,omitempty"`
	Keywords   string            `json:"keywords,omitempty"   xml:"itunes:keywords,omitempty"`
	Owner      *ITunesOwner      `json:"owner,omitempty"      xml:"itunes:owner,omitempty"`
	Subtitle   string            `json:"subtitle,omitempty"   xml:"itunes:subtitle,omitempty"`
	Summary    string            `json:"summary,omitempty"    xml:"itunes:summary,omitempty"`
	Image      string            `json:"image,omitempty"      xml:"itunes:image,omitempty"`
	Complete   string            `json:"complete,omitempty"   xml:"itunes:complete,omitempty"`
	NewFeedURL string            `json:"newFeedUrl,omitempty" xml:"itunes:newFeedUrl,omitempty"`
	Type       string            `json:"type,omitempty"       xml:"itunes:type,omitempty"`
}

// Encode will encode the itunes extension in the provided xml encoder
func (itunes ITunesFeedExtension) Encode(e *xml.Encoder) error {
	encode(e, "itunes:author", itunes.Author)
	encode(e, "itunes:block", itunes.Block)
	encode(e, "itunes:category", itunes.Categories)
	encode(e, "itunes:explicit", itunes.Explicit)
	encode(e, "itunes:keywords", itunes.Keywords)
	encode(e, "itunes:owner", itunes.Owner)
	encode(e, "itunes:subtitle", itunes.Subtitle)
	encode(e, "itunes:summary", itunes.Summary)
	encode(e, "itunes:image", "", xml.Attr{Name: xml.Name{Local: "href"}, Value: itunes.Image})
	encode(e, "itunes:complete", itunes.Complete)
	encode(e, "itunes:newFeedUrl", itunes.NewFeedURL)
	encode(e, "itunes:type", itunes.Type)

	return nil
}

// ITunesItemExtension is a set of extension
// fields for RSS items.
type ITunesItemExtension struct {
	Author            string `json:"author,omitempty"            xml:"itunes:author,omitempty"`
	Block             string `json:"block,omitempty"             xml:"itunes:block,omitempty"`
	Duration          string `json:"duration,omitempty"          xml:"itunes:duration,omitempty"`
	Explicit          string `json:"explicit,omitempty"          xml:"itunes:explicit,omitempty"`
	Keywords          string `json:"keywords,omitempty"          xml:"itunes:keywords,omitempty"`
	Subtitle          string `json:"subtitle,omitempty"          xml:"itunes:subtitle,omitempty"`
	Summary           string `json:"summary,omitempty"           xml:"itunes:summary,omitempty"`
	Image             string `json:"image,omitempty"             xml:"itunes:image,omitempty"`
	IsClosedCaptioned string `json:"isClosedCaptioned,omitempty" xml:"itunes:isClosedCaptioned,omitempty"`
	Episode           string `json:"episode,omitempty"           xml:"itunes:episode,omitempty"`
	Season            string `json:"season,omitempty"            xml:"itunes:season,omitempty"`
	Order             string `json:"order,omitempty"             xml:"itunes:order,omitempty"`
	EpisodeType       string `json:"episodeType,omitempty"       xml:"itunes:episodeType,omitempty"`
}

// Encode will encode the itunes item in the provided xml encoder
func (itunes ITunesItemExtension) Encode(e *xml.Encoder) error {
	encode(e, "itunes:author", itunes.Author)
	encode(e, "itunes:block", itunes.Block)
	encode(e, "itunes:duration", itunes.Duration)
	encode(e, "itunes:explicit", itunes.Explicit)
	encode(e, "itunes:keywords", itunes.Keywords)
	encode(e, "itunes:subtitle", itunes.Subtitle)
	encode(e, "itunes:summary", itunes.Summary)
	encode(e, "itunes:image", itunes.Image)
	encode(e, "itunes:isClosedCaptioned", itunes.IsClosedCaptioned)
	encode(e, "itunes:episode", itunes.Episode)
	encode(e, "itunes:season", itunes.Season)
	encode(e, "itunes:order", itunes.Order)
	encode(e, "itunes:episodeType", itunes.EpisodeType)

	return nil
}

// ITunesCategory is a category element for itunes feeds.
type ITunesCategory struct {
	XMLName xml.Name `xml:"itunes:category"`

	Text        string          `json:"text,omitempty"        xml:"text,attr"`
	Subcategory *ITunesCategory `json:"subcategory,omitempty" xml:"itunes:category,omitempty"`
}

// ITunesOwner is the owner of a particular itunes feed.
type ITunesOwner struct {
	XMLName xml.Name `xml:"itunes:owner"`

	Email string `json:"email,omitempty" xml:"itunes:email,omitempty"`
	Name  string `json:"name,omitempty"  xml:"itunes:name,omitempty"`
}

// NewITunesFeedExtension creates an ITunesFeedExtension given an
// extension map for the "itunes" key.
func NewITunesFeedExtension(extensions map[string][]Extension) *ITunesFeedExtension {
	feed := &ITunesFeedExtension{}
	feed.Author = parseTextExtension("author", extensions)
	feed.Block = parseTextExtension("block", extensions)
	feed.Explicit = parseTextExtension("explicit", extensions)
	feed.Keywords = parseTextExtension("keywords", extensions)
	feed.Subtitle = parseTextExtension("subtitle", extensions)
	feed.Summary = parseTextExtension("summary", extensions)
	feed.Image = parseImage(extensions)
	feed.Complete = parseTextExtension("complete", extensions)
	feed.NewFeedURL = parseTextExtension("new-feed-url", extensions)
	feed.Categories = parseCategories(extensions)
	feed.Owner = parseOwner(extensions)
	feed.Type = parseTextExtension("type", extensions)
	return feed
}

// NewITunesItemExtension creates an ITunesItemExtension given an
// extension map for the "itunes" key.
func NewITunesItemExtension(extensions map[string][]Extension) *ITunesItemExtension {
	entry := &ITunesItemExtension{}
	entry.Author = parseTextExtension("author", extensions)
	entry.Block = parseTextExtension("block", extensions)
	entry.Duration = parseTextExtension("duration", extensions)
	entry.Explicit = parseTextExtension("explicit", extensions)
	entry.Subtitle = parseTextExtension("subtitle", extensions)
	entry.Summary = parseTextExtension("summary", extensions)
	entry.Keywords = parseTextExtension("keywords", extensions)
	entry.Image = parseImage(extensions)
	entry.IsClosedCaptioned = parseTextExtension("isClosedCaptioned", extensions)
	entry.Episode = parseTextExtension("episode", extensions)
	entry.Season = parseTextExtension("season", extensions)
	entry.Order = parseTextExtension("order", extensions)
	entry.EpisodeType = parseTextExtension("episodeType", extensions)
	return entry
}

func parseImage(extensions map[string][]Extension) (image string) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["image"]
	if !ok || len(matches) == 0 {
		return
	}

	image = matches[0].Attrs["href"]
	return
}

func parseOwner(extensions map[string][]Extension) (owner *ITunesOwner) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["owner"]
	if !ok || len(matches) == 0 {
		return
	}

	owner = &ITunesOwner{}
	if name, ok := matches[0].Children["name"]; ok {
		owner.Name = name[0].Value
	}
	if email, ok := matches[0].Children["email"]; ok {
		owner.Email = email[0].Value
	}
	return
}

func parseCategories(extensions map[string][]Extension) (categories []*ITunesCategory) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["category"]
	if !ok || len(matches) == 0 {
		return
	}

	categories = []*ITunesCategory{}
	for _, cat := range matches {
		c := &ITunesCategory{}
		if text, ok := cat.Attrs["text"]; ok {
			c.Text = text
		}

		if subs, ok := cat.Children["category"]; ok {
			s := &ITunesCategory{}
			if text, ok := subs[0].Attrs["text"]; ok {
				s.Text = text
			}
			c.Subcategory = s
		}
		categories = append(categories, c)
	}
	return
}
