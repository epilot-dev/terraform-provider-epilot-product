resource "epilot-product_tax" "my_tax" {
  schema      = "tax"
  active      = false
  description = "...my_description..."
  rate        = "...my_rate..."
  region      = "DE"
  tax_id      = "123e4567-e89b-12d3-a456-426614174000"
  type        = "Custom"
}