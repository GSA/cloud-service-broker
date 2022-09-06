FROM cfplatformeng/csb:0.10.0 as upstream

FROM alpine/k8s:1.20.7

COPY --from=upstream /bin/cloud-service-broker /bin/cloud-service-broker

RUN apk update
RUN apk upgrade
# Install git so we can use it to grab Terraform modules
RUN apk add --update git zip

# Enable re-templates of Terraform Code
# Reference: https://github.com/GSA/data.gov/issues/3083
ENV BROKERPAK_UPDATES_ENABLED true

ENV PORT 8080
EXPOSE 8080/tcp

WORKDIR /bin
ENTRYPOINT ["/bin/cloud-service-broker"]
CMD ["help"]
