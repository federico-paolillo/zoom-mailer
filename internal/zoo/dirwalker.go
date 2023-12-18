package zoo

import (
	"io/fs"
	"path/filepath"
)

func newWalkFn(filesSeen *[]string) fs.WalkDirFunc {

	hasDescended := false

	return func(path string, d fs.DirEntry, err error) error {

		if err != nil {

			return err

		}

		if d.IsDir() {

			if hasDescended {

				return fs.SkipDir

			} else {

				hasDescended = true // We do not want to visit sub-folders

			}

		} else {

			*filesSeen = append(*filesSeen, path)

		}

		return nil
	}

}

func gatherAvailabilityFiles(dirToScanPath string) (*[]string, error) {

	filePathsGathered := make([]string, 0)

	walkFn := newWalkFn(&filePathsGathered)

	err := filepath.WalkDir(dirToScanPath, walkFn)

	if err != nil {

		return nil, err

	}

	return &filePathsGathered, nil

}
