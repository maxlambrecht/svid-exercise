## SVID basic validation

##### This example contains: 
 - A HTTP-based SPIFFE Verifiable Identity Document (SVID) validation service that takes a SPIFFE ID as a configuration parameter. 
 - A client that connects over mTLS using an SVID that is validated.
     
###### The service returns:

- HTTP 200 if the URI SAN of the SVID matches the configured SPIFFE ID.
- HTTP 401 if the URI SAN does not match the configured SPIFFE ID.

##### How to run the example

###### Run the Server
```
go run service/main.go --spiffeid spiffe://example.com/service --cert certs/server_cert.pem --key certs/server_key.pem
```

###### Run the Client

```
go run client/main.go --cert certs/client_cert.pem --key certs/client_key.pem 
Response code: 200
Authentication Succeed
```

```
go run client/main.go --cert certs/unknown_client_cert.pem --key certs/unknown_client_key.pem
Response code: 401
Authentication Failed. Invalid SpiffeID
```




