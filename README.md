## SVID basic validation

##### This example contains: 
 - A HTTP-based SPIFFE Verifiable Identity Document [SVID](https://github.com/spiffe/spiffe/blob/master/standards/SPIFFE-ID.md) validation service that takes a SPIFFE ID as a configuration parameter. 
 The service uses [go-spiffe](https://github.com/spiffe/go-spiffe) to parse the SVID and obtain the Subject Alternative Name
 - A client that connects over mTLS using an SVID.
     
###### The service returns:

- HTTP 200 if the URI SAN of the SVID matches the configured SPIFFE ID.
- HTTP 401 if the URI SAN does not match the configured SPIFFE ID.

##### How to run the example

###### Get the code and dependencies

```
$ go get -u github.com/maxlambrecht/svid-exercise
$ cd ~/go/src/github.com/maxlambrecht/svid-exercise
$ go get ./... 
```


###### Run the Server
```
$ go run service/main.go --spiffeid spiffe://example.com/service --cert certs/server_cert.pem --key certs/server_key.pem
Server listening on address :3000
```

By default the server listens on _https://localhost:3000_
To listen on another address use the option _--addr_

```
$ go run service/main.go --spiffeid spiffe://example.com/service --cert certs/server_cert.pem --key certs/server_key.pem --addr localhost:4000
Server listening on address localhost:4000
```

###### Run the Client

```
# using a certificate that has a Subject Alternative Name=spiffe://example.com/service:
$ go run client/main.go --cert certs/client.crt --key certs/client.key --ca certs/rootCA.crt
Response code: 200
Authentication Succeed
```

```
# using a certificate with an untrusted Subject Alternative Name or not SAN at all
$  go run client/main.go --cert certs/other_client.crt --key certs/other_client.key --ca certs/rootCA.crt 
Response code: 401
Authentication Failed. Invalid SpiffeID
```

By default the client sends the requests to _https://localhost:3000/auth_
To send the requests to another address use the option _--url_ 

```
$ go run client/main.go --cert certs/client.crt --key certs/client.key --ca certs/rootCA.crt --url https://localhost:4000/auth
Response code: 200
Authentication Succeed
```


###### Run Tests


```
$ go test ./... -v -cover

=== RUN   TestAuthenticateHandler
=== RUN   TestAuthenticateHandler/valid_ID
=== RUN   TestAuthenticateHandler/invalid_ID
--- PASS: TestAuthenticateHandler (0.00s)
    --- PASS: TestAuthenticateHandler/valid_ID (0.00s)
    --- PASS: TestAuthenticateHandler/invalid_ID (0.00s)
=== RUN   TestHTTPSServer
--- PASS: TestHTTPSServer (0.05s)
PASS
coverage: 87.5% of statements
ok      github.com/maxlambrecht/svid-exercise/service/server    0.055s  coverage: 87.5% of statements
=== RUN   TestValidateID
=== RUN   TestValidateID/Recognized_Spiffe_ID
=== RUN   TestValidateID/Unrecognized_Spiffe_ID
--- PASS: TestValidateID (0.00s)
    --- PASS: TestValidateID/Recognized_Spiffe_ID (0.00s)
    --- PASS: TestValidateID/Unrecognized_Spiffe_ID (0.00s)
PASS
coverage: 85.7% of statements
ok      github.com/maxlambrecht/svid-exercise/service/validator        coverage: 85.7% of statements

```

### Appendix

##### Create certificates

###### Create a ROOT CA key
openssl genrsa -des3 -out rootCA.key 4096

###### Create and self sign the Root Certificate
openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 1024 -out rootCA.crt

##### Client Certificate

###### Generate key
openssl genrsa -out client.key 2048

###### Generate csr
openssl req -new -key client.key -out client.csr -subj "/C=US/ST=US/L=us/O=client/OU=client/CN=client" 


Edit SpiffeID in file _certs/conf.cnf_:

```
[alt_names]
URI.1  = spiffe://example.com/service

```

###### Generate certificate signed with rootCA
openssl x509 -req -in client.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out client.crt -days 500 -sha256 -extfile conf.cnf -extensions usr_cert

###### View
openssl x509 -in client.crt -text -noout 


##### Server Certificate

###### Generate key
openssl genrsa -out server.key 2048

###### Generate csr
openssl req -new -key server.key -out server.csr -subj "/C=US/ST=US/L=us/O=server/OU=server/CN=localhost" 

###### Generate certificate signed with rootCA
openssl x509 -req -in server.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out server.crt -days 500 -sha256

###### View
openssl x509 -in server.crt -text -noout 


