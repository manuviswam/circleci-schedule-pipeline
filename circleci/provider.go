package circleci

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"project_slug": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"circle_token": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
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
	projectSlug := data.Get("project_slug").(string)
	token := data.Get("circle_token").(string)
	fmt.Println(projectSlug, token)
	c := newClient(projectSlug, token)
	return c, diags
}
