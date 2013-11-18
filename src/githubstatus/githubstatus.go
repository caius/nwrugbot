package githubstatus

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var StatusURL = "https://status.github.com/api/last-message.json"

type Response struct {
	Mood        string
	Description string
	UpdatedAt   string
}

type RawResponse struct {
	Status    string
	Body      string
	CreatedOn string
}

func Status() (Response, error) {
	r := Response{}

	err := r.Run()
	if err != nil {
		return Response{}, err
	}

	return r, nil
}

func (r *Response) Run() error {
	resp, err := http.Get(StatusURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	raw := RawResponse{}

	err = json.Unmarshal(body, &raw)
	if err != nil {
		return err
	}

	r.Mood = raw.Status
	r.Description = raw.Body
	r.UpdatedAt = raw.CreatedOn

	return nil
}
