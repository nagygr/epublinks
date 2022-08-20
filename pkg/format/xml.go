package format

/*
 * Using code snippets from:
 * https://github.com/nguyenthenguyen/docx/
 * (c) Nguyen The Nguyen
 */

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/nagygr/epublinks/pkg/archive"
	"io"
	"io/ioutil"
	"strings"
)

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
