package main

import (
	"etcd/etcd"
	"fmt"
)

type CJPayConfig struct {
	PartnerId    string
	PayPublicKey string
	PublicKey    string
	PrivateKey   string
	CallbackURL  string
	CjSingleurl  string
}



func main() {
	//resp, err := etcd.EtcdGet("/jf.com/saas/jf/ZmConfig/Url")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//sort.Sort(resp.Node.Nodes)

	//fmt.Println(resp.Node.Key)
	//fmt.Println(resp.Node.Value)

	//for _, v := range resp.Node.Nodes {
	//	fmt.Println(v.Key)
	//	fmt.Println(v.Value)
	//}

	var zm CJPayConfig


	err := etcd.EtcdUnmarshal("/jf.com/saas/jf/CJPayConfig", &zm)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(zm)

}
