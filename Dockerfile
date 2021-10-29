
FROM alpine/k8s:1.18.2

ADD https://github.com/cloudfoundry-incubator/cloud-service-broker/releases/download/0.4.0/cloud-service-broker.linux /bin/cloud-service-broker

RUN chmod 755 /bin/cloud-service-broker

# Install git so we can use it to grab Terraform modules
RUN apk update
RUN apk add --update git cgo

ENV PORT 8080
EXPOSE 8080/tcp

WORKDIR /bin
ENTRYPOINT ["/bin/cloud-service-broker"]
CMD ["help"]
