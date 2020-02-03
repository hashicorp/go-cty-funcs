go fmt ./... 
goimports -w */*.go 
for f in */
do
    pushd $f > /dev/null
    go mod tidy
    go test .
    popd > /dev/null
done