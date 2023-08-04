test:
	go test ./...

coverage: run-cover
	go tool cover -html=.cover/cover.out -o coverage.html

SHELL := /bin/bash
run-cover:
	@rm -rf .cover
	@mkdir .cover
	@for SUBPKG in $(shell go list -mod=vendor ./...) ; do \
		FILE=$${SUBPKG##$(PKG)/} ; go test \
		-covermode="set" \
		-coverprofile=".cover/$${FILE////-}.cover" \
		"$$SUBPKG" | tee -a .cover/stdout.txt || exit 1; \
	done
	echo "mode: set" > .cover/cover.out
	grep -h -v "^mode:" .cover/*.cover >> .cover/cover.out
