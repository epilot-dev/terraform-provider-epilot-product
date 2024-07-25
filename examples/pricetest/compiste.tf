# epilot-product_price.my_price:
resource "epilot-product_price" "my_price" {
    active                    = true
    billing_duration_unit     = "months"
    description               = "Composite"
    id                        = "e79c79ee-47e6-47dc-bb0a-e37757ff0467"
    is_composite_price        = true
    is_tax_inclusive          = true
    notice_time_unit          = "months"
    price_components          = {
        dollar_relation = [
            {
                entity_id = "5cc7d072-8d81-4510-aebb-d487593d3fe4"
                tags      = []
            },
            {
                entity_id = "aef586cc-c5ec-4e19-93af-de4c9ac8357f"
                tags      = []
            },
        ]
    }
    price_display_in_journeys = "show_price"
    pricing_model             = "per_unit"
    renewal_duration_unit     = "months"
    termination_time_unit     = "months"
    tiers                     = []
    type                      = "one_time"
    unit_amount               = 0
    unit_amount_currency      = "EUR"
    unit_amount_decimal       = "0.00"
    variable_price            = false
}
