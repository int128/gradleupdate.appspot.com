package domain

import (
	"fmt"
	"strings"
)

// RepositoryID points to a repository.
type RepositoryID struct {
	Owner string
	Name  string
}

func (r RepositoryID) FullName() string {
	return r.Owner + "/" + r.Name
}

func (r RepositoryID) String() string {
	return r.FullName()
}

func (r RepositoryID) Branch(name string) BranchID {
	return BranchID{Repository: r, Name: name}
}

// Repository represents a GitHub repository.
type Repository struct {
	ID            RepositoryID
	DefaultBranch BranchID
	Description   string
	AvatarURL     string
}

func (r Repository) String() string {
	return r.ID.String()
}

// RepositoryURL represents URL for a GitHub repository.
type RepositoryURL string

// Parse returns owner and repo for the repository.
func (url RepositoryURL) Parse() *RepositoryID {
	s := strings.Split(string(url), "/")
	if len(s) < 2 {
		return nil
	}
	return &RepositoryID{s[len(s)-2], s[len(s)-1]}
}

// FileContent represents content of a file.
type FileContent []byte

func (fc FileContent) String() string {
	return string(fc)
}

// File represents a file in a commit.
type File struct {
	Path    string
	Content FileContent
}

// PullRequestID points to a pull request.
type PullRequestID struct {
	Repository RepositoryID
	Number     int
}

func (p PullRequestID) String() string {
	return fmt.Sprintf("%s/pulls#%d", p.Repository, p.Number)
}

// PullRequest represents a pull request.
type PullRequest struct {
	ID         PullRequestID
	BaseBranch BranchID
	HeadBranch BranchID
	Title      string
	Body       string
}

func (p PullRequest) String() string {
	return p.ID.String()
}

// BranchID points to a branch in a repository.
type BranchID struct {
	Repository RepositoryID
	Name       string
}

func (b BranchID) String() string {
	return b.Repository.String() + ":" + b.Name
}

func (b BranchID) Ref() string {
	return "refs/heads/" + b.Name
}

// Branch represents a branch in a repository.
type Branch struct {
	ID     BranchID
	Commit Commit
}

func (b Branch) String() string {
	return b.ID.String()
}

// CommitID points to a branch in a repository.
type CommitID struct {
	Repository RepositoryID
	SHA        CommitSHA
}

type CommitSHA string

func (sha CommitSHA) String() string {
	return string(sha)
}

// Commit represents a commit in a repository.
type Commit struct {
	ID      CommitID
	Parents []CommitID
	Tree    TreeID
}

func (c Commit) IsBasedOn(base CommitID) bool {
	if len(c.Parents) != 1 {
		return false
	}
	parent := c.Parents[0]
	return parent.SHA == base.SHA
}

type TreeSHA string

func (sha TreeSHA) String() string {
	return string(sha)
}

// TreeID points to a tree in a repository.
type TreeID struct {
	Repository RepositoryID
	SHA        TreeSHA
}