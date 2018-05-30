

IN="./exB.3dd"

rm ./a.out
rm ./main
rm *.out
rm *.plt
rm *.out.expected 
rm Kd
rm Ks
rm Md

CODE_PATH="$GOPATH/src/github.com/Konstantin8105/History_frame3DD/src"
RESULT="result.txt"

# Remove all text in result file
echo "" > $RESULT

# -----------------------------
echo "Compilation"

gcc "-I$CODE_PATH/viewer/" "-I$CODE_PATH/microstran/" "$CODE_PATH/main.c" "$CODE_PATH/frame3dd.c" "$CODE_PATH/frame3dd_io.c" "$CODE_PATH/coordtrans.c" "$CODE_PATH/eig.c" "$CODE_PATH/HPGmatrix.c" "$CODE_PATH/HPGutil.c" "$CODE_PATH/NRutil.c" -lm

echo "Execution"
./a.out

# Expected
cp "$IN.out" "$IN.out.expected" 

# -----------------------------
echo "Transpilation"
# -clang-flag="-lm" 
c4go transpile -o="main.go" -clang-flag="-I$CODE_PATH/viewer/" -clang-flag="-I$CODE_PATH/microstran/" "$CODE_PATH/main.c" "$CODE_PATH/frame3dd.c" "$CODE_PATH/frame3dd_io.c" "$CODE_PATH/coordtrans.c" "$CODE_PATH/eig.c" "$CODE_PATH/HPGmatrix.c" "$CODE_PATH/HPGutil.c" "$CODE_PATH/NRutil.c"

echo "Amount lines into file main.go: " >> $RESULT 
cat ./main.go | wc -l >> $RESULT
# -----------------------------
echo "Go build"
go build main.go > go_build_result.txt 2>&1

echo "Amount lines into file go_build_result.txt: " >> $RESULT 
cat ./go_build_result.txt | wc -l >> $RESULT
# -----------------------------
echo "Go vet"
go vet > govet.txt 2>&1

echo "Amount lines into file govet.txt: " >> $RESULT 
cat ./govet.txt | wc -l >> $RESULT
# -----------------------------
echo "Golint"
golint > golint.txt 2>&1

cat ./golint.txt | grep -v "should have comment or be unexported" | grep -v "use ALL_CAPS in Go names" | grep -v "use underscores in Go names" > /tmp/golint.txt
cat /tmp/golint.txt > golint.txt

echo "Amount lines into file golint.txt: " >> $RESULT 

cat ./golint.txt |  wc -l >> $RESULT
# -----------------------------
echo "Deadcode"
deadcode > deadcode.txt 2>&1

echo "Amount lines into file deadcode.txt: " >> $RESULT 
cat ./deadcode.txt | wc -l >> $RESULT
# -----------------------------
echo "Gosimple"
gosimple > gosimple.txt 2>&1

echo "Amount lines into file gosimple.txt: " >> $RESULT 
cat ./gosimple.txt | wc -l >> $RESULT
# -----------------------------
echo "Write warnings into file"
cat ./main.go | grep "// Warning" > warning.txt

echo "Write amount warnings into file"
cat ./main.go | grep "// Warning" | wc -l > warning_amount.txt

echo "Amount lines into file warning.txt: " >> $RESULT 
cat ./warning.txt | wc -l >> $RESULT
# -----------------------------

# Run
rm ./a.out
rm *.out
rm *.plt
rm Kd
rm Ks
rm Md

cat "" > "$IN.out" 
./main

diff "$IN.out" "$IN.out.expected" 
diff "$IN.out" "$IN.out.expected" > result_diff.txt
