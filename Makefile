.PHONY: all docs
all: speakeasy docs

original.yaml: product.yaml
#cp product.yaml original.yaml
	curl https://docs.api.epilot.io/product.yaml > original.yaml


original_modified.yaml: original.yaml overlay.yaml
	speakeasy overlay apply -s original.yaml -o overlay.yaml > original_modified.yaml

overlay.yaml:
	speakeasy overlay compare -s original.yaml -s original_modified.yaml > overlay.yaml

speakeasy: product.yaml
	$(eval TMP := $(shell mktemp -d))
#cp product.yaml $(TMP)/openapi.yaml
	curl https://docs.api.epilot.io/product.yaml > $(TMP)/openapi.yaml
	speakeasy overlay apply -s $(TMP)/openapi.yaml -o overlay.yaml > $(TMP)/final.yaml
	speakeasy generate sdk --lang terraform -o . -s $(TMP)/final.yaml

docs:
	go generate ./...
