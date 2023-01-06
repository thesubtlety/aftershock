terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

# Configure the AWS Provider
# Will automatically pick up env vars
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs
provider "aws" {
  region = var.region
  access_key = var.aws_access_key_id
  secret_key = var.aws_secret_access_key
  #session_token = var.aws_session_token
  #shared_config_files      = ["/Users/tf_user/.aws/conf"]
  #shared_credentials_files = ["/Users/tf_user/.aws/creds"]
  # assume_role {
  #   role_arn     = var.role_arn
  #   session_name = "SESSION_NAME"
  #   external_id  = "EXTERNAL_ID"
  # }
}