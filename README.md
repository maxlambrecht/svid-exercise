
go run client/main.go --cert certs/client_cert.pem --key certs/client_key.pem 

go run client/main.go --cert certs/unknown_client_cert.pem --key certs/unknown_client_key.pem

go run service/main.go --spiffeid spiffe://example.com/service --cert certs/server_cert.pem --key certs/server_key.pem