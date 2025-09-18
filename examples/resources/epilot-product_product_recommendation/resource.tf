resource "epilot-product_product_recommendation" "my_productrecommendation" {
  additional = "{ \"see\": \"documentation\" }"
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
  offers = "{ \"see\": \"documentation\" }"
  purpose = [
    "..."
  ]
  schema         = "product_recommendation"
  source_price   = "{ \"see\": \"documentation\" }"
  source_product = "{ \"see\": \"documentation\" }"
  tags = [
    "..."
  ]
  type = "change"
}