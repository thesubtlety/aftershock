variable "jql_text" {
  type = string
  description = "JQL query string"
  default     = "default"
}

variable "url" {
  type = string
  description = "Atlassian url"
  default     = "default"
}

variable "user" {
  type = string
  description = "Atlassian user"
  default     = "default"
}

variable "password" {
  type = string
  description = "Atlassian password"
  default     = "default"
}

variable "new_user" {
  type = string
  description = "New Atlassian user name"
  default     = "default"
}

variable "new_user_email" {
  type = string
  description = "New Atlassian user email"
  default     = "default"
}
