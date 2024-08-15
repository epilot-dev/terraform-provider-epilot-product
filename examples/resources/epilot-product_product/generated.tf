# __generated__ by Terraform
# Please review these resources and move them into your main configuration files.

# __generated__ by Terraform from "6f19b201-996c-4052-8fb1-df36b02e4e03"
resource "epilot-product_product" "my_product" {
  active = true
  additional = {
    new_attribute_1723730533875 = "\"test\""
  }
  availability_files = null
  code               = null
  description        = null
  feature            = null
  files              = null
  internal_name      = "Test product with simple price"
  name               = "Test product with simple price"
  price_options = {
    dollar_relation = [
      {
        entity_id = "2a1dbfe9-8b65-4499-88d9-554cedaa253c"
        tags      = []
      },
    ]
  }
  product_downloads = null
  product_images    = null
  purpose           = []
  schema            = "product"
  tags              = []
  type              = "product"
}
