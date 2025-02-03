resource "epilot-product_coupon" "my_coupon" {
  active = true
  additional = {
    key = jsonencode("value"),
  }
  cashback_period = "12"
  category        = "discount"
  description     = "...my_description..."
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
  fixed_value          = 3.79
  fixed_value_currency = "EUR"
  fixed_value_decimal  = "...my_fixed_value_decimal..."
  manifest = [
    "123e4567-e89b-12d3-a456-426614174000"
  ]
  name             = "...my_name..."
  percentage_value = "...my_percentage_value..."
  prices = {
    dollar_relation = [
      {
        entity_id = "123e4567-e89b-12d3-a456-426614174000"
        tags = [
          "..."
        ]
      }
    ]
  }
  promo_code_usage = "{ \"see\": \"documentation\" }"
  promo_codes = [
    {
      code            = "...my_code..."
      has_usage_limit = false
      id              = "...my_id..."
      usage_limit     = 8.12
    }
  ]
  requires_promo_code = false
  schema              = "coupon"
  tags = [
    "..."
  ]
  type = "fixed"
}