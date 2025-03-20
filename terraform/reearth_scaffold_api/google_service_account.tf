data "google_service_account" "main" {
  project    = data.google_project.project.project_id
  account_id = var.service_account_email
}
