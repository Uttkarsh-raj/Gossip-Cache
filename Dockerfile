FROM golang:1.22.2-alpine

RUN mkdir /app 

ADD . /app

WORKDIR /app

COPY . /app/

EXPOSE 3000

RUN go build -o main.exe

CMD [ "/app/main.exe" ]