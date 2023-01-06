# Get current user
data "bitbucket_current_user" "aftershock" {}
output "bitbucket_current_user_info" {
  description = "bitbucket current user info"
  value = "${data.bitbucket_current_user.aftershock}"
}
