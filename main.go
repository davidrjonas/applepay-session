package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"github.com/davidrjonas/applepay-session/session"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var version = "1.0.0"

var (
	merchId     = kingpin.Flag("merchid", "Merchant id, e.g. merchant.com.example").Required().String()
	domainName  = kingpin.Flag("domain", "Domain name").Required().String()
	displayName = kingpin.Flag("displayname", "Display name").Required().String()

	certFile = kingpin.Flag("cert", "Client certificate file").Required().String()
	keyFile  = kingpin.Flag("key", "Client certificate key file").Required().String()
	caFile   = kingpin.Flag("ca", "Root CA file for validation").Required().String()

	url = kingpin.Arg("url", "New session url").Required().String()
)

func mustLoadCert(certFile, keyFile string) tls.Certificate {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	return cert
}

func mustLoadCa(caFile string) *x509.CertPool {
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatalln("Failed to create cert pool from caFile")
	}

	return caCertPool
}

func main() {
	kingpin.Version(version)
	kingpin.Parse()

	s := session.New(*merchId, *domainName, mustLoadCert(*certFile, *keyFile), mustLoadCa(*caFile))

	println(s.Start(*url, *displayName))
}
