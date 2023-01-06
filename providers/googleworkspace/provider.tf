terraform {
  required_providers {
    googleworkspace = {
      source = "hashicorp/googleworkspace"
      version = "0.7.0"
    }
  }
}

# https://registry.terraform.io/providers/hashicorp/googleworkspace/latest/docs
provider "googleworkspace" {
  customer_id  = var.customer_id
  access_token = var.access_token

  #credentials = var.credentials_file
  #oauth_scopes = [
  #  "https://www.googleapis.com/auth/admin.directory.user",
  #  "https://www.googleapis.com/auth/admin.directory.userschema",
  #  # include scopes as needed
  #]
}