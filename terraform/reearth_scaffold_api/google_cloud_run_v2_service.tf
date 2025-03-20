resource "google_cloud_run_v2_service" "main" {
  project  = data.google_project.project.project_id
  name     = local.service_name
  location = var.region
  ingress  = "INGRESS_TRAFFIC_ALL"

  template {
    containers {
      image = var.image

      resources {
        cpu_idle          = true
        startup_cpu_boost = true
      }

      env {
        name  = "REEARTH_AUTH0_AUDIENCE"
        value = var.auth0_audience
      }

      env {
        name  = "REEARTH_AUTH0_DOMAIN"
        value = var.auth0_domain
      }

      env {
        name = "REEARTH_DB"
        value_source {
          secret_key_ref {
            secret  = var.database_secret_id
            version = "latest"
          }
        }
      }

      env {
        name  = "REEARTH_HOST"
        value = "https://${local.api_domain}"
      }

      env {
        name  = "REEARTH_CMS_ORIGINS"
        value = "https://${local.web_domain}"
      }
    }

    scaling {
      max_instance_count = 100
      min_instance_count = 0
    }

    service_account = var.service_account_email
  }

  traffic {
    percent = 100
    type    = "TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST"
  }

  lifecycle {
    ignore_changes = [
      client,
      client_version,

      # Currently, we restrict production environment to use all ingress.
      # It will be released in 2025 so I'll just ignore changes as for now.
      ingress,

      template[0].containers[0].image,
      template[0].containers[0].name,
      template[0].revision,
      traffic[0].revision,
      traffic[0].type,
    ]
  }
}
