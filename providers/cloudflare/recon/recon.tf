# Get account data
data "cloudflare_accounts" "aftershock" { }
output "cloudflare_account_info" {
  description = "Cloudflare account info"
  value = "${data.cloudflare_accounts.aftershock.accounts}"
}

# Get all active zones in an account
data "cloudflare_zones" "aftershock" {
  filter {
    lookup_type = "contains"
    match       = ".*"
  }
}
output "cloudflare_zone_info" {
  description = "All cloudflare zones"
  value = "${data.cloudflare_zones.aftershock.zones}"
}
