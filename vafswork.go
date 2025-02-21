package vafswork

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// get full path of Go executable
func GetExePath() string {
	// get executable's working dir
	exe, _ := os.Executable()
	exePath := filepath.Dir(exe)
	return exePath
}

// Rotate files: keep <num> of most recent files and delete other
func RotateFilesByMtime(filesDir string, filesToKeep int) error {
	fileToDel := ""

	files, err := os.ReadDir(filesDir)
	if err != nil {
		return err
	}

	// sort file slice by modification time(asc)
	sort.Slice(
		files,
		func(i, j int) bool {
			fileI, _ := files[i].Info()
			fileJ, _ := files[j].Info()
			return fileI.ModTime().After(fileJ.ModTime())
		},
	)

	// delete files which index more than <filesToKeep> value
	for ind, file := range files {
		// skip dir
		if file.IsDir() {
			continue
		}

		// deleting old files
		if ind+1 > filesToKeep {
			fileToDel = fmt.Sprintf("%s/%s", filesDir, file.Name())

			// skip all errors of os.Remove(Path errors)
			os.Remove(fileToDel)
		}
	}

	return nil
}
