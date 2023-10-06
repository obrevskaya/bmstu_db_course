
goimports:
	goimports -w .

gen-client: _gen-client goimports

_gen-client:
	echo ${PWD}
	docker run --rm \
      -v ${PWD}:/console openapitools/openapi-generator-cli generate \
      -i /console/api/swagger.yaml \
      -g go \
      -o /console/internal/client/ \
      --additional-properties=outputAsLibrary=true,sourceFolder=.