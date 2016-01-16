package slacker

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// Response is the simplest representation of a Slack API response
type Response struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

var (
	// ErrNotAuthed is returned when an API response returns with "not_authed"
	// as it's error attribute.
	ErrNotAuthed = errors.New("slacker: Not Authed")
)

// ParseResponse parses an io.Reader (usually the result of an API request),
// and see's if the response actually contains error information. If it does,
// it will return an error, leaving `dest` untouched. Otherwise, it will
// json decode onto the destination passed in.
func ParseResponse(r io.Reader, dest interface{}) error {
	d := json.NewDecoder(r)
	var rawData json.RawMessage

	if err := d.Decode(&rawData); err != nil {
		return nil
	}

	var resp Response
	if err := json.Unmarshal(rawData, &resp); err != nil {
		return err
	}

	if !resp.Ok {
		return responseError(resp.Error)
	}

	if err := json.Unmarshal(rawData, dest); err != nil {
		return err
	}

	return nil
}

func responseError(ident string) error {
	err, ok := errMap[ident]
	if !ok {
		return fmt.Errorf("slacker: unknown error returned: %s", ident)
	}

	return err
}
