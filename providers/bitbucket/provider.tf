terraform {
  required_providers {
    provider_example = {
      source  = "DrFaust92/bitbucket"
      version = "~> 1.0"
    }
  }
}
provider "bitbucket" {
  username = var.username
  password = var.password # you can also use app passwords
  #oauth_client_id = var.oauth_client_id
  #oauth_client_secret = var.oauth_client_secret
  #oauth_token = var.oauth_token
}
