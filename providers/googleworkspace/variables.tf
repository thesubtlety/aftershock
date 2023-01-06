# Provision vars
variable "customer_id" {
  type = string
  description = "The Google Workspace customer ID"
  default     = "default"
}
variable "access_token" {
  type = string
  description = "The Google Workspace customer Access Token"
  default     = "default"
}
variable "credentials_file" {
  type = string
  description = "The Google Workspace cred file, like /Users/mscott/my-project-c633d7053aab.json"
  default     = "default"
}

# Create user vars
variable "create_user_primary_email" {
  type = string
  description = "The Google Workspace create_user_primary_email"
  default     = "default"
}
variable "create_user_md5password" {
  type = string
  description = "The Google Workspace create_user_md5password"
  default     = "default"
}
variable "create_user_firstname" {
  type = string
  description = "The Google Workspace create_user_firstname"
  default     = "default"
}
variable "create_user_lastname" {
  type = string
  description = "The Google Workspace create_user_lastname"
  default     = "default"
}
variable "create_user_recover_email" {
  type = string
  description = "The Google Workspace create_user_recover_email"
  default     = "default"
}
