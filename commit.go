package gosimplegit

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

/* Commit commits the changes to the repository.
 *
 * message: The message to be used for the commit.
 * files: The files to be committed.
 *
 * Returns error if something goes wrong.
 */
func (c *Repository) Commit(message string, files string) error {
	var err error

	w, err := c.repo.Worktree()
	if err != nil {
		return fmt.Errorf("Fetching work tree failed with %s", err.Error())
	}

	_, err = w.Add(files)
	if err != nil {
		return fmt.Errorf("Adding files failed with %s", err.Error())
	}

	commit, err := w.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  c.user.name,
			Email: c.user.email,
			When:  time.Now(),
		},
	})

	_, err = c.repo.CommitObject(commit)
	if err != nil {
		return fmt.Errorf("Commiting changes failed with %s", err.Error())
	}

	return nil
}
