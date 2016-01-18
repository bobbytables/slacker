package slacker

import (
	"encoding/json"

	"github.com/sdming/gosnow"
)

var sf *gosnow.SnowFlake

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
