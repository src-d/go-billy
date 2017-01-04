package billy_test

import (
	"errors"
	"testing"

	"srcd.works/go-billy.v1"
	"srcd.works/go-billy.v1/memory"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type UtilsSuite struct{}

var _ = Suite(&UtilsSuite{})

func (s *UtilsSuite) TestExistsFalse(c *C) {
	fs := memory.New()

	e, err := billy.Exists(fs, "non-existent")
	c.Assert(err, IsNil)
	c.Assert(e, Equals, false)
}

func (s *UtilsSuite) TestExistsTrue(c *C) {
	fs := memory.New()
	f, err := fs.Create("existent")
	c.Assert(err, IsNil)
	c.Assert(f.Close(), IsNil)

	e, err := billy.Exists(fs, "existent")
	c.Assert(err, IsNil)
	c.Assert(e, Equals, true)
}

func (s *UtilsSuite) TestExistsError(c *C) {
	fs := &errorStatFs{memory.New()}

	e, err := billy.Exists(fs, "existent")
	c.Assert(err, ErrorMatches, "test error")
	c.Assert(e, Equals, false)
}

type errorStatFs struct {
	billy.Filesystem
}

func (fs *errorStatFs) Stat(path string) (billy.FileInfo, error) {
	return nil, errors.New("test error")
}
