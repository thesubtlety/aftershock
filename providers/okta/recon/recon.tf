# Get users data
data "okta_users" "all" {
  include_groups = true
  include_roles = true
  search {
    expression = "status eq \"ACTIVE\""
  }
}
output "okta_users_output" {
  description = "okta users info"
  value = "${data.okta_users.all.users}"
}

# Get groups info
data "okta_groups" "all" { }
output "okta_groups_output" {
  description = "okta groups info"
  value = "${data.okta_groups.all.groups}"
}

# Get trusted origins info
data "okta_trusted_origins" "all" { }
output "okta_trusted_origins_output" {
  description = "okta trusted origins info"
  value = "${data.okta_trusted_origins.all.trusted_origins}"
}

# Get Okta brands
data "okta_brands" "all" { }
output "okta_brands_output" {
  description = "okta brands info"
  value = "${data.okta_brands.all.brands}"
}

