terraform {
  required_providers {
    epilot-product = {
      source  = "epilot-dev/epilot-product"
      version = "0.18.0"
    }
  }
}

provider "epilot-product" {
  server_url = "..." # Optional
}