package lines

import (
	"bufio"
	"io"
	"os"
)

type FileLineScanner struct {
	scanner *bufio.Scanner
}

func (fileLineScanner *FileLineScanner) Scan() (string, error) {

	bufioScanner := fileLineScanner.scanner

	if bufioScanner.Scan() {

		lineScanned := bufioScanner.Text()

		return lineScanned, nil

	}

	bufioErr := bufioScanner.Err()

	if bufioErr == nil {

		return "", io.EOF

	}

	return "", bufioErr

}

func NewFileLineScanner(fileToScanHandle *os.File) *FileLineScanner {

	bufioScanner := bufio.NewScanner(fileToScanHandle)

	return &FileLineScanner{scanner: bufioScanner}

}
