FROM golang:1.19-alpine3.16

##buat folder APP
RUN mkdir /deall

##set direktori utama
WORKDIR /deall

##copy seluruh file
ADD . .

##buat executeable
RUN go build -o master .

##jalankan executeable
CMD ["./master"]