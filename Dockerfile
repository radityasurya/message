FROM scratch

# ADD ca-certificates.crt /etc/ssl/certs/

COPY ./bin/message /

CMD ["/message", "serve"]
