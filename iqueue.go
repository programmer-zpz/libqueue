// Copyright 2021 The Cloudbases Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package libqueue

type IQueue interface {
	SetClient(iClient *IClient) IQueue
	SetTopic(topic string) IQueue
	Enqueue(string) error
	Dequeue() (string, error)
}
