# User permissions
data "cloudflare_api_token_permission_groups" "all" {}

# Token allowed to create new tokens.
resource "cloudflare_api_token" "api_token_create" {
  name = "api_token_create"

  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.permissions["API Tokens Write"],
    ]
    resources = {
      "com.cloudflare.api.user.${var.cloudflare_user_id}" = "*"
    }
  }
}
