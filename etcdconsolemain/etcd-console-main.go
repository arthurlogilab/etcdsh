package etcdconsolemain

import (
	"flag"
	"fmt"
)

func Start() {
	var url = flag.String("url", "http://localhost:4001", "etcd url")

	flag.Parse()

	fmt.Println(*url)

}
