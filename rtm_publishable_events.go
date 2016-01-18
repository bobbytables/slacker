package slacker

import (
	"encoding/json"

	"github.com/sdming/gosnow"
)

var sf *gosnow.SnowFlake

// Some events pushed to Slack require an ID that is always incrementing,
// the Snowflake ID generator is perfect for this.
func init() {
	snowflake, err := gosnow.Default()
	if err != nil {
		panic(err)
	}

	sf = snowflake
}

// Publishable implements Publishable
func (e RTMMessage) Publishable() ([]byte, error) {
	nextID, err := sf.Next()
	if err != nil {
		return nil, err
	}

	e.ID = nextID
	e.Type = "message"

	return json.Marshal(e)
}
