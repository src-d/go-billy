package billy

import "os"

// Exists returns true if the path exists in the filesystem. False, otherwise.
// If there is an I/O error that prevents checking the existence of the file
// an error is returned.
func Exists(fs Filesystem, path string) (bool, error) {
	_, err := fs.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}

	return err == nil, err
}
