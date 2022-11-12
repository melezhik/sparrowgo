set -e
rm -rf go.sum
find examples/ -name "go.mod" -exec rm {} \;
find examples/ -name "go.sum" -exec rm {} \;
find examples/ -name "*.bin" -exec rm {} \;
