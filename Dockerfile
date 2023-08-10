FROM golang:1.20.6

LABEL maintainer="Jorge <jorgewpgomes@gmail.com>"

WORKDIR /app/src/ChallangeStone

# aponta a variavel gopath do go para o diretorio app
ENV GOPATH=/app

# copia os arquivos do projeto para o workdir do container
COPY . /app/src/ChallangeStone

# execulta o main.go e baixa as dependencias do projeto
RUN go build main.go

# Comando para rodar o executavel
ENTRYPOINT ["./main"]

# expõe a pota 8080
EXPOSE 8080