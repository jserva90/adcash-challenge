FROM golang:1.19 AS build

LABEL version="1.0" Author="Jserva90"

WORKDIR /go/src/loanapp

COPY . ./

RUN go build -o /bin/loanapp

FROM gcr.io/distroless/base-debian11
COPY --from=build /bin/loanapp /loanapp
COPY data/ /data

EXPOSE 8080
ENTRYPOINT ["/loanapp"]