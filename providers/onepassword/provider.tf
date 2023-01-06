terraform {
  required_providers {
    onepassword = {
      source = "anasinnyk/onepassword"
      version = "1.3.0"
    }
  }
}

# if email, password, secret_key not set will check OP_SESSION_<subdomain>
# https://registry.terraform.io/providers/milosbackonja/1password/latest/docs
provider "onepassword" {
    email      = var.email
    password   = var.password
    secret_key = var.secret_key
    subdomain  = var.subdomain
}