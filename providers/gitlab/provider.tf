terraform {
  required_providers {
    gitlab = {
      source  = "gitlabhq/gitlab"
      version = "~> 15.0"
    }
  }
}

provider "gitlab" {
  token = var.gitlab_token
}