FROM golang AS builder

WORKDIR /yaml2yaml

COPY . .

RUN go build yaml2yaml.go

FROM alpine

COPY --from=builder /yaml2yaml/yaml2yaml /usr/local/bin/yaml2yaml

ENTRYPOINT [ "/usr/local/bin/yaml2yaml" ]
