resource "epilot-product_product" "my_product" {
  active            = false
  code              = "...my_code..."
  description       = "...my_description..."
  internal_name     = "...my_internal_name..."
  name              = "Elbert Stamm"
  product_downloads = "{ \"see\": \"documentation\" }"
  product_images    = "{ \"see\": \"documentation\" }"
  product_id        = "123e4567-e89b-12d3-a456-426614174000"
  type              = "product"
}