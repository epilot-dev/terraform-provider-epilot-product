resource "epilot-product_product_recommendation" "my_productrecommendation" {
  additional = {
    key = jsonencode("value")
  }
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
  offers = [
    {
      price_id   = "...my_price_id..."
      product_id = "...my_product_id..."
      target_id  = "...my_target_id..."
    }
  ]
  purpose = [
    "..."
  ]
  schema = "product_recommendation"
  source_price = {
    dollar_relation = [
      {
        dollar_relation = [
          {
            entity_id = "123e4567-e89b-12d3-a456-426614174000"
            tags = [
              "..."
            ]
          }
        ]
      }
    ]
  }
  source_product = {
    dollar_relation = [
      {
        dollar_relation = [
          {
            entity_id = "123e4567-e89b-12d3-a456-426614174000"
            tags = [
              "..."
            ]
          }
        ]
      }
    ]
  }
  tags = [
    "..."
  ]
  type = "change"
}