package etcdv3

import (
	"context"
	"errors"
	"path"
	"strings"
	"time"

	mvccpb "github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
)

// List 获取目录下列表
func (c *Etcd3Client) List(key string) (nodes []*Node, err error) {
	if key == "" {
		return make([]*Node, 0), errors.New("key is empty")
	}
	// 兼容key前缀设置为 /
	dir := key
	if key != "/" {
		key = strings.TrimRight(key, "/")
		dir = key + "/"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	txn := c.Client.Txn(ctx)
	txn.If(
		clientv3.Compare(
			clientv3.Value(key),
			"=",
			DEFAULT_DIR_VALUE,
		),
	).Then(
		clientv3.OpGet(dir, clientv3.WithPrefix()),
	)

	txnResp, err := txn.Commit()
	if err != nil {
		return nil, err
	}

	if !txnResp.Succeeded {
		return nil, ErrorListKey
	} else {
		if len(txnResp.Responses) > 0 {
			rangeResp := txnResp.Responses[0].GetResponseRange()
			return c.list(dir, rangeResp.Kvs)
		} else {
			// empty directory
			return []*Node{}, nil
		}
	}

	return []*Node{}, nil
}

// Value 获取一个key的值
func (c *Etcd3Client) Value(key string) (val *Node, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.Client.Get(ctx, key)
	if err != nil {
		return
	}
	if resp.Kvs != nil && len(resp.Kvs) > 0 {
		val = &Node{
			Value:   string(resp.Kvs[0].Value),
			FullDir: key,
			Version: resp.Kvs[0].Version,
		}
	} else {
		err = ErrorKeyNotFound
	}
	return
}

func (c *Etcd3Client) list(dir string, kvs []*mvccpb.KeyValue) ([]*Node, error) {
	nodes := []*Node{}
	for _, kv := range kvs {
		name := strings.TrimPrefix(string(kv.Key), dir)
		if strings.Contains(name, "/") {
			// secondary directory
			continue
		}
		nodes = append(nodes, NewNode(dir, kv))
	}
	return nodes, nil
}

func (c *Etcd3Client) ensureKey(key string) (string, string) {
	key = strings.TrimRight(key, "/")
	if key == "" {
		return "/", ""
	}
	if strings.Contains(key, "/") == true {
		return key, path.Clean(key + "/../")
	} else {
		return key, ""
	}
}

// Put 添加一个key
func (c *Etcd3Client) Put(key string, value string, mustEmpty bool) error {
	// log.Println(key)
	key, parentKey := c.ensureKey(key)
	//  需要判断的条件
	cmp := make([]clientv3.Cmp, 0)

	if parentKey != "" {
		cmp = append(cmp, clientv3.Compare(
			clientv3.Value(parentKey),
			"=",
			DEFAULT_DIR_VALUE,
		))
	}

	if mustEmpty {
		cmp = append(
			cmp,
			clientv3.Compare(
				clientv3.Version(key),
				"=",
				0,
			),
		)
	} else {
		cmp = append(
			cmp,
			clientv3.Compare(
				clientv3.Value(key),
				"!=",
				DEFAULT_DIR_VALUE,
			),
		)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	txn := c.Client.Txn(ctx)
	// make sure the parentKey is a directory
	txn.If(
		cmp...,
	).Then(
		clientv3.OpPut(key, value),
	)

	txnResp, err := txn.Commit()
	if err != nil {
		return err
	}

	if !txnResp.Succeeded {
		return ErrorPutKey
	}
	return nil
}

// Delete 删除key
func (c *Etcd3Client) Delete(key string) error {
	key = strings.TrimRight(key, "/")
	dir := key + "/"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	txn := c.Client.Txn(ctx)
	// 如果是目录，删除整个目录
	txn.If(
		clientv3.Compare(
			clientv3.Value(key),
			"=",
			DEFAULT_DIR_VALUE,
		),
	).Then(
		clientv3.OpDelete(key),
		clientv3.OpDelete(dir, clientv3.WithPrefix()),
	).Else(
		clientv3.OpDelete(key),
	)

	_, err := txn.Commit()
	return err
}

// GetRecursiveValue 获取前缀下的所有key
func (c *Etcd3Client) GetRecursiveValue(key string) (list []*Node, err error) {
	list = make([]*Node, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp, err := c.Client.KV.Get(ctx, key, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		return nil, err
	}
	for _, kv := range resp.Kvs {
		list = append(list, &Node{
			Value:   string(kv.Value),
			FullDir: string(kv.Key),
			Version: kv.Version,
		})
	}

	return
}
