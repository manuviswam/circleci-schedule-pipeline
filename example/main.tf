terraform {
  required_providers {
    circleci = {
      version = "0.2"
      source = "manuviswam.com/manuviswam/circleci-schedule-pipeline"
    }
  }
}

provider "circleci" {
  project_slug = "slugggg"
  circle_token = "tooookeeennnn"
}

resource "circleci_schedule" "daily_security_audit" {
  name = "daily-security-audit"
  timetable {
    per_hour = 1
    hours_of_day = [1,2,3]
    days_of_week = ["MON","TUE"]
  }
  attribution_actor = "current"
  parameters = {
    foo = "bar"
  }
  description = "Pipeline to run security audit every day"
}