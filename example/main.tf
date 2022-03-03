terraform {
  required_providers {
    circleci = {
      version = "0.2"
      source = "manuviswam.com/manuviswam/circleci-schedule-pipeline"
    }
  }
}

provider "circleci" {
  project_slug = "gh/manuviswam/Awesome-schedule"
  circle_token = ""
}

resource "circleci_schedule" "daily_security_audit" {
  name = "daily-security-audit"
  timetable {
    per_hour = 1
    hours_of_day = [1,2,3]
    days_of_week = ["MON","WED"]
  }
  attribution_actor = "current"
  parameters = {
    branch: "main"
  }
  description = "Pipeline to run security audit every day"
}