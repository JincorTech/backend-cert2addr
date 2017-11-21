# Jincor Backend Cert2Addr

Generate a ethereum like address by passed certificate encoded in pem format.

Only one endpoint is available:

1. `/api/certificates/actions/getaddress [POST]`
Default request:
```
{
  "pem": "-----BEGIN CERTIFICATE----\n.....\n-----END CERTIFICATE-----\n"
}
```
Successfull response:
```
{
  "status": 200,
  "address": "ed280b6c632e6fb7233ec29539ec0bb799d1ce24"
}
```
