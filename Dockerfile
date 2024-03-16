FROM golang:1.21.2

WORKDIR /app

#COPY go.mod go.sum ./
#COPY go.mod .
COPY . .
RUN go mod download

#COPY *.go ./

RUN go build -o /docker-rest-api .

EXPOSE 5007

CMD ["/docker-rest-api"]

# Билдить:
# docker build --tag docker-rest-api .