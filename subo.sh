git add .
git commit -m "Último Commit"
git push
GOOS=linux GOARCH=amd64  go build -o bootstrap main.go
rm main.zip
zip main.zip main
zip myFunction.zip bootstrap