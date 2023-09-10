package app

import (
	"fyfirman-blog-service/internal/github_contribution/domain"
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

	c.JSON(http.StatusOK, contrib)
}
