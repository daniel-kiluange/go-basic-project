FROM golang:1.16-alpine as build
RUN apk add --no-cache git
RUN apk add --no-cache make
RUN mkdir "/src"
WORKDIR /src
COPY . .
RUN make

FROM scratch
COPY --from=build /src/bin /app/app
ENTRYPOINT ["/app/app/api"]
