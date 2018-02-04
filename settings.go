package wordpress

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

type SettingsCollection struct {
	client *Client
	url    string
}

func (col *SettingsCollection) List() (*Settings, *Response, []byte, error) {
	var settings Settings
	resp, body, err := col.client.List(col.url, nil, &settings)
	return &settings, newResponse(resp), body, err
}
