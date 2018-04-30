#!/usr/bin/env bash

CERT=$1
openssl x509 -in ${CERT} -text -noout \
-certopt no_subject,no_header,no_version,no_serial,no_signame,no_validity,no_subject,no_issuer,no_pubkey,no_sigdump,no_aux  | awk '/X509v3 Subject Alternative Name/','/X509v3 Basic Constraints/'
