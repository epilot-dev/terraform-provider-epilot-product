terraform {
  required_providers {
    epilot-product = {
      source  = "epilot-dev/epilot-product"
      version = "0.18.1"
    }
  }
}

provider "epilot-product" {
  server_url = "..." # Optional
}