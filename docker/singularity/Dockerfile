FROM golang:latest as builder

# Create directories.
RUN mkdir /build
RUN mkdir /files

# Build statically linked binary.
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o pep2gene .

# Create apline.
FROM alpine
COPY --from=builder /build/pep2gene /app/
COPY --from=builder /files /files

# Perform work in directory where user files will be located.
WORKDIR /files

ENTRYPOINT ["/app/pep2gene"]
