package session

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type HttpDoer interface {
	Do(*http.Request) (*http.Response, error)
}

type Starter interface {
	SetTimeout(time.Duration) Starter
	Start(string, string) string
}

type client struct {
	merchId    string
	domainName string
	cert       tls.Certificate
	caPool     *x509.CertPool
	http       HttpDoer
	timeout    time.Duration
}

func New(merchId, domainName string, cert tls.Certificate, caPool *x509.CertPool) Starter {
	return &client{
		merchId:    merchId,
		domainName: domainName,
		cert:       cert,
		caPool:     caPool,
		timeout:    10 * time.Second,
	}
}

func (a *client) SetTimeout(t time.Duration) Starter {
	a.timeout = t
	return a
}

func (a *client) SetHttp(client HttpDoer) *client {
	a.http = client
	return a
}

func (a *client) GetHttp() HttpDoer {
	if a.http != nil {
		return a.http
	}

	a.http = a.buildHttpClient()

	return a.http
}

func checkRedirectPolicy(req *http.Request, via []*http.Request) error {
	// Never allow redirects
	return http.ErrUseLastResponse
}

func (a *client) buildHttpClient() *http.Client {

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{a.cert},
		RootCAs:            a.caPool,
		InsecureSkipVerify: false,
		MinVersion:         tls.VersionTLS12,
	}

	tlsConfig.BuildNameToCertificate()

	transport := &http.Transport{TLSClientConfig: tlsConfig}

	return &http.Client{
		Transport:     transport,
		Timeout:       a.timeout,
		CheckRedirect: checkRedirectPolicy,
	}
}

func (a *client) Start(url, displayName string) string {
	// Create request
	body := fmt.Sprintf("{\"merchantIdentifier\":\"%s\",\"domainName\":\"%s\",\"displayName\":\"%s\"}", a.merchId, a.domainName, displayName)

	request, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/json")

	// Make the request
	resp, err := a.GetHttp().Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Start session failed; %d %s", resp.StatusCode, resp.Status)
	}

	// read response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
