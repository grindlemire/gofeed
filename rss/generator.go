package rss

import "encoding/xml"

// Marshal the serialized xml for the parsed rss feed
func Marshal(f *Feed) ([]byte, error) {
	output, err := xml.MarshalIndent(f, "", "  ")
	if err != nil {
		return []byte{}, err
	}
	return []byte(xml.Header + string(output)), nil
}
