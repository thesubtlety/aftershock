# Get detailed information for your zones
data "ultradns_zone" "zone" {
    name = var.zone_name
}

# Print it
output "ultradns_zone_info" {
  description = "ultradns_zone info"
  value = "${data.ultradns_zone.zone}"
}
