package repo

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

type contributionResponse struct {
	Data struct {
		Viewer struct {
			ContributionsCollection struct {
				ContributionCalendar struct {
					Weeks []struct {
						ContributionDays []struct {
							Date              string `json:"date"`
							ContributionCount int    `json:"contributionCount"`
						} `json:"contributionDays"`
					} `json:"weeks"`
				} `json:"contributionCalendar"`
			} `json:"contributionsCollection"`
		} `json:"viewer"`
	} `json:"data"`
}

// GetContributionCount queries GitHub for the number of contributions on a specific day.
func GetContributionCount(day time.Time) (int, error) {
	utcDay := day.UTC()
	
	startOfDay := time.Date(
		utcDay.Year(), utcDay.Month(), utcDay.Day(),
		00, 00, 00, 0, time.UTC,
	).Format("2006-01-02T15:04:05Z")
	
	endOfDay := time.Date(
		utcDay.Year(), utcDay.Month(), utcDay.Day(),
		23, 59, 59, 0, time.UTC,
	).Format("2006-01-02T15:04:05Z")

	query := `query { viewer { contributionsCollection(from: "%s", to: "%s") { contributionCalendar { weeks { contributionDays { date contributionCount } } } } } }`
	formattedQuery := fmt.Sprintf(query, startOfDay, endOfDay)

	cmd := exec.Command("gh", "api", "graphql",
		"-f", fmt.Sprintf("query=%s", formattedQuery),
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("failed to query GitHub API: %w", err)
	}

	var response contributionResponse
	if err := json.Unmarshal(output, &response); err != nil {
		return 0, fmt.Errorf("failed to parse API response: %w", err)
	}

	targetDateStr := utcDay.Format("2006-01-02")

	for _, week := range response.Data.Viewer.ContributionsCollection.ContributionCalendar.Weeks {
		for _, contributionDay := range week.ContributionDays {
			if contributionDay.Date == targetDateStr {
				return contributionDay.ContributionCount, nil
			}
		}
	}

	return 0, nil
}
 