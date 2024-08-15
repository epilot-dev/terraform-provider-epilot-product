terraform {
  required_providers {
    epilot-product = {
      source  = "epilot-dev/epilot-product"
      version = "0.10.1"
    }
  }
}

variable "epilot_auth" {
  type        = string
  description = "epilot_auth"
}


provider "epilot-product" {
  # Configuration options
  epilot_auth = var.epilot_auth
  server_url  = "https://product.dev.sls.epilot.io"
}
