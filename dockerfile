FROM golang:1.22.4 AS deps

WORKDIR /gcp_practice
ADD *.mod *.sum ./
RUN go mod download

FROM deps as dev
ADD . .
EXPOSE 8080
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -X main.docker=true" -o api cmd/main.go
CMD ["/gcp_practice/api"]

FROM scratch as prod

WORKDIR /
EXPOSE 8080
COPY --from=dev /gcp_practice/api /
CMD ["/api"]