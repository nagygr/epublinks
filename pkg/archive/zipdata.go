// Package archive contains types related to zip handling.
package archive

/*
 * Using code snippets from:
 * https://github.com/nguyenthenguyen/docx/
 * (c) Nguyen The Nguyen
 */

import (
	"archive/zip"
)

// ZipData defines the common interface for different zip-handling types.
type ZipData interface {
	Files() []*zip.File
	close() error

	FileByName(name string) (file *zip.File, err error)
	FilesByName(substring string) (files []*zip.File, err error)
}
