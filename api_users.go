package slacker

// User represents a Slack user object
// https://api.slack.com/types/user
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Deleted  bool   `json:"deleted"`
	Color    string `json:"color"`
	IsAdmin  bool   `json:"is_admin"`
	IsOwner  bool   `json:"is_owner"`
	Has2fa   bool   `json:"has_2fa"`
	HasFiles bool   `json:"has_files"`
}

// UsersList returns all users in the team
func (c *APIClient) UsersList() ([]*User, error) {
	// resp, err := c.slackMethod("users.list")
	return nil, nil
}
