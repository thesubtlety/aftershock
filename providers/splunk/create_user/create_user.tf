resource "random_string" "password" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

# This does not print any output if successful
resource "splunk_authentication_users" "aftershock" {
  name              = "aftershock"
  email             = "aftershock@example.com"
  password          = random_string.password.result
  force_change_pass = false
  roles             = ["admin"]
}