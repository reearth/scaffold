output "network_endpoint_group" {
  value = google_compute_region_network_endpoint_group.main
}

output "service" {
  value = google_cloud_run_v2_service.main
}
