package etcd

import (
	"context"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
	"strings"
	"time"

	"go.etcd.io/etcd/client"
	"go.etcd.io/etcd/pkg/transport"
)

var ETCD_URL  = "http://127.0.0.1:2379"
var kapi client.KeysAPI
var etcdClient *client.Client
var RunMode = "dev"
func init() {
	var err error
	if RunMode == "prod" {
		etcdClient, err = NewEtcdClientTLS()
	} else {
		etcdClient, err = NewEtcdClient()
	}
	if err != nil {
		log.Fatal(err)
	}
	kapi = client.NewKeysAPI(*etcdClient)
}

func NewEtcdClient() (*client.Client, error) {
	cfg := client.Config{
		Endpoints: strings.Split(ETCD_URL, ","),
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func NewEtcdClientTLS() (*client.Client, error) {
	tls := transport.TLSInfo{
		CertFile:      "keys/etcd/ssl/etcd.pem",
		KeyFile:       "keys/etcd/ssl/etcd-key.pem",
		TrustedCAFile: "keys/etcd/ssl/ca.pem",
	}
	tp, err := transport.NewTransport(tls, 30*time.Second)
	if err != nil {
		return nil, err
	}
	cfg := client.Config{
		Endpoints: strings.Split("", ","),
		Transport: tp,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	return &c, nil
}



func EtcdSet(key, value string, opts *client.SetOptions) (resp *client.Response, err error) {
	return kapi.Set(context.Background(), key, value, opts)
}

func EtcdGet(key string) (resp *client.Response, err error) {
	resp, err = kapi.Get(context.Background(), key, &client.GetOptions{
		Recursive: true,
	})
	return
}

func Get(key string) (value string, err error) {
	resp, err := kapi.Get(context.Background(), key, &client.GetOptions{
		Recursive: true,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	value = resp.Node.Value
	return
}

func EtcdUnmarshal(path string, config interface{}) (err error) {
	resp, err := kapi.Get(context.Background(), path, &client.GetOptions{
		Recursive: true,
		Quorum:    true,
	})
	if err != nil {
		return
	}
	m := make(map[string]interface{})

	for _, v := range resp.Node.Nodes {

		sk := strings.Split(v.Key, "/")
		m[sk[len(sk)-1]] = v.Value

		fmt.Println(sk[len(sk)-1])
	}
	mapstructure.Decode(m, &config)
	return
}
