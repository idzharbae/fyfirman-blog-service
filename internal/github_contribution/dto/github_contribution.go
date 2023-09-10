package dto

type GetGithubContributionDTO struct {
	Message string `json:"message"`
	Data    struct {
		TotalContributions int `json:"totalContributions"`
		Weeks              []struct {
			ContributionDays []struct {
				ContributionCount int    `json:"contributionCount"`
				Date              string `json:"date"`
			} `json:"contributionDays"`
		} `json:"weeks"`
	} `json:"data"`
}
