package service

import (
	"context"
	"fmt"

	"github.com/int128/gradleupdate/domain"

	"github.com/google/go-github/v18/github"
)

const gradleWrapperPropsPath = "gradle/wrapper/gradle-wrapper.properties"

func GetGradleWrapperVersion(ctx context.Context, owner, repo string) (domain.GradleVersion, error) {
	c := github.NewClient(nil)

	fc, _, _, err := c.Repositories.GetContents(ctx, owner, repo, gradleWrapperPropsPath, nil)
	if err != nil {
		return "", fmt.Errorf("Could not get content: %s", err)
	}
	if fc == nil {
		return "", fmt.Errorf("No such file: %s", gradleWrapperPropsPath)
	}
	content, err := fc.GetContent()
	if err != nil {
		return "", fmt.Errorf("Could not decode content: %s", err)
	}
	v := domain.FindGradleWrapperVersion(content)
	if v == "" {
		return "", fmt.Errorf("Could not determine version from gradle-wrapper.properties")
	}
	return v, nil
}
