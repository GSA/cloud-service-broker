FROM alpine/k8s:1.20.7 AS upstream

RUN apk update && apk upgrade && apk add --update wget

RUN wget --no-verbose \
    https://github.com/cloudfoundry/cloud-service-broker/releases/download/v2.5.6/cloud-service-broker.linux \
    -O /bin/cloud-service-broker

FROM alpine/k8s:1.20.7

COPY --from=upstream /bin/cloud-service-broker /bin/cloud-service-broker
RUN chmod a+x /bin/cloud-service-broker

# Install git so we can use it to grab Terraform modules
RUN apk update && apk upgrade && apk add --update git zip


# Enable re-templates of Terraform Code
# Reference: https://github.com/GSA/data.gov/issues/3083
ENV BROKERPAK_UPDATES_ENABLED=true

ENV PORT=8080
EXPOSE 8080/tcp

WORKDIR /bin
ENTRYPOINT ["/bin/cloud-service-broker"]
CMD ["help"]
