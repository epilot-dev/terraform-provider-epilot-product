resource "epilot-product_tax" "my_tax" {
  active      = true
  additional  = "{ \"see\": \"documentation\" }"
  description = "...my_description..."
  files = {
    dollar_relation = [
      {
        entity_id = "123e4567-e89b-12d3-a456-426614174000"
        tags = [
          "..."
        ]
      }
    ]
  }
  manifest = [
    "123e4567-e89b-12d3-a456-426614174000"
  ]
  purpose = [
    "..."
  ]
  rate   = "{ \"see\": \"documentation\" }"
  region = "DE"
  schema = "tax"
  tags = [
    "..."
  ]
  type = "VAT"
}