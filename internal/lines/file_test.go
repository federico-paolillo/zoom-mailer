package lines

import (
	"io"
	"os"
	"testing"
)

func TestScanFile(t *testing.T) {

	fileHandle, openError := os.Open("testdata/lorem.txt")

	if openError != nil {

		t.Fatal("could not open test file testdata/lorem.txt")

	}

	defer fileHandle.Close()

	scanner := NewFileLineScanner(fileHandle)

	line1, _ := scanner.Scan()

	if line1 != "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua." {

		t.Fatal("line 1 has an unexpected value")

	}

	line2, _ := scanner.Scan()

	if line2 != "Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat." {

		t.Fatal("line 2 has an unexpected value")

	}

	line3, _ := scanner.Scan()

	if line3 != "Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur." {

		t.Fatal("line 3 has an unexpected value")

	}

	line4, _ := scanner.Scan()

	if line4 != "Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum." {

		t.Fatal("line 3 has an unexpected value")

	}

	_, eofError := scanner.Scan()

	if eofError != io.EOF {

		t.Fatal("expected End Of File")

	}

}
