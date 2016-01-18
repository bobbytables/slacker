package slacker

// Channel represents a Slack channel
// https://api.slack.com/types/channel
type Channel struct {
	Created    int      `json:"created"`
	Creator    string   `json:"creator"`
	ID         string   `json:"id"`
	IsArchived bool     `json:"is_archived"`
	IsChannel  bool     `json:"is_channel"`
	IsGeneral  bool     `json:"is_general"`
	IsMember   bool     `json:"is_member"`
	Members    []string `json:"members"`
	Name       string   `json:"name"`
	NumMembers int      `json:"num_members"`

	Purpose ChannelPurpose `json:"purpose"`
	Topic   ChannelTopic   `json:"topic"`
}

// ChannelPurpose represents a channels' purpose in Slack
type ChannelPurpose struct {
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
	Value   string `json:"value"`
}

// ChannelTopic represents a channel's topic in Slack
type ChannelTopic struct {
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
	Value   string `json:"value"`
}

// ChannelsList is a wrapper for a channels.list API call
type ChannelsList struct {
	Channels []*Channel `json:"channels"`
}

// ChannelsList returns a list of Channels from Slack
func (c *APIClient) ChannelsList() ([]*Channel, error) {
	dest := ChannelsList{}
	if err := c.slackMethodAndParse("channels.list", &dest); err != nil {
		return nil, err
	}

	return dest.Channels, nil
}
