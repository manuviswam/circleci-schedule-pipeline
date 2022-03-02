package circleci

type Client struct {
	ProjectSlug string
	Token       string
}

type TimeTable struct {
	PerHour    int
	HoursOfDay []int
	DaysOfWeek []string
}

type Schedule struct {
	Name             string
	TimeTable        TimeTable
	AttributionActor string
	Parameters       map[string]interface{}
	Description      string
}

func newClient(projectSlug, token string) *Client {
	return &Client{
		ProjectSlug: projectSlug,
		Token:       token,
	}
}

func (c *Client) createSchedule(s Schedule) (id string, e error) {
	return "foo", nil
}

func (c *Client) updateSchedule(id string, s Schedule) error {
	return nil
}

func (c *Client) readSchedule(id string) (s Schedule, e error) {
	return
}

func (c *Client) deleteSchedule(id string) error {
	return nil
}
