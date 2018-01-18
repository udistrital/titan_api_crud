FROM debian:stretch
COPY main /main
RUN chmod +x main
ENTRYPOINT ["/main"]
