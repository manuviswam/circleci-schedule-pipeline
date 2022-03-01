package circleci

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"project-slug": &schema.Schema{
				Type:     schema.TypeString,
				Optional: false,
			},
			"circle-token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: false,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"circleci_schedule": resourceSchedule(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}
