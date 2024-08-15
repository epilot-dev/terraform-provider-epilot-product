data "epilot-product_price" "my_price" {
  hydrate  = true
  price_id = "123e4567-e89b-12d3-a456-426614174000"
  strict   = true
}