#!/usr/bin/env bash

CERT=$1
openssl x509 -in ${CERT} -text
