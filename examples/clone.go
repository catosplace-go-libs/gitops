package main

import (
	"fmt"
	"log"

	"github.com/catosplace-go-libs/gitops/pkg/gitops"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func main() {
	// Define the options for cloning the repository
	options := gitops.CloneOptions{
		RepoURL:     "https://github.com/some-user/some-repo.git",
		Destination: "./some-repo-clone",
		Auth: &http.BasicAuth{
			Username: "your-username", // Replace with your GitHub username
			Password: "your-token",    // Replace with your GitHub token
		},
	}

	// Use the default Git client implementation
	gitClient := &gitops.GitClientImpl{}

	// Clone the repository
	cloneOptions := gitops.CloneOptions{
		RepoURL:     "YOUR_GITHUB_REPO_URL",
		Destination: "./template",
		Auth:        auth,
	}

	// Clone the repository
	err := gitops.CloneRepo(gitClient, options)
	if err != nil {
		log.Fatalf("Error cloning repository: %v", err)
	}

	fmt.Println("Repository cloned successfully!")
}
