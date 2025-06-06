resource "google_monitoring_slo" "availability" {
  project = data.google_project.project.project_id
  service = google_monitoring_service.main.service_id

  calendar_period = "MONTH"
  display_name    = "95% Availability Monthly"
  slo_id          = "95-availability-monthly"
  goal            = 0.95

  basic_sli {
    availability {
      enabled = true
    }
  }
}

resource "google_monitoring_slo" "latency" {
  project = data.google_project.project.project_id
  service = google_monitoring_service.main.service_id

  calendar_period = "MONTH"
  display_name    = "95% Latency Monthly"
  goal            = 0.95
  slo_id          = "95-latency-monthly"

  basic_sli {
    latency {
      threshold = "10s"
    }
  }
}
