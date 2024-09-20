package gitops

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// CloneOptions holds the options for cloning a git repository.
type CloneOptions struct {
	RepoURL     string
	Destination string
	Auth        *http.BasicAuth
}

// CloneRepo clones a git repository to a specified location using the user's git credentials.
func CloneRepo(options CloneOptions) error {
	// Ensure the destination directory does not already exist
	if _, err := os.Stat(options.Destination); !os.IsNotExist(err) {
		return fmt.Errorf("destination directory %s already exists", options.Destination)
	}

	// Clone the repository
	_, err := git.PlainClone(options.Destination, false, &git.CloneOptions{
		URL:      options.RepoURL,
		Auth:     options.Auth,
		Progress: os.Stdout,
	})
	if err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}

	return nil
}
