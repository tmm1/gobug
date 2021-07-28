FROM arm32v7/alpine
WORKDIR /usr/src
RUN apk add go git gcc make musl-dev
ADD go.* *.go .
ENV GOARM=7
RUN go get
#RUN go version && go env
RUN go test -v ./
