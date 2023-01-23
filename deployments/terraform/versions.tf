terraform {
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "blackhorseya"

    workspaces {
      name = "irent"
    }
  }

  required_providers {
    cloudflare = {
      source = "cloudflare/cloudflare"
    }
  }
}
