package echo

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	echopb "github.com/bradenbass/echo/proto"
)

const (
	clientCert           = "tls/Client.crt"
	clientKey            = "tls/Client.key"
	certificateAuthority = "tls/Echo_CA.crt"
)

func NewClient(address string, secure bool) (echopb.EchoerClient, error) {
	var opts []grpc.DialOption

	if secure {
		certificate, err := tls.LoadX509KeyPair(clientCert, clientKey)
		if err != nil {
			return nil, fmt.Errorf("could not load client key pair: %s", err)
		}

		// Create a certificate pool from the certificate authority
		certPool := x509.NewCertPool()
		ca, err := ioutil.ReadFile(certificateAuthority)
		if err != nil {
			return nil, fmt.Errorf("failed to read ca: %v", err)
		}

		// Append the certificates from the CA
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			return nil, fmt.Errorf("unable to append certs from certificate authority %v", ca)
		}

		tlsCreds := credentials.NewTLS(&tls.Config{
			Certificates: []tls.Certificate{certificate},
			RootCAs:      certPool,
		})
		opts = append(opts, grpc.WithTransportCredentials(tlsCreds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// Dial a new connection
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, err
	}
	return echopb.NewEchoerClient(conn), nil
}
