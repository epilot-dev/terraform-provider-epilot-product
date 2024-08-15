resource "epilot-product_price" "my_price" {
  schema                    = "price"
  active                    = true
  billing_duration_amount   = 66.76
  billing_duration_unit     = "weeks"
  description               = "...my_description..."
  is_composite_price        = false
  is_tax_inclusive          = true
  long_description          = "...my_long_description..."
  notice_time_amount        = 2.66
  notice_time_unit          = "months"
  price_display_in_journeys = "show_as_starting_price"
  price_id                  = "123e4567-e89b-12d3-a456-426614174000"
  pricing_model             = "tiered_graduated"
  renewal_duration_amount   = 13.87
  renewal_duration_unit     = "months"
  tax                       = "{ \"see\": \"documentation\" }"
  termination_time_amount   = 70.52
  termination_time_unit     = "weeks"
  type                      = "one_time"
  unit                      = "...my_unit..."
  unit_amount               = 62.81
  unit_amount_currency      = "EUR"
  unit_amount_decimal       = "...my_unit_amount_decimal..."
  variable_price            = true
}