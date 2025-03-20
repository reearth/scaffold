resource "google_project_iam_member" "main" {
  for_each = toset([
  ])

  project = data.google_project.project.project_id
  role    = each.value
  member  = "serviceAccount:${data.google_service_account.main.email}"
}
