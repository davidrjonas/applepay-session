ApplePay Session
================

Starts a new session with Apple Pay and outputs the response data.

Motivation
----------

On Ubuntu 12.04, even with OpenSSL that supports TLS 1.2 (a requirement for Apple Pay), cURL will not use it. The upshot is that PHP with cURL will not use TLS 1.2 either. And PHP as well, even in the newest versions since they are compiled against the same libcurl. This helper program can handle the server side session initiation and Etsy's [applepay-php](https://github.com/etsy/applepay-php) can handle the rest.

