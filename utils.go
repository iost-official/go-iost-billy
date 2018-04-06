package billy

import "os"

func Exists(fs Filesystem, path string) (bool, error) {
	_, err := fs.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}

	return err == nil, err
}