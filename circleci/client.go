package circleci

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

type actor struct {
	Id    string `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

type ScheduleResponse struct {
	Schedule
	Id          string    `json:"id"`
	UpdatedAt   time.Time `json:"updated-at"`
	CreatedAt   time.Time `json:"created-at"`
	ProjectSlug string    `json:"project-slug"`
	Actor       actor     `json:"actor"`
}

const scheduleApi = "https://circleci.com/api/v2/project/%s/schedule"
const scheduleApiWithId = "https://circleci.com/api/v2/schedule/%s"

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
	return createOrUpdate(c, s, "PATCH", fmt.Sprintf(scheduleApiWithId, id), 200)
}

func (c *Client) readSchedule(id string) (s ScheduleResponse, err error) {
	url := fmt.Sprintf(scheduleApiWithId, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Add("circle-token", c.Token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		return s, errors.New("Request failed with status : " + res.Status + string(body))
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &s)
	return
}

func (c *Client) deleteSchedule(id string) (err error) {

	url := fmt.Sprintf(scheduleApiWithId, id)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return
	}
	req.Header.Add("circle-token", c.Token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		return errors.New("Request failed with status : " + res.Status + string(body))
	}
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
