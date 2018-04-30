## SVID basic validation

##### This example contains: 
 - A HTTP-based SPIFFE Verifiable Identity Document (SVID) validation service that takes a SPIFFE ID as a configuration parameter. 
 - A client that connects over mTLS using an SVID.
     
###### The service returns:

- HTTP 200 if the URI SAN of the SVID matches the configured SPIFFE ID.
- HTTP 401 if the URI SAN does not match the configured SPIFFE ID.

##### How to run the example

###### Get the code and dependencies

```
go get github.com/maxlambrecht/svid-exercise
cd ~/go/src/github.com/maxlambrecht/svid-exercise
go get ./... 
```


###### Run the Server
```
go run service/main.go --spiffeid spiffe://example.com/service --cert certs/server_cert.pem --key certs/server_key.pem
Server listening on address :3000
```

By default the server listens on _https://localhost:3000_
To listen on another address use the option _--addr_
For example

```
go run service/main.go --spiffeid spiffe://example.com/service --cert certs/server_cert.pem --key certs/server_key.pem --addr localhost:4000
Server listening on address localhost:4000
```

###### Run the Client

```
# using a certificate that has a Subject Alternative Name=spiffe://example.com/service:
go run client/main.go --cert certs/client_cert.pem --key certs/client_key.pem 
Response code: 200
Authentication Succeed
```

```
# using a certificate with an untrusted Subject Alternative Name or not SAN at all
go run client/main.go --cert certs/unknown_client_cert.pem --key certs/unknown_client_key.pem
Response code: 401
Authentication Failed. Invalid SpiffeID
```

By default the client sends the requests to _https://localhost:3000/auth_
To send the requests to another address use the option _--url_ 
For example

```
go run client/main.go --cert certs/client_cert.pem --key certs/client_key.pem --url https://localhost:4000/auth
Response code: 200
Authentication Succeed
```


###### Run Tests


```
go test ./... -v -cover

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

#### Appendix

##### Create certificate

Edit SpiffeID in file _certs/conf.cnf_:

```
[alt_names]
URI.1  = spiffe://example.com/service

```

Run _generate-cert.sh_ in directory _certs_:

```
sh generate-cert.sh cert_file.pem key_file.pem
```

Run _view-cert-san.sh_ to verify the SpiffeID in the Certificate:


```
sh view-cert-san.sh cert1.pem

X509v3 Subject Alternative Name: 
                URI:spiffe://example.com/service

```

