package gitops

import (
	"errors"
	"os"
	"testing"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/stretchr/testify/assert"
)

// StubGitClient is a stub implementation of GitClient.
type StubGitClient struct {
	CloneFunc func(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error)
}

// Clone calls the stubbed CloneFunc.
func (s *StubGitClient) Clone(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error) {
	return s.CloneFunc(path, isBare, o)
}

func TestCloneRepo_Success(t *testing.T) {
	// Arrange
	stubClient := &StubGitClient{
		CloneFunc: func(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error) {
			return &git.Repository{}, nil
		},
	}

	options := CloneOptions{
		RepoURL:     "https://github.com/some/repo.git",
		Destination: "./tmp/destination",
		Auth:        &http.BasicAuth{Username: "user", Password: "pass"},
	}

	// Ensure the destination does not exist
	os.RemoveAll(options.Destination)

	// Act
	err := CloneRepo(stubClient, options)

	// Assert
	assert.NoError(t, err)
}

func TestCloneRepo_DestinationExists(t *testing.T) {
	// Arrange
	stubClient := &StubGitClient{}
	options := CloneOptions{
		RepoURL:     "https://github.com/some/repo.git",
		Destination: "./tmp/destination",
		Auth:        &http.BasicAuth{Username: "user", Password: "pass"},
	}

	// Create the destination directory
	err := os.MkdirAll(options.Destination, os.ModePerm)
	assert.NoError(t, err)

	// Act
	err = CloneRepo(stubClient, options)

	// Assert
	assert.Error(t, err)
	assert.EqualError(t, err, "destination directory ./tmp/destination already exists")

	// Cleanup
	os.RemoveAll(options.Destination)
}

func TestCloneRepo_CloneFailure(t *testing.T) {
	// Arrange
	stubClient := &StubGitClient{
		CloneFunc: func(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error) {
			return nil, errors.New("clone error")
		},
	}

	options := CloneOptions{
		RepoURL:     "https://github.com/some/repo.git",
		Destination: "./tmp/destination",
		Auth:        &http.BasicAuth{Username: "user", Password: "pass"},
	}

	// Ensure the destination does not exist
	os.RemoveAll(options.Destination)

	// Act
	err := CloneRepo(stubClient, options)

	// Assert
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to clone repository: clone error")
}
