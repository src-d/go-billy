package os

import (
	"gopkg.in/src-d/go-git.v4/utils/fs/os"

	"srcd.works/billy"
)

func New(baseDir string) billy.Filesystem {
	return os.New(baseDir)
}
