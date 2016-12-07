package memory

import (
	"gopkg.in/src-d/go-git.v4/utils/fs/memory"

	"srcd.works/billy"
)

func New() billy.Filesystem {
	return memory.New()
}
