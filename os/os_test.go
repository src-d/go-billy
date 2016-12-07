package os_test

import (
	"srcd.works/billy"
	"srcd.works/billy/os"
)

func Example() {
	var fs billy.Filesystem
	fs = os.New("/bin")

	doSomethingWith(fs)
}

func doSomethingWith(fs billy.Filesystem) {
	// Intentionally left blank
}
