resource "epilot-product_product" "my_product" {
  active = false
  additional = {
    "see" : jsonencode("documentation"),
  }
  availability_files = {
    dollar_relation = [
      {
        entity_id = "123e4567-e89b-12d3-a456-426614174000"
        tags = [
          "..."
        ]
      }
    ]
  }
  code        = "...my_code..."
  description = "...my_description..."
  feature = [
    "{ \"see\": \"documentation\" }"
  ]
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
  internal_name = "...my_internal_name..."
  manifest = [
    "123e4567-e89b-12d3-a456-426614174000"
  ]
  name = "...my_name..."
  price_options = {
    dollar_relation = [
      {
        entity_id = "123e4567-e89b-12d3-a456-426614174000"
        tags = [
          "..."
        ]
      }
    ]
  }
  product_downloads = {
    dollar_relation = [
      {
        entity_id = "123e4567-e89b-12d3-a456-426614174000"
        tags = [
          "..."
        ]
      }
    ]
  }
  product_images = {
    dollar_relation = [
      {
        entity_id = "123e4567-e89b-12d3-a456-426614174000"
        tags = [
          "..."
        ]
      }
    ]
  }
  purpose = [
    "..."
  ]
  schema = "product"
  tags = [
    "..."
  ]
  type = "service"
}