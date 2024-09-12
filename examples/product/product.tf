# epilot-product_product.my_product:
resource "epilot-product_product" "my_product" {
    active         = true
    additional     = {
        "_acl_sync"    = "\"2024-06-29T19:59:22.804Z\""
        "order_number" = "\"PRODUCT-1029\""
    }
    internal_name  = "A normal product"
    manifest       = ["a4890a91-013a-4354-a0ba-5c8fc0d22cae", "92e0d59b-437c-4da8-a1d5-ef2fdbaab543"]
    name           = "A normal product"
    price_options  = {
        dollar_relation = [
            {
                entity_id = "f107dfd7-4bef-4c3e-9a08-c8d3746b9168"
                tags      = []
            },
        ]
    }
    product_images = {
        dollar_relation = [
            {
                entity_id = "64414958-2744-40af-8ccb-4a49111e25bd"
                tags      = []
            },
        ]
    }
    purpose        = []
    schema         = "product"
    tags           = []
    type           = "product"
}
