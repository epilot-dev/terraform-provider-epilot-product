resource "epilot-product_price" "my_price" {

}

terraform {
  required_providers {
    epilot-product = {
      source  = "epilot-dev/epilot-product"
      version = "0.7.0"
    }
  }
}

variable epilot_auth {
  type        = string
}

variable server_url {
  type        = string
}


provider "epilot-product" {
  # Configuration options
  epilot_auth= var.epilot_auth
  server_url= var.server_url
}