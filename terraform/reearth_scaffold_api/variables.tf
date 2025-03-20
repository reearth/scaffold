variable "domain" {
  type        = string
  description = "Custom domain to be used."
}

variable "auth0_audience" {
  type        = string
  description = "Audient of the Auth0 credential."
}

variable "auth0_domain" {
  type        = string
  description = "Domain of the Auth0 tenant."
}

variable "database_secret_id" {
  type        = string
  description = "Secret ID of the database connection URL."
}

variable "image" {
  type        = string
  description = "Image of the Cloud Run service."
}

variable "project" {
  description = "ID of the GCP project."
  type        = string

  validation {
    condition     = can(regex("^[a-z][a-z0-9-]{4,28}[a-z0-9]$", var.project))
    error_message = "Project ID must start with a lowercase letter, and can include lowercase letters, numbers, or hyphens. It must be between 6 and 30 characters long."
  }
}

variable "region" {
  type        = string
  description = "Region to host the resources."

  validation {
    condition     = can(regex("^[a-z]+-[a-z]+[0-9]+$", var.region))
    error_message = "Region must be in the format of <region>-<zone>."
  }
}

variable "service_account_email" {
  type        = string
  description = "Email of the service account to be used."

  validation {
    condition     = can(regex("^[a-z0-9-_]+@[a-z0-9-_.]+$", var.service_account_email))
    error_message = "Service account email must be in the format of <name>@<domain>."
  }
}
