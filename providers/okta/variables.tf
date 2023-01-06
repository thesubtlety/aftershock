variable "api_token" {
  type = string
  description = "The Okta API token."
  default     = "default"
}
variable "org_name" {
  type = string
  description = "The Okta org name"
  default     = "default"
}
variable "base_url" {
  type = string
  description = "The Okta base URL, e.g okta.com or oktapreview.com"
  default     = "default"
}
