# Retrieve information about the currently authenticated user.
data "github_user" "current" {
  username = ""
}
output "current_github_login" {
  value = "${data.github_user.current.login}"
}
output "user_info" {
  value = data.github_user.current
}

# Private Repo Information
data "github_repositories" "recon" {
  query = "is:private user:${data.github_user.current.login}"
}
output "private_repos_recon" {
  value = data.github_repositories.recon
}

# Get detail info on private repos
data "github_repository" "all" {
  for_each = toset(data.github_repositories.recon.full_names)
  full_name = each.key
}
output "github_repository_all_output" {
  value = data.github_repository.all
}
