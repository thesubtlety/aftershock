variable "cloudflare_api_token" {
  type = string
  description = "The Cloudflare API token."
  default     = "default"
}

variable "cloudflare_user_id" {
  description = "The Cloudflare user ID."
  default     = "default"
}

# variable "cloudflare_zone" {
#   description = "The Cloudflare zone."
#   default     = "default"
# }