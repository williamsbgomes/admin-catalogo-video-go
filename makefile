# Nome do arquivo de cobertura
COVERAGE=coverage.out

# Diretório dos pacotes (ajuste se necessário)
PKG=./...

.PHONY: test coverage html lint fmt clean

## Roda todos os testes com saída detalhada
test:
	go test -v $(PKG)

## Roda testes com cobertura e salva no arquivo coverage.out
coverage:
	go test -coverprofile=$(COVERAGE) $(PKG)

## Abre relatório de cobertura no navegador
html: coverage
	go tool cover -html=$(COVERAGE)

## Formata o código com go fmt
fmt:
	go fmt $(PKG)

## Roda golint (precisa instalar com: go install golang.org/x/lint/golint@latest)
lint:
	golint ./...

## Remove arquivos de cobertura
clean:
	rm -f $(COVERAGE)