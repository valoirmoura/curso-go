Client gRPC: Evans
- package nome_do_pacote
- service nome_do_service
- call nome_do_metodo

para sair do evans: ctrl + D quando estiver usando streams


Comando para rodar ProtoBuf:
```protoc --go_out=. --go-grpc_out=. proto/course_category.proto (para gerar o arquivo entities.pb.go)

