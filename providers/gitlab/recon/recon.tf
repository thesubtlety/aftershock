# Get current user data
data "gitlab_current_user" "aftershock" {}
output "gitlab_current_users_info" {
  description = "current gitlab user info"
  value = "${data.gitlab_current_user.aftershock}"
}

# Get gitlab_users data
# times out
# data "gitlab_users" "aftershock" { 
#   active = true
# }
# output "gitlab_users_info" {
#   description = "gitlab_users info"
#   value = "${data.gitlab_users.aftershock}"
# }
