# Get caller identity
data "aws_caller_identity" "current" {}
output "aws_caller_identity_output" {
  description = "Get caller identity info"
  value = "${data.aws_caller_identity.current}"
}
