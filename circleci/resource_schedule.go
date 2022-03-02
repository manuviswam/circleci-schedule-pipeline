package circleci

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
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
	circleClient := i.(*Client)
	s := getScheduleFromData(data)

	id, err := circleClient.createSchedule(s)
	if err != nil {
		return diag.FromErr(err)
	}
	data.SetId(id)

	return resourceScheduleRead(ctx, data, i)
}

func resourceScheduleRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	circleClient := i.(*Client)
	_, err := circleClient.readSchedule(data.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	//if err := data.Set("schedule", schedule); err != nil {
	//	return diag.FromErr(err)
	//}
	return nil
}

func resourceScheduleUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	circleClient := i.(*Client)
	s := getScheduleFromData(data)
	err := circleClient.updateSchedule(data.Id(), s)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("last_updated", time.Now().Format(time.RFC850)); err != nil {
		return diag.FromErr(err)
	}
	return resourceScheduleRead(ctx, data, i)
}

func resourceScheduleDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	circleClient := i.(*Client)
	err := circleClient.deleteSchedule(data.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	data.SetId("")
	return nil
}

func getScheduleFromData(data *schema.ResourceData) Schedule {
	name := data.Get("name").(string)
	timetable := data.Get("timetable").(*schema.Set).List()[0].(map[string]interface{})
	perHour := timetable["per_hour"].(int)
	hoursOfDay := make([]int, len(timetable["hours_of_day"].([]interface{})))
	for _, v := range timetable["hours_of_day"].([]interface{}) {
		hoursOfDay = append(hoursOfDay, v.(int))
	}
	daysOfWeek := make([]string, len(timetable["days_of_week"].([]interface{})))
	for _, v := range timetable["days_of_week"].([]interface{}) {
		daysOfWeek = append(daysOfWeek, v.(string))
	}

	attributionActor := data.Get("attribution_actor").(string)
	parameters := data.Get("parameters").(map[string]interface{})
	description := data.Get("description").(string)

	s := Schedule{
		Name: name,
		TimeTable: TimeTable{
			PerHour:    perHour,
			HoursOfDay: hoursOfDay,
			DaysOfWeek: daysOfWeek,
		},
		AttributionActor: attributionActor,
		Parameters:       parameters,
		Description:      description,
	}
	return s
}
