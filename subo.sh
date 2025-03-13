git add .
git commit -m "Último Commit"
git push
go build -o bootstrap main.go
rm main.zip
zip main.zip main