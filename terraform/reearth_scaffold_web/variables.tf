variable "api_url" {
  type        = string
  description = "URL of the API server."
  default     = ""
}

variable "auth0_audience" {
  type        = string
  description = "Audient of the Auth0 credential."
}

variable "auth0_client_id" {
  type        = string
  description = "Client ID of the Auth0 SPA client."
}

variable "auth0_domain" {
  type        = string
  description = "Domain of the Auth0 tenant."
}

variable "image" {
  type        = string
  description = "Image of the Cloud Run service."
}

variable "favicon_url" {
  type        = string
  description = "URL of the favicon."
  default     = ""
}

variable "title" {
  type        = string
  description = "Title of HTML."
  default     = ""
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
