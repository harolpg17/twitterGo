git add .
git commit -m "Último Commit"
git push
GOOS=linux GOARCH=amd64  go build -o bootstrap main.go
rm main.zip
rm myFunction.zip
zip main.zip main
zip myFunction.zip bootstrap