package circleci

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	ProjectSlug string
	Token       string
}

type TimeTable struct {
	PerHour    int      `json:"per-hour"`
	HoursOfDay []int    `json:"hours-of-day"`
	DaysOfWeek []string `json:"days-of-week"`
}

type Schedule struct {
	Name             string                 `json:"name"`
	TimeTable        TimeTable              `json:"timetable"`
	AttributionActor string                 `json:"attribution-actor"`
	Parameters       map[string]interface{} `json:"parameters"`
	Description      string                 `json:"description"`
}

const scheduleApi = "https://circleci.com/api/v2/project/%s/schedule"

func newClient(projectSlug, token string) *Client {
	return &Client{
		ProjectSlug: projectSlug,
		Token:       token,
	}
}

func (c *Client) createSchedule(s Schedule) (string, error) {
	return createOrUpdate(c, s, "POST", fmt.Sprintf(scheduleApi, c.ProjectSlug), 201)
}

func (c *Client) updateSchedule(id string, s Schedule) (string, error) {
	return createOrUpdate(c, s, "PATCH", fmt.Sprintf(scheduleApi, c.ProjectSlug)+"/"+id, 200)
}

func (c *Client) readSchedule(id string) (s Schedule, e error) {
	return
}

func (c *Client) deleteSchedule(id string) error {
	return nil
}

func createOrUpdate(c *Client, s Schedule, method, url string, expectedStatus int) (id string, err error) {
	httpClient := http.Client{}

	schedule, err := json.Marshal(s)
	if err != nil {
		return
	}
	data := bytes.NewBuffer(schedule)
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("circle-token", c.Token)
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != expectedStatus {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", errors.New("Request failed with status : " + resp.Status + string(body))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var response map[string]json.RawMessage
	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}
	err = json.Unmarshal(response["id"], &id)
	return
}
