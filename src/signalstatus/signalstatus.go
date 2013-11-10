package signalstatus

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

var StatusURL string = "http://status.37signals.com/status.json"

type StatusCheck struct {
	Status StatusBody
}

type StatusBody struct {
	Mood        string
	Description string
	UpdatedAt   time.Time
}

type StatusResponse struct {
	Status StatusBody
}

func Status() (StatusCheck, error) {
	sc := StatusCheck{}

	err := sc.Run()
	if err != nil {
		return StatusCheck{}, err
	}

	return sc, nil
}

func (sc *StatusCheck) Run() error {
	resp, err := http.Get(StatusURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	status_response := StatusResponse{}
	err = json.Unmarshal(body, &status_response)
	if err != nil {
		return err
	}

	sc.Status = status_response.Status

	return nil
}

func (sc *StatusCheck) OK() bool {
	return sc.Status.Mood == "good"
}
