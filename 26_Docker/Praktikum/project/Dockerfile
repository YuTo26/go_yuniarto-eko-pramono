FROM golang:1.19-alpine
WORKDIR /apk
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /project_1
EXPOSE  80
CMD ["/project_1"]
