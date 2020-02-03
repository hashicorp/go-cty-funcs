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
