package circleci

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceSchedule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: false,
			},
			"timetable": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"per-hour": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: false,
							//	validate number 1 to 60
						},
						"hours-of-day": &schema.Schema{
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
							Optional: false,
						},
						"days-of-week": &schema.Schema{
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional: false,
							//	"TUE" "SAT" "SUN" "MON" "THU" "WED" "FRI"
						},
					},
				},
			},
			"attribution-actor": &schema.Schema{
				Type: schema.TypeString,
				//	"current" "system"
				Optional: false,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
