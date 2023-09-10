package githubgql

import (
	"context"

	"github.com/pkg/errors"

	"github.com/shurcooL/githubv4"
	"golang.org/x/exp/slog"
	"golang.org/x/oauth2"
)

type Response struct {
	User struct {
		ContributionsCollection struct {
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
	}
}

func GetContributionFromGithub(username string) (*Response, error) {
	token := ""

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	var query struct {
		User struct {
			ContributionsCollection struct {
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
		} `graphql:"user(login: $userName)"`
	}

	var response Response

	err := client.Query(context.Background(), &query, map[string]interface{}{
		"userName": username,
	})

	if err != nil {
		return nil, errors.Wrap(err, "Failed to get contribution to Github")
	}

	slog.Info("Fetched. Total Contributions: %d\n", query.User.ContributionsCollection.ContributionCalendar.TotalContributions)

	response.User = query.User

	return &response, nil
}
