package circleci

type Client struct {
	ProjectSlug string
	Token       string
}

type Schedule struct {
}

func newClient(projectSlug, token string) *Client {
	return &Client{
		ProjectSlug: projectSlug,
		Token:       token,
	}
}

func (c *Client) createSchedule(s Schedule) error {
	return nil
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
