resource "epilot-product_tax" "my_tax" {
  active      = false
  description = "...my_description..."
  rate        = "...my_rate..."
  region      = "AT"
  tax_id      = "123e4567-e89b-12d3-a456-426614174000"
  type        = "Custom"
}