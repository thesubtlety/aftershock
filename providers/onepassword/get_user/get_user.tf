# Get user info
data "onepassword_user" "this" {
    email = var.target_email
}
output "onepassword_user_output" {
  description = "Onepassword user info"
  value = "${data.onepassword_user.this}"
}
