terraform {
  required_providers {
    okta = {
      source = "okta/okta"
      version = "~> 3.39"
    }
  }
}

# Configure the Okta Provider
provider "okta" {
  org_name  = var.org_name
  base_url  = var.base_url
  api_token = var.api_token
}
