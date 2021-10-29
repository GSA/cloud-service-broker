
# Modified from upstream, 
# https://github.com/cloudfoundry-incubator/cloud-service-broker/blob/main/Dockerfile
FROM golang:1.17-alpine AS build
RUN apk update
RUN apk upgrade
RUN apk add --update gcc g++
WORKDIR /app

ADD https://github.com/cloudfoundry-incubator/cloud-service-broker/archive/refs/tags/0.4.0.tar.gz /app/source.tar.gz
RUN tar -xvf source.tar.gz && \
  cd cloud-service-broker-0.4.0 && \
  CGO_ENABLED=1 GOOS=linux go build -o ./build/cloud-service-broker


FROM alpine/k8s:1.18.2

COPY --from=build /app/cloud-service-broker-0.4.0/build/cloud-service-broker /bin/cloud-service-broker

# Install git so we can use it to grab Terraform modules
RUN apk add --update git

ENV PORT 8080
EXPOSE 8080/tcp

WORKDIR /bin
ENTRYPOINT ["/bin/cloud-service-broker"]
CMD ["help"]
