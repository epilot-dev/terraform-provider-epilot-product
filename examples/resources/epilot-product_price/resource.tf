resource "epilot-product_price" "my_price" {
  active                  = true
  additional              = "{ \"see\": \"documentation\" }"
  billing_duration_amount = 4.23
  billing_duration_unit   = "years"
  description             = "...my_description..."
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
  is_composite_price = true
  is_tax_inclusive   = false
  long_description   = "...my_long_description..."
  manifest = [
    "123e4567-e89b-12d3-a456-426614174000"
  ]
  notice_time_amount = 7.88
  notice_time_unit   = "weeks"
  price_components = {
    dollar_relation = [
      {
        entity_id = "...my_entity_id..."
        tags = [
          "..."
        ]
      }
    ]
  }
  price_display_in_journeys = "show_price"
  pricing_model             = "tiered_graduated"
  purpose = [
    "..."
  ]
  renewal_duration_amount = 3.79
  renewal_duration_unit   = "years"
  schema                  = "price"
  tags = [
    "..."
  ]
  tax = {
    dollar_relation = [
      {
        entity_id = "123e4567-e89b-12d3-a456-426614174000"
        tags = [
          "..."
        ]
      }
    ]
  }
  termination_time_amount = 2.23
  termination_time_unit   = "years"
  tiers = [
    {
      display_mode            = "on_request"
      flat_fee_amount         = 1.58
      flat_fee_amount_decimal = "...my_flat_fee_amount_decimal..."
      unit_amount             = 4.92
      unit_amount_decimal     = "...my_unit_amount_decimal..."
      up_to                   = 5.32
    }
  ]
  type                 = "one_time"
  unit                 = "...my_unit..."
  unit_amount          = 6.37
  unit_amount_currency = "EUR"
  unit_amount_decimal  = "...my_unit_amount_decimal..."
  variable_price       = false
}