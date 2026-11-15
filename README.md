# http-sample

API HTTP mínima para demonstrar o fluxo de build do devex.

```sh
go test ./...
go run .
curl http://127.0.0.1:8080/
curl http://127.0.0.1:8080/healthz
```

Para verificar que o teste pode reprovar de forma controlada:

```sh
DEVEX_TEST_FORCE_FAIL=1 go test ./...
```

O `.devex.yaml` declara o comando de teste e os caminhos de build que o worker lerá do commit remoto. Este diretório é a fonte da aplicação de demonstração no monorepo; para o fluxo ponta a ponta, seu conteúdo também precisa existir em um repositório GitHub público de teste.
