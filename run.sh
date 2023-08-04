#!/bin/bash

rm -rf .cover
mkdir .cover
for SUBPKG in $(go list -mod=vendor ./...) ; do \
    FILE=$${SUBPKG##$(PKG)/} ; go test \
    -covermode="set" \
    -coverprofile=".cover/$${FILE////-}.cover" \
    "$$SUBPKG" | tee -a .cover/stdout.txt || exit 1; \
done
echo "mode: set" > .cover/cover.out
grep -h -v "^mode:" .cover/*.cover >> .cover/cover.out