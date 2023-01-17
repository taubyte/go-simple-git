package gosimplegit

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

/*
ListBranches will return a list of branches for the repository

fetch true will search remote origin to gather all branches
fetch false will search .git/config to gather branches
*/
func (c *Repository) ListBranches(fetch bool) (branches []string, fetchErr error, err error) {
	if fetch == true {
		return c.fetchAndListBranches()
	}

	branchRef, err := c.repo.Branches()
	if err != nil {
		return nil, fetchErr, fmt.Errorf("listing branches for repository: `%s` failed with: %s", c.url, err)
	}

	branches = make([]string, 0)
	err = branchRef.ForEach(func(r *plumbing.Reference) error {
		branches = append(branches, r.Name().Short())
		return nil
	})
	if err != nil {
		return nil, fetchErr, fmt.Errorf("branchRef.ForEach() for repository: `%s` failed with: %s", c.url, err)
	}

	return branches, fetchErr, nil
}

func (c *Repository) fetchAndListBranches() (branches []string, fetchErr error, err error) {
	fetchErr = c.Fetch()

	rem, err := c.repo.Remote("origin")
	if err != nil {
		return nil, fetchErr, fmt.Errorf("getting remote origin for repository: `%s` failed with: %s", c.url, err)
	}

	remoteLister, err := rem.List(&git.ListOptions{
		Auth: c.auth,
	})
	if err != nil {
		return nil, fetchErr, fmt.Errorf("listing origin references for repository: `%s` failed with: %s", c.url, err)
	}

	branches = make([]string, 0)
	for _, branch := range remoteLister {
		if branch.Name().IsBranch() == true {
			branches = append(branches, branch.Name().Short())
		}
	}

	return
}
