variable "aws_access_key_id" {
  type = string
  description = "The AWS access key"
  default     = "default"
}
variable "aws_secret_access_key" {
  type = string
  description = "The AWS secret key"
  default     = "default"
}
variable "aws_session_token" {
  type = string
  description = "The AWS session token"
  default     = "default"
}
variable "region" {
  type = string
  description = "The AWS region"
  default     = "default"
}
variable "profile" {
  type = string
  description = "The AWS profile"
  default     = "default"
}
