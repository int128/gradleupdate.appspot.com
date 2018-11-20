package registry

import (
	"context"

	"github.com/int128/gradleupdate/app/domain/repositories"
	"github.com/int128/gradleupdate/app/infrastructure"
	impl "github.com/int128/gradleupdate/app/infrastructure/repositories"
)

type Repositories interface {
	Repository(context.Context) repositories.Repository
	PullRequest(context.Context) repositories.PullRequest
	Branch(context.Context) repositories.Branch
	Commit(context.Context) repositories.Commit
}

type defaultRepositories struct{}

func (*defaultRepositories) Repository(ctx context.Context) repositories.Repository {
	return &impl.Repository{GitHub: infrastructure.GitHubClient(ctx)}
}

func (*defaultRepositories) PullRequest(ctx context.Context) repositories.PullRequest {
	return &impl.PullRequest{GitHub: infrastructure.GitHubClient(ctx)}
}

func (*defaultRepositories) Branch(ctx context.Context) repositories.Branch {
	return &impl.Branch{GitHub: infrastructure.GitHubClient(ctx)}
}

func (*defaultRepositories) Commit(ctx context.Context) repositories.Commit {
	return &impl.Commit{GitHub: infrastructure.GitHubClient(ctx)}
}

func NewRepositoriesRegistry() Repositories {
	return &defaultRepositories{}
}
