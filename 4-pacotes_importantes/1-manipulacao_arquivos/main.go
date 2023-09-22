package main

import (
	"os"
)

func main() {

	//Escrita de Arquivo

	//f, err := os.Create("arquivo.txt")
	//if err != nil {
	//	panic(err)
	//}
	//
	////Se eu não souber o tipo de dados que eu vou gravar eu passo o Write com parâmetro []bytes
	//tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	//
	////Se eu souber que meus dados são Strings posso usar o WriteString
	////tamanho, err := f.WriteString("Hello, World!")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
	//f.Close()

	//Leitura de Arquivo

	//arquivo, err := os.ReadFile("arquivo.txt")
	//if err != nil {
	//	panic(err)
	//}
	////convertendo o conteudo do arquivo ([]bytes) para string
	//fmt.Println(string(arquivo))

	//Como ler um arquivo muito grande tipo de 1GB....
	//No GO temos como ler um arquivo por partes...

	//leitura de pouco em pouco abrindo o arquivo
	//arquivo2, err := os.Open("arquivo.txt")
	//if err != nil {
	//	panic(err)
	//}
	//
	////O pacote bufio bufferiza.... definimos aqui que a leitura será de 10 em 10 bytes
	//reader := bufio.NewReader(arquivo2)
	//buffer := make([]byte, 3)
	//
	//for {
	//	n, err := reader.Read(buffer)
	//	if err != nil {
	//		break
	//	}
	//	fmt.Println(string(buffer[:n]))
	//}

	//Removendo arquivo

	err := os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
