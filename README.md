ApplePay Session
================

Starts a new session with Apple Pay and outputs the response data.

Usage
-----

```
usage: applepay-session --merchid=MERCHID --domain=DOMAIN --displayname=DISPLAYNAME --cert=CERT --key=KEY --ca=CA [<flags>] <url>

Flags:
  --help                     Show context-sensitive help (also try --help-long
                             and --help-man).
  --merchid=MERCHID          Merchant id, e.g. merchant.com.example
  --domain=DOMAIN            Domain name
  --displayname=DISPLAYNAME  Display name
  --cert=CERT                Client certificate file
  --key=KEY                  Client certificate key file
  --ca=CA                    Root CA file for validation
  --version                  Show application version.

Args:
  <url>  New session url
```

Library Usage
-------------

See main.go for an example, but the interface is simple:

```
s := session.New(merchId, domainName, cert, caCertPool)

session_data := s.Start(url, displayName)
```

Motivation
----------

On Ubuntu 12.04, even with OpenSSL that supports TLS 1.2 (a requirement for Apple Pay), cURL will not use it. The upshot is that PHP with cURL will not use TLS 1.2 either. And PHP as well, even in the newest versions since they are compiled against the same libcurl. This helper program can handle the server side session initiation and Etsy's [applepay-php](https://github.com/etsy/applepay-php) can handle the rest.

