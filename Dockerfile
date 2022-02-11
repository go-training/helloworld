FROM scratch
COPY helloworld /helloworld
ENTRYPOINT ["/helloworld"]
