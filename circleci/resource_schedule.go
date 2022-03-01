package circleci

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSchedule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceScheduleCreate,
		ReadContext:   resourceScheduleRead,
		UpdateContext: resourceScheduleUpdate,
		DeleteContext: resourceScheduleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"timetable": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"per_hour": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
							//	validate number 1 to 60
						},
						"hours_of_day": &schema.Schema{
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
							Required: true,
						},
						"days_of_week": &schema.Schema{
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Required: true,
							//	"TUE" "SAT" "SUN" "MON" "THU" "WED" "FRI"
						},
					},
				},
			},
			"attribution_actor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				//	"current" "system"
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceScheduleCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	circleClient := i.(*Client)
	name := data.Get("name").(string)
	timetable := data.Get("timetable").(map[string]interface{})
	perHour := timetable["per_hour"].(int)
	hoursOfDay := timetable["hours_of_day"].([]int)
	daysOfWeek := timetable["days_of_week"].([]string)
	attributionActor := data.Get("attribution_actor").(string)
	parameters := data.Get("parameters").(map[string]string)
	description := data.Get("description").(string)
	fmt.Println("client is {}", circleClient)
	fmt.Println("name is {}", name)
	fmt.Println("timetable is {}", timetable)
	fmt.Println("perHour is {}", perHour)
	fmt.Println("hoursOfDay is {}", hoursOfDay)
	fmt.Println("daysOfWeek is {}", daysOfWeek)
	fmt.Println("attributionActor is {}", attributionActor)
	fmt.Println("parameters is {}", parameters)
	fmt.Println("description is {}", description)
	return diags
}

func resourceScheduleRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

func resourceScheduleUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

func resourceScheduleDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}
