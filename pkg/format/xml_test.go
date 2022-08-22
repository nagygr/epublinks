package format

import (
	"testing"
)

func TestLocalEpubFile(t *testing.T) {
	urls, err := EpubLinksFromFile("../../attic/sample_en.epub")

	if err != nil {
		t.Fatalf(
			"Error extracting URLs from EPUB: %s",
			err.Error(),
		)
	}

	expectedLen := 274
	if actualLen := len(urls); actualLen != expectedLen {
		t.Fatalf(
			"Expected to find %d links, found %d",
			expectedLen, actualLen,
		)
	}

	expectedUrl := "https://en.wikipedia.org/wiki/Information"
	if actualUrl := urls[0]; actualUrl != expectedUrl {
		t.Fatalf(
			"Expected to find %s as first URL but found: %s",
			expectedUrl, actualUrl,
		)
	}
}
