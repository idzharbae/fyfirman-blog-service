package domain

import (
	githubgql "fyfirman-blog-service/internal/github_contribution/infras/github_gql"

	"github.com/pkg/errors"
	"github.com/shurcooL/githubv4"
)

type ContributionsCollection struct {
	ContributionCalendar struct {
		TotalContributions int
		Weeks              []struct {
			ContributionDays []struct {
				ContributionCount int
				Date              githubv4.DateTime
			}
		}
	}
}

func GetGithubContribution(username string) (*ContributionsCollection, error) {
	contrib := GetContributionFromDB(username)

	if contrib != nil {
		return contrib, nil
	}

	contribFromGithub, err := githubgql.GetContributionFromGithub(username)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to get contribution to Github")
	}

	return (*ContributionsCollection)(&contribFromGithub.User.ContributionsCollection), nil
}

func GetContributionFromDB(username string) *ContributionsCollection {
	return nil
}
