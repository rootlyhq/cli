package models

import "time"

// Structure for a pulse
type Pulse struct {
	Summary        string
	StartedAt      time.Time
	EndedAt        time.Time
	ServiceIds     []string
	EnvironmentIds []string
	Labels         []map[string]string
	Source         string
	Refs           []map[string]string
}

// Structure for a created pulse response
type CreatedPulseResponse struct {
	Data struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Summary string `json:"summary"`
			Labels  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"labels"`
			Services []struct {
				ID                     string        `json:"id"`
				TeamID                 int           `json:"team_id"`
				Name                   string        `json:"name"`
				Slug                   string        `json:"slug"`
				Description            string        `json:"description"`
				SlackChannels          []string      `json:"slack_channels"`
				DeletedAt              interface{}   `json:"deleted_at"`
				CreatedAt              string        `json:"created_at"`
				UpdatedAt              string        `json:"updated_at"`
				OpsgenieID             interface{}   `json:"opsgenie_id"`
				PagerdutyID            interface{}   `json:"pagerduty_id"`
				NotifyEmails           []interface{} `json:"notify_emails"`
				PublicDescription      string        `json:"public_description"`
				GithubRepositoryBranch string        `json:"github_repository_branch"`
				GithubRepositoryName   interface{}   `json:"github_repository_name"`
				Color                  string        `json:"color"`
				HerokuAppName          interface{}   `json:"heroku_app_name"`
				GitlabRepositoryName   string        `json:"gitlab_repository_name"`
				GitlabRepositoryBranch interface{}   `json:"gitlab_repository_branch"`
			} `json:"services"`
			Environments []struct {
				ID             string        `json:"id"`
				TeamID         int           `json:"team_id"`
				Name           string        `json:"name"`
				Slug           string        `json:"slug"`
				Description    string        `json:"description"`
				Color          string        `json:"color"`
				DeletedAt      interface{}   `json:"deleted_at"`
				CreatedAt      string        `json:"created_at"`
				UpdatedAt      string        `json:"updated_at"`
				IncidentsCount int           `json:"incidents_count"`
				NotifyEmails   []interface{} `json:"notify_emails"`
				SlackChannels  []interface{} `json:"slack_channels"`
			} `json:"environments"`
			StartedAt string `json:"started_at"`
			EndedAt   string `json:"ended_at"`
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at"`
		} `json:"attributes"`
	} `json:"data"`
}
