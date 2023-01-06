variable "github_api_token" {
  type = string
  description = "The GitHub API token."
  default     = "default"
}
variable "ssh_key_pub" {
  type = string
  description = "The public SSH key"
  default     = "default"
}

variable "github_repo_name" {
  type = string
  description = "The repo name to check for action secrets"
  default     = "default"
}
