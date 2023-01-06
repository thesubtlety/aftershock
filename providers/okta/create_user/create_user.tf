# Create user
resource "okta_user" "aftershock" {
  first_name         = "John"
  last_name          = "Aftershock"
  login              = "john.aftershock@example.com"
  email              = "john.aftershock@example.com"
}

# Create admin user
resource "okta_user_admin_roles" "aftershock" {
  user_id     = okta_user.aftershock.id
  admin_roles = [
    "APP_ADMIN",
  ]
}