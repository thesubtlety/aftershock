terraform {
  required_providers {
    salesforce = {
      source = "hashicorp/salesforce"
      version = "0.1.0"
    }
  }
}

#https://registry.terraform.io/providers/hashicorp/salesforce/latest/docs
#export SALESFORCE_CLIENT_ID=""
#export SALESFORCE_PRIVATE_KEY=""
#export SALESFORCE_API_VERSION=""
#export SALESFORCE_USERNAME=""
provider "salesforce" {
  client_id   = var.client_id
  private_key = var.private_key
  api_version = "53.0"
  username    = var.username
}
