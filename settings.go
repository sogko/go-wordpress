package wordpress

import "context"

// Settings represents a WordPress settings.
type Settings struct {
	Title                string `json:"title"`
	Description          string `json:"description"`
	URL                  string `json:"url"`
	Email                string `json:"email"`
	Timezone             string `json:"timezone"`
	DateFormat           string `json:"date_format"`
	TimeFormat           string `json:"time_format"`
	StartOfWeek          int    `json:"start_of_week"`
	Language             string `json:"language"`
	UseSmilies           bool   `json:"use_smilies"`
	DefaultCategory      int    `json:"default_category"`
	DefaultPostFormat    string `json:"default_post_format"`
	PostsPerPage         int    `json:"posts_per_page"`
	DefaultPingStatus    string `json:"default_ping_status"`
	DefaultCommentStatus string `json:"default_comment_status"`
}

// SettingsService provides access to the settings related functions in the WordPress REST API.
type SettingsService service

// List returns a list of settingss.
func (c *SettingsService) List(ctx context.Context) (*Settings, *Response, error) {
	var settings Settings
	resp, err := c.client.List(ctx, "settings", nil, &settings)
	return &settings, resp, err
}
