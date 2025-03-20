resource "google_monitoring_service" "main" {
  project      = data.google_project.project.project_id
  service_id   = "${local.service_name}-${var.region}" # This has to be unique across the project so we add the region suffix.
  display_name = local.service_name

  basic_service {
    service_type = "CLOUD_RUN"
    service_labels = {
      location     = google_cloud_run_v2_service.main.location
      service_name = google_cloud_run_v2_service.main.name
    }
  }
}
