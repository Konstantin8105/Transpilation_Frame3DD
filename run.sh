
CODE_PATH="$GOPATH/src/github.com/Konstantin8105/History_frame3DD/src"

echo "Write version c2go"
c2go -v > c2go_version.txt

echo "Transpilation"
c2go transpile -o="main.go" "$CODE_PATH/main.c" "$CODE_PATH/frame3dd.c" "$CODE_PATH/frame3dd_io.c" "$CODE_PATH/coordtrans.c" "$CODE_PATH/eig.c" "$CODE_PATH/HPGmatrix.c" "$CODE_PATH/HPGutil.c" "$CODE_PATH/NRutil.c"

echo "Write warnings into file"
cat ./main.go | grep "// Warning" > warning.txt

echo "Write amount warnings into file"
cat ./main.go | grep "// Warning" | wc -l > warning_amount.txt

