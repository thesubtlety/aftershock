terraform {
  required_providers {
    ultradns = {
      source = "ultradns/ultradns"
      version = ">= 1.3.0"
    }
  }
}

# Configure the UltraDNS Provider
provider "ultradns" {
  username = var.username
  password = var.password
  hosturl = var.hosturl
}