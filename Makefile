install-spec-package:
	cd spec && npm install

tsp-compile:
	cd spec && tsp compile . --output-dir ../tsp-output