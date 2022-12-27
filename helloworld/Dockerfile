FROM golang:1.16-alpine
RUN adduser -S goruntime
USER goruntime
WORKDIR /work
COPY src-code/ /work/
CMD ["go","run","."]
EXPOSE 8000