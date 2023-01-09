# Aftershock

A toolkit leveraging Terraform to quickly query SaaS providers or create SaaS resources. 

There are countless SaaS providers and each has their own authentication standards and APIs. Why recreate the wheel each time you need to perform some action when others have already done the work with Terraform Provider integrations? Rather than learn the intricacies of some API or build custom connectors, just write or (largely) copy paste a quick HCL script and drop in your credentials to execute your desired recon or post exploit actions.

Find an API key or creds to a SaaS provider? Plug it into the `terraform.tfvars` file and see what it can do. Useful for offensive security to quickly pull data from a provider or create a user/resource and for detection engineering and blue teams to do the same to audit logs of those actions.

## Usage

**Download**

Download the repo: `git clone https://github.com/thesubtlety/aftershock.git`

Get the binary: https://github.com/thesubtlety/aftershock/releases

**Or build it**

`./build/build.sh . aftershock`

**Run**

1. Add your tokens or credentials to the appropriate providers' `terraform.tfvars.example` file.
   1. Or set env vars prepended with `TF_VAR` like: `TF_VAR_password=$(echo $ATLASSIAN_TOKEN) aftershock run atlassian create-jira-user`
   2. Some providers automatically look for well known environment variables such as `AWS_SECRET_ACCESS_KEY` or `GITHUB_TOKEN`
2. Rename to `terraform.tfvars`, and then execute your desired module.
3. Then execute, for example: `aftershock run github recon`

```
% aftershock -h
A tool to quickly leverage terraform providers and modules
Examples
        aftershock run github recon
        aftershock run splunk create-user

Usage:
  aftershock [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List available terraform providers and actions
  new         Add terraform template files for a new provider or action
  run         Run a provider action

Flags:
      --clean                   Delete provider terraform state files
      --debug-path string       TF_LOG=DEBUG output file
      --destroy                 Destroy (undo) what the last plan you applied
      --force                   Don't prompt to delete state files
  -h, --help                    help for aftershock
      --ignore-path             Ignore your installed version of terraform
      --providers-path string   Path to terraform templates folder (default "providers")
      --verbose                 Print standard terraform output
```

## Supported Providers and Actions

These are simply the actions I chose to copy paste for quick testing. As long as the resource or data source is defined you can copy paste the terraform definition and add the functionality.

* Atlassian
  * Recon
    * JQL search
  * Create User
* Splunk
  * Create user
* GitHub
  * Recon
    * get user information
    * get private repo info
  * Create SSH Key
* GitLab
  * Recon
    * get user information
* Cloudflare
  * Create API Key
  * Recon
    * Get account data
    * Get all zones
* UltraDNS
  * Recon
    * Get DNS Zone info
* Okta
  * Recon
    * Get users, groups, trusted origins
  * Create admin user
* Bitbucket
  * Recon
    * Get current user
* AWS
  * Recon
    * sts get-caller-identity
* 1Password
  * Recon
    * get user info
* Google Workspace
  * Recon
    * Get workspace users
  * Create User
* Salesforce
  * Get user profile id

## Other Providers
* Lots(!) more
  - https://registry.terraform.io/browse/providers

## Contributing or adding a new provider

1. Search the Terraform [registry](https://registry.terraform.io/)
2. Check out the documentation tab for example usage
   1. Typically the main page has provider/auth info you can copy paste into `provider.tf`. Click `Use Provider` for a template.
   2. The docs are broken out into Data Sources (good for recon) and Resources (good for post exploit actions)
3. Create templates

```
aftershock new <provider> <action>
e.g. aftershock new jamf dump-scripts
```

3. Update the terraform files based on the docs

If you're just adding an action to an existing provider all you need to do is configure the `<action>.tf` file and add any required variables to the `variables.tf` file.

**provider.tf**

Copy paste the provider auth example into `provider.tf`. You'll want to replace the auth strings with `var.api_key` for example and add that variable to `variables.tf`

**variables.tf**

Used for variables passed from the `terraform.tfvars` file, referenced in the `provider.tf` or action terraform files

**terraform.tf**

Variables used by terraform

You'll rename/copy the `terraform.tfvars.example` file to `terraform.tfvars`. If there's a better way to keep .tfvars out of git let me know.

**\<action>.tf**

Now update the `action.tf` file as appropriate, to create resources or output data

**Caveats**

Some providers don't have Data Sources which means you'd need to write one to perform recon or rely on an SDK. Typically you'll still have resources available for post exploit actions.

## Troubleshooting

Your state files may get into a broken state, especially if you're adding new actions

Try to force clean the tf files with `aftershock run <provider> <action> --clean --force`

Use the `terraform` binary for more fine grained actions

## References
- https://registry.terraform.io/
- https://developer.hashicorp.com/terraform/tutorials/providers
