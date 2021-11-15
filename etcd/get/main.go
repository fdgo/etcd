package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"time"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		kv     clientv3.KV
		err    error
		getResp *clientv3.GetResponse
		kvpair  *mvccpb.KeyValue
	)
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	kv = clientv3.NewKV(client)
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/",clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(getResp.Kvs,getResp.Count)
		for _, kvpair = range getResp.Kvs {
			fmt.Println("get: ", string(kvpair.Key), string(kvpair.Value))
		}
	}
}
