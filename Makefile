install-spec-package:
	cd spec && npm install

tsp-compile:
	cd spec && tsp compile . --output-dir ../tsp-output

go-generate:
	cd server-go && go generate ./...

go-build:
	cd server-go && go build -o server .