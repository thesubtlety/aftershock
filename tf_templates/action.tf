# Data Sources
data "example_provider_accounts" "aftershock" { }
output "example_provider_account_info" {
  description = "example_provider account info"
  value = "${data.example_provider_accounts.aftershock.accounts}"
}



# Resources
resource "example_provider_resource" "aftershock" {
  name = "user"
}
