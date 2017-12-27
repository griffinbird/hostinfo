FROM busybox:1.26
COPY ./hostinfo /app/hostinfo
WORKDIR "/app"
CMD ["./hostinfo"]