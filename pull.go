// Currently runs into failed with object not found

package gosimplegit

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

/* Pull pulls the changes from the repository.
 *
 * Returns error if something goes wrong.
 */
func (c *Repository) Pull() error {
	w, err := c.repo.Worktree()
	if err != nil {
		return fmt.Errorf("Fetching worktree when pulling failed with %s", err.Error())
	}

	err = w.PullContext(c.ctx, &git.PullOptions{
		Progress: os.Stdout,
		Auth:     c.auth,
	})
	if err != nil {
		return fmt.Errorf("Pulling from repo failed with %+v", err)
	}
	return nil
}
