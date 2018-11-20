package repositories

import (
	"context"

	"github.com/google/go-github/v18/github"
	"github.com/int128/gradleupdate/app/domain"
	"github.com/int128/gradleupdate/app/domain/repositories"
	"github.com/pkg/errors"
)

type PullRequest struct {
	GitHub *github.Client
}

func (r *PullRequest) Query(ctx context.Context, q repositories.PullRequestQuery) ([]domain.PullRequest, error) {
	payloads, _, err := r.GitHub.PullRequests.List(ctx, q.Base.Owner, q.Base.Repo, &github.PullRequestListOptions{
		Base:        q.Base.Branch,
		Head:        q.Head.Branch,
		State:       q.State,
		Direction:   q.Direction,
		Sort:        q.Sort,
		ListOptions: github.ListOptions{Page: q.Page, PerPage: q.PerPage},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "GitHub API returned error")
	}
	pulls := make([]domain.PullRequest, len(payloads))
	for i, payload := range payloads {
		head := payload.GetHead()
		base := payload.GetBase()
		pulls[i] = domain.PullRequest{
			PullRequestIdentifier: domain.PullRequestIdentifier{
				RepositoryIdentifier: domain.RepositoryIdentifier{Owner: base.GetUser().GetLogin(), Repo: base.GetRepo().GetName()},
				PullRequestNumber:    payload.GetNumber(),
			},
			Head: domain.BranchIdentifier{
				RepositoryIdentifier: domain.RepositoryIdentifier{Owner: head.GetUser().GetLogin(), Repo: head.GetRepo().GetName()},
				Branch:               head.GetRef(),
			},
			Base: domain.BranchIdentifier{
				RepositoryIdentifier: domain.RepositoryIdentifier{Owner: base.GetUser().GetLogin(), Repo: base.GetRepo().GetName()},
				Branch:               base.GetRef(),
			},
			Title: payload.GetTitle(),
			Body:  payload.GetBody(),
		}
	}
	return pulls, nil
}

func (r *PullRequest) Create(ctx context.Context, pull domain.PullRequest) (domain.PullRequest, error) {
	payload, _, err := r.GitHub.PullRequests.Create(ctx, pull.Owner, pull.Repo, &github.NewPullRequest{
		Base:  github.String(pull.Base.Branch),
		Head:  github.String(pull.Head.Branch),
		Title: github.String(pull.Title),
		Body:  github.String(pull.Body),
	})
	if err != nil {
		return domain.PullRequest{}, errors.Wrapf(err, "GitHub API returned error")
	}
	head := payload.GetHead()
	base := payload.GetBase()
	return domain.PullRequest{
		PullRequestIdentifier: domain.PullRequestIdentifier{
			RepositoryIdentifier: domain.RepositoryIdentifier{Owner: base.GetUser().GetLogin(), Repo: base.GetRepo().GetName()},
			PullRequestNumber:    payload.GetNumber(),
		},
		Head: domain.BranchIdentifier{
			RepositoryIdentifier: domain.RepositoryIdentifier{Owner: head.GetUser().GetLogin(), Repo: head.GetRepo().GetName()},
			Branch:               head.GetRef(),
		},
		Base: domain.BranchIdentifier{
			RepositoryIdentifier: domain.RepositoryIdentifier{Owner: base.GetUser().GetLogin(), Repo: base.GetRepo().GetName()},
			Branch:               base.GetRef(),
		},
		Title: payload.GetTitle(),
		Body:  payload.GetBody(),
	}, nil
}

func (r *PullRequest) Update(ctx context.Context, pull domain.PullRequest) (domain.PullRequest, error) {
	payload, _, err := r.GitHub.PullRequests.Edit(ctx, pull.Owner, pull.Repo, pull.PullRequestNumber, &github.PullRequest{
		Title: github.String(pull.Title),
		Body:  github.String(pull.Body),
	})
	if err != nil {
		return domain.PullRequest{}, errors.Wrapf(err, "GitHub API returned error")
	}
	head := payload.GetHead()
	base := payload.GetBase()
	return domain.PullRequest{
		PullRequestIdentifier: domain.PullRequestIdentifier{
			RepositoryIdentifier: domain.RepositoryIdentifier{Owner: base.GetUser().GetLogin(), Repo: base.GetRepo().GetName()},
			PullRequestNumber:    payload.GetNumber(),
		},
		Head: domain.BranchIdentifier{
			RepositoryIdentifier: domain.RepositoryIdentifier{Owner: head.GetUser().GetLogin(), Repo: head.GetRepo().GetName()},
			Branch:               head.GetRef(),
		},
		Base: domain.BranchIdentifier{
			RepositoryIdentifier: domain.RepositoryIdentifier{Owner: base.GetUser().GetLogin(), Repo: base.GetRepo().GetName()},
			Branch:               base.GetRef(),
		},
		Title: payload.GetTitle(),
		Body:  payload.GetBody(),
	}, nil
}