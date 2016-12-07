package memory_test

import (
	"srcd.works/billy"
	"srcd.works/billy/memory"
)

func Example() {
	var fs billy.Filesystem
	fs = memory.New()

	doSomethingWith(fs)
}

func doSomethingWith(fs billy.Filesystem) {
	// Intentionally left blank
}
