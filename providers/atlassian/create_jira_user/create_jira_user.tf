# Create a JIRA user
resource "jira_user" "aftershock" {
  name = var.new_user
  email = var.new_user_email
  display_name = "Aftershock"
}

# Print it
output "new_jira_user_output" {
  description = "New JIRA user information"
  value = "${data.jira_user.aftershock}"
}
