install-spec-package:
	cd spec && npm install

tsp-compile:
	cd spec && tsp compile . --output-dir ../tsp-output

go-generate:
	cd server-go && go generate ./...

go-build:
	cd server-go && go build -o server .

jvm-generate:
	cd server-jvm && ./mvnw clean compile

run-portman:
	# Port OpenAPI Spec to Postman Collection
	npx @apideck/portman -l tsp-output/@typespec/openapi3/openapi.v1.yaml -o events-postman-collection.json -b http://localhost:8084 -t false