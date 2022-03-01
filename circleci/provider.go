package circleci

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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
				Type:      schema.TypeString,
				Optional:  false,
				Sensitive: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"circleci_schedule": resourceSchedule(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	projectSlug := data.Get("project-slug").(string)
	token := data.Get("circle-token").(string)
	c := newClient(projectSlug, token)
	return c, diags
}
