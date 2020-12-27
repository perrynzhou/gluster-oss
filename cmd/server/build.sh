rm -rf server
go build -mod=vendor -o server
rm -rf /mnt/dht/*