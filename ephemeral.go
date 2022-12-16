package gosimplegit

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func (c *Repository) handle_ephemeral() (err error) {
	if strings.HasPrefix(c.root, "/") {
		return fmt.Errorf("Root: `%s` must be relative when using Temporary()", c.root)
	}

	c.workDir, err = ioutil.TempDir("", "repo-*")
	c.root = path.Join(c.workDir, c.root)
	if err != nil {
		return fmt.Errorf(
			"Creating new git client (url=%s) failed with %v",
			c.url,
			err,
		)
	}
	// make root dir if needed
	if len(c.workDir) > 0 {
		err = os.MkdirAll(c.root, 0755)
		if err != nil {
			return fmt.Errorf("Failed making root dir `%s` with %v", c.root, err)
		}
	} else {
		c.workDir = c.root
	}

	// make sure we cleanup
	go func(_c *Repository) {
		select {
		case <-_c.ctx.Done():
			if c.ephemeralNoDelete == false {
				os.RemoveAll(_c.workDir)
			}
		}
	}(c)

	return nil
}
