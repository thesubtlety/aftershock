resource "github_user_ssh_key" "example" {
  title = "example title"
  key   = file(var.ssh_key_pub) # To do variable
}