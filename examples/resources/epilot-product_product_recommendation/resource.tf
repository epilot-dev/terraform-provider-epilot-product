# resource "epilot-product_product_recommendation" "my_productrecommendation" {
#   additional = {
#     key = jsonencode("value")
#   }
#   files = {
#     dollar_relation = [
#       {
#         entity_id = "123e4567-e89b-12d3-a456-426614174000"
#         tags = [
#           "..."
#         ]
#       }
#     ]
#   }
#   manifest = [
#     "123e4567-e89b-12d3-a456-426614174000"
#   ]
#   offers = [
#     {
#       price_id   = "...my_price_id..."
#       product_id = "...my_product_id..."
#       target_id  = "...my_target_id..."
#     }
#   ]
#   purpose = [
#     "..."
#   ]
#   schema = "product_recommendation"
#   source_price = {
#     dollar_relation = [
#       {
#         dollar_relation = [
#           {
#             entity_id = "123e4567-e89b-12d3-a456-426614174000"
#             tags = [
#               "..."
#             ]
#           }
#         ]
#       }
#     ]
#   }
#   source_product = {
#     dollar_relation = [
#       {
#         dollar_relation = [
#           {
#             entity_id = "123e4567-e89b-12d3-a456-426614174000"
#             tags = [
#               "..."
#             ]
#           }
#         ]
#       }
#     ]
#   }
#   tags = [
#     "..."
#   ]
#   type = "change"
# }

terraform {
  required_providers {
    epilot-product = {
      source  = "epilot-dev/epilot-product"
      version = "0.15.1"
    }
  }
}

provider "epilot-product" {
  # Configuration options

  epilot_auth="eyJraWQiOiJyd1wvSUdNeXJIU1wvV1wvMHlkOWoyT3Z6eURxTDEwM1RCUUNaUkVBejVMSUpBPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiI3YmNjMzE2Ni03YmRhLTRmN2EtOTdjZi1mYjgxODAxYjUzNWUiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLmV1LWNlbnRyYWwtMS5hbWF6b25hd3MuY29tXC9ldS1jZW50cmFsLTFfNGh3VXlSWkVCIiwiY3VzdG9tOml2eV9vcmdfaWQiOiI3MzkyMjQiLCJjb2duaXRvOnVzZXJuYW1lIjoibi5nb2VsQGVwaWxvdC5jbG91ZCIsImN1c3RvbTppdnlfdXNlcl9pZCI6IjExMDAwMDM3IiwiYXVkIjoiNDdwcjdzdDdsNHVtYm1wZmJpanY2N21odWEiLCJldmVudF9pZCI6IjZiYzYzOGNhLWNlYTQtNGU1Ny05NjM0LWUwMGJhMTExMTRlZiIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNzU4MDI4OTcxLCJleHAiOjE3NTgxMjM4MDgsImlhdCI6MTc1ODEyMDIwOCwiZW1haWwiOiJuLmdvZWxAZXBpbG90LmNsb3VkIn0.ShSKJIMTmb-ySXs4v8bkgFuqAKtyQEcvoLjEOhyL82g9Z4DlJGU0wQrpj0VWjsYbcfueXd0bZ2c7hKvye7qx2rOkFQ7MGDYgAszWM0VBv58s-kAuOvPHMENwtAfze4HAsjrqykhvjGlEnu70IDO7Z2MNlH10mdsLWXLOWTgYxJDvdPA8c0U3Mb2qiSZsMnjZ4IwCL-1RMt2K4W-Jx-9npQuzlWnXaoD-bIOKioit6y_v97U9mWff2T9eRnZen2BQkvQ0F1d1IyP1Ydsn4fCDO7Ng1-90muX2D0CUCbpw_86F141U64v4xA89wc1cMbg8ok6daj6vg-AincMUJrJ_vw"
  server_url="https://product.dev.sls.epilot.io"
}


# import {
#   to = epilot-product_product_recommendation.my_epilot-product_product_recommendation
#   id = "22c134b7-8dd8-4872-a5c1-8fc6dd4ac2ee"
# }

resource "epilot-product_product_recommendation" "my_epilot-product_product_recommendation" {
  # id = "22c134b7-8dd8-4872-a5c1-8fc6dd4ac2ee"
}