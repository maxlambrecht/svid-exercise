## SVID basic validation

##### This example contains: 
 - A HTTP-based SPIFFE Verifiable Identity Document (SVID) validation service that takes a SPIFFE ID as a configuration parameter. 
 - A client that connects over mTLS using an SVID.
     
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


#### Appendix

##### Create certificate

Edit SpiffeID in file certs/conf.cnf:

```
[alt_names]
URI.1  = spiffe://example.com/service

```

Run generate-cert.sh in directory _certs_:

```
sh generate-cert.sh cert_file.pem key_file.pem
```

Run view-cert-san.sh to verify the SpiffeID in the Certificate:


```
sh view-cert-san.sh cert1.pem

X509v3 Subject Alternative Name: 
                URI:spiffe://example.com/service

```

