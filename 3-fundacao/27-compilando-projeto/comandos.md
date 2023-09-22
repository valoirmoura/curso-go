go build nome_arquivo.go

//Aqui excolhemos o sistema operacional que ira rodar nosso Binario
GOOS=linux go build nome_arquivo.go 
GOOS=windows go build nome_arquivo.go 
GOOS=windows GOARCH=amd64 go build nome_arquivo.go //nesse caso executamentos para windows na plataforma 64bits
GOOS=windows GOARCH=386 go build nome_arquivo.go //nesse caso executamentos para windows na plataforma 64bits


comando legal para verificar todas as plataformas: go tool dist list
