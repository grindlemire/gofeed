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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser_Parse(t *testing.T) {
	files, _ := filepath.Glob("../testdata/parser/rss/*.xml")
	for _, f := range files {
		testFile(t, f)
	}
}

func TestIsolateProblem(t *testing.T) {

	name := "../testdata/parser/rss/rss_channel_item_custom.xml"

	testFile(t, name)
}

// TODO: Examples
func TestGenerateXML(t *testing.T) {
	name := "./realistic_test.xml"

	raw, err := ioutil.ReadFile(name)
	assert.Nil(t, err)

	fmt.Printf("input:\n%s\n\n\n\n\n", raw)

	fp := &rss.Parser{}
	feed, err := fp.Parse(bytes.NewReader(raw))
	assert.Nil(t, err)

	output, err := feed.Marshal()
	assert.Nil(t, err)

	fmt.Printf("output:\n\n\n\n\n%s\n", output)
}

func testFile(t *testing.T, filename string) {
	base := filepath.Base(filename)
	name := strings.TrimSuffix(base, filepath.Ext(base))

	fmt.Printf("Testing %s... ", name)

	// Get actual source feed
	ff := fmt.Sprintf("../testdata/parser/rss/%s.xml", name)
	f, _ := ioutil.ReadFile(ff)

	// Parse actual feed
	fp := &rss.Parser{}
	actual, _ := fp.Parse(bytes.NewReader(f))
	newInput, err := actual.Marshal()
	assert.Nil(t, err)

	actual.RootName = ""
	actual.RootAttrs = nil

	// Get json encoded expected feed result
	ef := fmt.Sprintf("../testdata/parser/rss/%s.json", name)
	e, _ := ioutil.ReadFile(ef)

	// Unmarshal expected feed
	expected := &rss.Feed{}
	json.Unmarshal(e, &expected)

	if assert.Equal(t, expected, actual, "Feed file %s.xml did not match expected output %s.json", name, name) {
		fmt.Printf("OK...")
	} else {
		fmt.Printf("Failed\n")
	}

	fp = &rss.Parser{}
	newActual, err := fp.Parse(bytes.NewReader(newInput))
	require.Nil(t, err)
	newActual.RootAttrs = nil
	newActual.RootName = ""

	if assert.Equal(t, expected, newActual, "Remarshalled Feed file %s.xml did not match expected output %s.json", name, name) {
		fmt.Printf("Remarshalling OK\n")
	} else {
		fmt.Printf("Failed Remarshalling\n")
	}
}
