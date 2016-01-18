package slacker

// User represents a Slack user object
// https://api.slack.com/types/user
type User struct {
	ID       string `json:"id"`
	TeamID   string `json:"team_id"`
	Name     string `json:"name"`
	RealName string `json:"real_name"`
	Deleted  bool   `json:"deleted"`
	Color    string `json:"color"`
	IsAdmin  bool   `json:"is_admin"`
	IsOwner  bool   `json:"is_owner"`
	Has2fa   bool   `json:"has_2fa"`
	HasFiles bool   `json:"has_files"`

	Profile `json:"profile"`
}

// Profile represents a more detailed profile of a Slack user, including things
// like avatars.
type Profile struct {
	AvatarHash         string `json:"avatar_hash"`
	Email              string `json:"email"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Image192           string `json:"image_192"`
	Image24            string `json:"image_24"`
	Image32            string `json:"image_32"`
	Image48            string `json:"image_48"`
	Image512           string `json:"image_512"`
	Image72            string `json:"image_72"`
	RealName           string `json:"real_name"`
	RealNameNormalized string `json:"real_name_normalized"`
}

// UsersList is a response object wrapper for users.list in Slack
// https://api.slack.com/methods/users.list
type UsersList struct {
	Users []*User `json:"members"`
}

// UsersList returns all users in the team
func (c *APIClient) UsersList() ([]*User, error) {
	dest := UsersList{}
	if err := c.slackMethodAndParse("users.list", &dest); err != nil {
		return nil, err
	}

	return dest.Users, nil
}
