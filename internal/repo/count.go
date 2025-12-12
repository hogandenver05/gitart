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
// Returns 0 and an error if the API call fails (lenient error handling).
func GetContributionCount(day time.Time) (int, error) {
	utcDay := day.UTC()
	year, month, date := utcDay.Date()
	normalizedDay := time.Date(year, month, date, 0, 0, 0, 0, time.UTC)

	fromDate := normalizedDay.Format(time.RFC3339)
	toDate := normalizedDay.AddDate(0, 0, 1).Format(time.RFC3339)

	query := `query($from: DateTime!, $to: DateTime!) { viewer { contributionsCollection(from: $from, to: $to) { contributionCalendar { weeks { contributionDays { date contributionCount } } } } } }`

	variablesJSON, err := json.Marshal(map[string]string{
		"from": fromDate,
		"to":   toDate,
	})
	if err != nil {
		return 0, fmt.Errorf("failed to marshal variables: %w", err)
	}

	cmd := exec.Command("gh", "api", "graphql", "-f", fmt.Sprintf("query=%s", query), "-f", fmt.Sprintf("variables=%s", string(variablesJSON)))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("failed to query GitHub API: %w, output: %s", err, string(output))
	}

	var response contributionResponse
	if err := json.Unmarshal(output, &response); err != nil {
		return 0, fmt.Errorf("failed to parse API response: %w, output: %s", err, string(output))
	}

	targetDateStr := normalizedDay.Format("2006-01-02")

	for _, week := range response.Data.Viewer.ContributionsCollection.ContributionCalendar.Weeks {
		for _, contributionDay := range week.ContributionDays {
			contributionDate, err := time.Parse(time.RFC3339, contributionDay.Date)
			if err != nil {
				continue
			}
			contributionDateStr := contributionDate.UTC().Format("2006-01-02")
			if contributionDateStr == targetDateStr {
				return contributionDay.ContributionCount, nil
			}
		}
	}

	return 0, nil
}
