package format

/*
 * Using code snippets from:
 * https://github.com/nguyenthenguyen/docx/
 * (c) Nguyen The Nguyen
 */

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/nagygr/epublinks/pkg/archive"
)

// EpubLinksFromFile extracts the URLs from an EPUB file given
// by its path. The return values are the slice of URLs and any
// error that happened during the extraction.
func EpubLinksFromFile(path string) ([]string, error) {
	zipFile, err := archive.NewZipFile(path)

	if err != nil {
		return nil, fmt.Errorf(
			"Error while opening EPUB file from path: %w",
			err,
		)
	}

	return ExtractLinksFromZipFile(zipFile, "OEBPS/sections/section")
}

// EpubLinksFromUrl extracts the URLs from an EPUB file given
// with an URL. The return values are the slice of URLs and any
// error that happened during the extraction.
func EpubLinksFromUrl(url string) ([]string, error) {
	zipFile, err := archive.NewZipFileFromUrl(url)

	if err != nil {
		return nil, fmt.Errorf(
			"Error while opening EPUB file from url: %w",
			err,
		)
	}

	return ExtractLinksFromZipFile(zipFile, "OEBPS/sections/section")
}

func ExtractLinksFromZipFile(zipFile archive.ZipData, nameFragment string) (
	[]string, error,
) {
	texts, err := ReadXmls(zipFile, nameFragment)

	if err != nil {
		return nil, fmt.Errorf(
			"Error while parsing EPUB structure: %w",
			err,
		)
	}

	urls, err := UrlsFromXmls(texts)

	if err != nil {
		return nil, fmt.Errorf(
			"Error while extracting URLs from EPUB: %w",
			err,
		)
	}

	return urls, nil
}

func ReadXmls(zipReader archive.ZipData, nameFragment string) (xmlTexts []string, err error) {
	xmlFiles, err := zipReader.FilesByName(nameFragment)

	if err != nil {
		return []string{}, err
	}

	xmlTexts, err = ReadTextFromXmls(xmlFiles)
	if err != nil {
		return []string{}, err
	}

	return
}

func ReadTextFromXmls(xmlFiles []*zip.File) ([]string, error) {
	xmlText := []string{}

	for _, element := range xmlFiles {
		documentReader, err := element.Open()
		if err != nil {
			return []string{}, err
		}

		text, err := XmlFileToString(documentReader)
		if err != nil {
			return []string{}, err
		}

		xmlText = append(xmlText, text)
	}

	return xmlText, nil
}

func XmlFileToString(reader io.Reader) (string, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf(
			"Error while extracting text: %w",
			err,
		)
	}
	return string(b), nil
}

func UrlsFromXml(xmlText string) ([]string, error) {
	var (
		contents = strings.NewReader(xmlText)
		decoder  = xml.NewDecoder(contents)
		results  = []string{}
	)

	for {
		token, err := decoder.Token()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf(
				"Error while parsing xml file: %w",
				err,
			)
		}

		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "a" {
				for _, v := range t.Attr {
					if v.Name.Local == "href" {
						results = append(results, v.Value)
					}
				}
			}
		default:
		}
	}

	return results, nil
}

func UrlsFromXmls(xmlTexts []string) ([]string, error) {
	results := []string{}

	for _, xml := range xmlTexts {
		urls, err := UrlsFromXml(xml)

		if err != nil {
			return nil, fmt.Errorf(
				"Error while extracting URLs from XML: %w",
				err,
			)
		}

		results = append(results, urls...)
	}

	return results, nil
}
