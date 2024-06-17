install-spec-package:
	cd spec && npm install

tsp-compile:
	cd spec && tsp compile . --output-dir ../tsp-output

go-generate:
	cd server-go && go generate ./...

go-build:
	cd server-go && go build -o server .

run-go: install-spec-package tsp-compile go-generate go-build
	cd server-go && ./server

jvm-generate:
	cd server-jvm && ./mvnw clean compile

run-jvm: install-spec-package tsp-compile
	cd server-jvm && ./mvnw spring-boot:run

run-portman:
	# Port OpenAPI Spec to Postman Collection
	npx @apideck/portman -l tsp-output/@typespec/openapi3/openapi.v1.yaml -o events-postman-collection.json -b http://localhost:8084 -t false