# Provider definitions
terraform {
  required_providers {
    epilot-product = {
      source  = "epilot-dev/epilot-product"
      version = "0.10.1"
    }
  }
}

# Variables
variable "epilot_auth" {
  type = string
}

variable "product_api_url" {
  type = string
  default = "https://product.dev.sls.epilot.io"
}
# Providers configuration

provider "epilot-product" {
  epilot_auth = var.epilot_auth
  server_url = var.product_api_url
}


# resource "epilot-product_product" "my_product" {

# }