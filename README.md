# epilot-product

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
</div>

<!-- Start SDK SDK Installation -->
## SDK Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    epilot-product = {
      source  = "epilot-dev/epilot-product"
      version = "0.2.0"
    }
  }
}

provider "epilot-product" {
  # Configuration options
}
```
<!-- End SDK SDK Installation -->

## SDK Example Usage

<!-- Start SDK SDK Example Usage -->
### Testing the provider locally

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

### Example

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```
<!-- End SDK SDK Example Usage -->


<!-- Start SDK SDK Available Operations -->

<!-- End SDK SDK Available Operations -->



<!-- Start Summary [summary] -->
## Summary

Product API: This API allows managing products, prices, taxes, and coupons.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [epilot-product](#epilot-product)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Installation](#installation)
  * [Testing the provider locally](#testing-the-provider-locally)
  * [Authentication](#authentication)
  * [Available Resources and Data Sources](#available-resources-and-data-sources)

<!-- End Table of Contents [toc] -->

<!-- Start Installation [installation] -->
## Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    epilot-product = {
      source  = "epilot-dev/epilot-product"
      version = "0.18.2"
    }
  }
}

provider "epilot-product" {
  server_url = "..." # Optional
}
```
<!-- End Installation [installation] -->



<!-- Start Testing the provider locally [usage] -->
## Testing the provider locally

#### Local Provider

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```

#### Compiled Provider

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

1. Execute `go build` to construct a binary called `terraform-provider-epilot-product`
2. Ensure that the `.terraformrc` file is configured with a `dev_overrides` section such that your local copy of terraform can see the provider binary

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/epilot-dev/epilot-product" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```
<!-- End Testing the provider locally [usage] -->

<!-- Start Authentication [security] -->
## Authentication

This provider supports authentication configuration via provider configuration.

Available configuration:

| Provider Attribute | Description |
|---|---|
| `epilot_auth` | Authorization header with epilot OAuth2 bearer token. |
| `epilot_org` | Overrides the target organization to allow shared tenantaccess. |
<!-- End Authentication [security] -->

<!-- Start Available Resources and Data Sources [operations] -->
## Available Resources and Data Sources

### Resources

* [epilot-product_coupon](docs/resources/coupon.md)
* [epilot-product_price](docs/resources/price.md)
* [epilot-product_product](docs/resources/product.md)
* [epilot-product_product_recommendation](docs/resources/product_recommendation.md)
* [epilot-product_tax](docs/resources/tax.md)
### Data Sources

* [epilot-product_coupon](docs/data-sources/coupon.md)
* [epilot-product_price](docs/data-sources/price.md)
* [epilot-product_product](docs/data-sources/product.md)
* [epilot-product_product_recommendation](docs/data-sources/product_recommendation.md)
* [epilot-product_tax](docs/data-sources/tax.md)
<!-- End Available Resources and Data Sources [operations] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/github.com/epilot-dev/terraform-provider-epilot-product/scaffolding" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

Your `<PATH>` may vary depending on how your Go environment variables are configured. Execute `go env GOBIN` to set it, then set the `<PATH>` to the value returned. If nothing is returned, set it to the default location, `$HOME/go/bin`.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)