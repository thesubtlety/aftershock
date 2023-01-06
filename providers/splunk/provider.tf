terraform {
  required_providers {
    splunk = {
      source  = "splunk/splunk"
    }
  }
}

provider "splunk" {
  url                  = var.url
  username             = var.username
  password             = var.password
  insecure_skip_verify = true
}

