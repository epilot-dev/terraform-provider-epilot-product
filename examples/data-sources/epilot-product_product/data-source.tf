data "epilot-product_product" "my_product" {
  hydrate    = false
  product_id = "123e4567-e89b-12d3-a456-426614174000"
  strict     = true
}