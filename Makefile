version := 99.99.0
provider_macos_path = registry.terraform.io/KeisukeYamashita/dummy/$(version)/darwin_arm64/

.PHONY: doc
doc:
	@go generate ./...
	@rsync -a docs/ website/docs/configurations

.PHONY: build
build: 
	@go build

.PHONY: install_macos
install_macos: build
	@mkdir -p ~/Library/Application\ Support/io.terraform/plugins/$(provider_macos_path)
	@mv terraform-provider-dummy ~/Library/Application\ Support/io.terraform/plugins/$(provider_macos_path)