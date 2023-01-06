terraform {
  required_providers {
    provider_example = {
      source  = "provider_example/provider_example"
      version = "~> 1.0"
    }
  }
}
provider "provider_example" {
  api_token = var.provider_example_api_token
}
