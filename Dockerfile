FROM golang:1.21.3 AS BUILDER

ENV CGO_ENABLED=0

WORKDIR /home/zoo-mailer/src

COPY go.mod    /home/zoo-mailer/src/
COPY go.sum    /home/zoo-mailer/src/

RUN go mod download && go mod verify

COPY cmd/      /home/zoo-mailer/src/cmd/
COPY internal/ /home/zoo-mailer/src/internal/

RUN mkdir /home/zoo-mailer/out

RUN go build -o /home/zoo-mailer/out/zoo /home/zoo-mailer/src/cmd/zoo

FROM scratch

WORKDIR /opt/zoo-mailer

ENV ZOO__folder=/var/zoo/files
ENV ZOO__mail_tmpl=/etc/zoo/templates/mail.tmpl
ENV ZOO__sendlist=/etc/zoo/send.list

COPY --from=BUILDER /home/zoo-mailer/out/zoo /opt/zoo-mailer/

ENTRYPOINT [ "/opt/zoo-mailer/zoo" ]