# Copyright IBM Corp. 2020, 2025
# SPDX-License-Identifier: MPL-2.0

set -e
clear
go fmt ./...
goimports -w */*.go
for f in */
do
    pushd $f > /dev/null
    go mod tidy
    go test .
    popd > /dev/null
done

git diff --exit-code --ignore-space-change --ignore-space-at-eol
