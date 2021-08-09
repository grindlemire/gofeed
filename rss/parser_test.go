package rss_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mmcdole/gofeed/rss"
	"github.com/stretchr/testify/require"
)

func TestParser_Parse(t *testing.T) {
	files, _ := filepath.Glob("../testdata/parser/rss/rss_channel_copyright*.xml")
	for _, f := range files {
		testFile(t, f)
	}
}

func TestIsolate(t *testing.T) {
	f := fmt.Sprintf("../testdata/parser/rss/%s.xml", "rss_channel_copyright_escaped_markup")
	testFile(t, f)
}

func testFile(t *testing.T, filename string) {
	base := filepath.Base(filename)
	name := strings.TrimSuffix(base, filepath.Ext(base))

	fmt.Printf("Testing %s...\n ", name)

	// Get actual source feed
	ff := fmt.Sprintf("../testdata/parser/rss/%s.xml", name)
	f, _ := ioutil.ReadFile(ff)

	// Parse actual feed
	fp := &rss.Parser{}
	actual, _ := fp.Parse(bytes.NewReader(f))
	newInput, err := actual.MarshalIndent("", "  ")
	require.Nil(t, err)

	actual.RootName = ""
	actual.RootAttrs = nil

	// Get json encoded expected feed result
	ef := fmt.Sprintf("../testdata/parser/rss/%s.json", name)
	e, _ := ioutil.ReadFile(ef)

	// Unmarshal expected feed
	expected := &rss.Feed{}
	json.Unmarshal(e, &expected)

	require.Equal(t, expected, actual, "Feed file %s.xml did not match expected output %s.json", name, name)

	fp = &rss.Parser{}
	newActual, err := fp.Parse(bytes.NewReader(newInput))
	require.Nil(t, err)
	newActual.RootAttrs = nil
	newActual.RootName = ""

	require.Equal(t, expected, newActual, "Remarshalled Feed file %s.xml did not match expected output %s.json", name, name)
}
