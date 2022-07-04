// Copyright 2021 The Cloudbases Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package etcd

import (
	i "cloudbases.io/libqueue"
	clientv3 "go.etcd.io/etcd/client/v3"
	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"
)

type EtcdClient struct {
	clientv3.Client
}

type EtcdQueue struct {
	EtcdClient
	Topic string
}

func (eq *EtcdQueue) SetTopic(topic string) i.IQueue {
	eq.Topic = topic
	return eq
}

func (eq *EtcdQueue) Enqueue(val string) error {
	etcdQueue := recipe.NewQueue(&eq.EtcdClient.Client, eq.Topic)
	return etcdQueue.Enqueue(val)
}

func (eq *EtcdQueue) Dequeue() (string, error) {
	etcdQueue := recipe.NewQueue(&eq.EtcdClient.Client, eq.Topic)
	return etcdQueue.Dequeue()
}

func (eq *EtcdQueue) SetClient(iClient *i.IClient) i.IQueue {
	eq.EtcdClient = (*iClient).(EtcdClient)
	return eq
}
