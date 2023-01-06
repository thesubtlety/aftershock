# Get users
data "googleworkspace_users" "aftershock" {}
output "googleworkspace_users_output" {
  description = "googleworkspace users"
  value = "${data.googleworkspace_users.aftershock.users}"
}
