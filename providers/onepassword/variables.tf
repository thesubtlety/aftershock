variable "email" {
  type = string
  description = "The OnePassword email"
  default     = "default"
}
variable "password" {
  type = string
  description = "The OnePassword password"
  default     = "default"
}
variable "secret_key" {
  type = string
  description = "The OnePassword secret_key"
  default     = "default"
}
variable "subdomain" {
  type = string
  description = "The OnePassword subdomain"
  default     = "default"
}
variable "target_email" {
  type = string
  description = "The onepassword_user target_email"
  default     = "default"
}
