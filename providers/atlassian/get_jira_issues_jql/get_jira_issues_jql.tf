# Get JQL Issues
data "jira_jql" "issues" {
  jql = "${var.jql_text}" # summary ~ "query here" OR description ~ "query here"
}

# Print it
output "jira_jql_issues" {
  description = "JIRA JQL response"
  value = "${data.jira_jql.issues}"
}
