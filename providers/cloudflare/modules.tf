# Run Recon
module "cloudflare_recon" {
  source = "./cloudflare_recon"
  cloudflare_api_token = var.cloudflare_api_token
}
output "recon_output" {
  value = module.cloudflare_recon
}

# Create API Key
module "cloudflare_create_api_key" {
  source = "./create_api_key"
  cloudflare_api_token = var.cloudflare_api_token
}
output "cloudflare_create_api_key_output" {
  value = module.cloudflare_create_api_key
}