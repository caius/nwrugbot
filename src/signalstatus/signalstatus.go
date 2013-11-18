package signalstatus

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var StatusURL string = "http://status.37signals.com/status.json"

type StatusCheck struct {
	Status StatusBody
}

type StatusBody struct {
	Mood        string
	Description string
	UpdatedAt   string
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

	err = json.Unmarshal(body, &sc)
	if err != nil {
		return err
	}

	return nil
}

func (sc *StatusCheck) OK() bool {
	return sc.Status.Mood == "good"
}
