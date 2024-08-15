# __generated__ by Terraform
# Please review these resources and move them into your main configuration files.

# __generated__ by Terraform from "2b7f9439-21bc-450b-91b4-aa82e6d3e53b"
resource "epilot-product_product" "my_product" {
  active        = true
  code          = null
  description   = null
  feature       = null
  internal_name = "_Product with a lot of files"
  name          = "_Product with a lot of files"
  price_options = {
    dollar_relation = [
      {
        entity_id = "21864c72-dc60-4a5e-8e4a-c5cb5daf52ba"
        tags      = []
      },
    ]
  }
  product_downloads = jsonencode({
    "$relation" = [{
      entity_id = "c40ab6e7-7abf-4a8a-afec-20756510ed6f"
    }]
  })
  product_images = jsonencode({
    "$relation" = [{
      _tags     = []
      entity_id = "6406688d-e6ca-414f-8a31-2c3c5661af3d"
    }]
  })
  type = "product"
}
