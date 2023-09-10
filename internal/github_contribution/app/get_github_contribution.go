package app

import (
	"fyfirman-blog-service/internal/github_contribution/domain"
	"fyfirman-blog-service/internal/github_contribution/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func GetGithubContribution(c *gin.Context) {
	slog.Info("GET: GetGithubContribution")

	contrib, err := domain.GetGithubContribution("fyfirman")

	if err != nil {
		slog.Default().Error(err.Error())
		errorResponse := map[string]interface{}{
			"error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	var response dto.GetGithubContributionDTO
	response.Message = "Data successfully fetched"
	response.Data = struct {
		TotalContributions int "json:\"totalContributions\""
		Weeks              []struct {
			ContributionDays []struct {
				ContributionCount int    "json:\"contributionCount\""
				Date              string "json:\"date\""
			} "json:\"contributionDays\""
		} "json:\"weeks\""
	}(contrib.ContributionCalendar)

	c.JSON(http.StatusOK, response)
}
