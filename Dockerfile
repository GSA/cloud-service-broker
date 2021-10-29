
FROM alpine/k8s:1.18.2

ADD https://github.com/cloudfoundry-incubator/cloud-service-broker/releases/download/0.4.0/cloud-service-broker.darwin /bin/cloud-service-broker

RUN chmod 755 /bin/cloud-service-broker

# Install git so we can use it to grab Terraform modules
RUN apk add --update git

ENV PORT 8080
EXPOSE 8080/tcp

WORKDIR /bin
ENTRYPOINT ["/bin/cloud-service-broker"]
CMD ["help"]
