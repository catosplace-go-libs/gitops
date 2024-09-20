package main

import (
	"fmt"
	"log"

	"github.com/catosplace-go-libs/gitops/pkg/gitops"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func main() {

	accessToken := "YOUR_GITHUB_ACCESS_TOKEN"

	// Create an http.BasicAuth object with the token
	auth := &http.BasicAuth{
		Username: "YOUR_GITHUB_USERNAE",
		Password: accessToken,
	}

	// Clone the repository
	cloneOptions := gitops.CloneOptions{
		RepoURL:     "YOUR_GITHUB_REPO_URL",
		Destination: "./template",
		Auth:        auth,
	}

	err := gitops.CloneRepo(cloneOptions)
	if err != nil {
		log.Fatalf("Failed to clone repository: %v", err)
	}

	fmt.Println("Repository cloned successfully!")
}
