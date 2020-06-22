default: build plan

deps:
	go install github.com/hashicorp/terraform

build:
	go build -o terraform-provider-salesforce && cp ./terraform-provider-salesforce tf/terraform.d/plugins/darwin_amd64

test:
	go test -v .

plan:
	@terraform plan
