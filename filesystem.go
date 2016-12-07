// Interface to access different filesystems
package billy

import "gopkg.in/src-d/go-git.v4/utils/fs"

type Filesystem interface {
	fs.Filesystem
}
type File interface {
	fs.File
}
type FileInfo interface {
	fs.FileInfo
}
