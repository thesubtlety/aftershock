# create new user
resource "googleworkspace_user" "aftershock" {
  primary_email = var.create_user_primary_email
  password      = var.create_user_md5password
  hash_function = "MD5"

  name {
    family_name = var.create_user_lastname
    given_name  = var.create_user_firstname
  }

  recovery_email = var.create_user_recover_email
}