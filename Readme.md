

Banking App

Este é um aplicativo de serviços bancários desenvolvido em Go usando o framework Proto.Actor e Gorilla Mux. O aplicativo fornece uma API para realizar transações bancárias, consultar resumos diários. 

Pré-requisitos

Certifique-se de ter as seguintes ferramentas instaladas em seu ambiente:

Go (versão 1.16 ou superior)
Docker (opcional, se desejar executar o aplicativo em um contêiner Docker)

Instalação

Clone este repositório para o seu diretório de trabalho:

bash

git clone https://github.com/cpsleme/LemeBank.git

Navegue até o diretório do projeto:

bash

cd banking-app

Instale as dependências do Go usando o go mod:

bash

go mod download

Execução

Existem duas opções para executar o aplicativo: diretamente no ambiente Go ou em um contêiner Docker. Executando diretamente no ambiente Go

No diretório do projeto, execute o seguinte comando para iniciar o aplicativo:

bash

go run main.go

Isso compilará e executará o aplicativo no ambiente Go. O aplicativo estará acessível em http://localhost:8080.

Executando em um contêiner Docker

Certifique-se de ter o Docker instalado em sua máquina.

No diretório do projeto, crie uma imagem do Docker executando o seguinte comando:

bash

docker build -t banking-app .

Isso criará uma imagem do Docker com o nome banking-app.

Após a criação da imagem, inicie o contêiner com o seguinte comando:

bash

docker run -p 8080:8080 banking-app

Isso iniciará o contêiner e mapeará a porta 8080 do contêiner para a porta 8080 do host. O aplicativo estará acessível em http://localhost:8080.

Uso da API

Uma vez que o aplicativo esteja em execução, você pode usar a API para realizar as seguintes operações:

Lançamento de débito: POST /debit
Lançamento de crédito: POST /credit
Recuperação do resumo diário: GET /daily-summary?accountnumber=<account_number>&date=<date>

Certifique-se de substituir <account_number> e pelos valores adequados ao fazer as solicitações.

Este projeto está licenciado sob a licença MIT. Consulte o arquivo `LICENSE
