terraform {
  required_providers {
    jira = {
      source = "fourplusone/jira"
      version = "0.1.16"
    }
  }
}

provider "jira" {
  url = var.url
  user = var.user
  password = var.password
}
