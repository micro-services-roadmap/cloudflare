package cf

import (
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/micro-services-roadmap/cloudflare/kv"
)

func GetWorkersKV(account, namespace, key string) ([]byte, error) {
	if kv.API == nil {
		return nil, fmt.Errorf("KV API not initialized")
	}

	resp, err := kv.API.GetWorkersKV(context.Background(),
		cloudflare.AccountIdentifier(account),
		cloudflare.GetWorkersKVParams{NamespaceID: namespace, Key: key})
	fmt.Printf("GetWorkersKV[key: %s] from namespace[%s] Response: %+v\n", key, namespace, resp)
	if err != nil {
		fmt.Printf("GetWorkersKV[key: %s] from namespace[%s] Error: %s\n", key, namespace, err)
	}

	return resp, nil
}

func WriteWorkersKV(account, namespace, key, value string) (*cloudflare.Response, error) {
	if kv.API == nil {
		return nil, fmt.Errorf("KV API not initialized")
	}

	resp, err := kv.API.WriteWorkersKVEntry(context.Background(),
		cloudflare.AccountIdentifier(account),
		cloudflare.WriteWorkersKVEntryParams{
			NamespaceID: namespace,
			Key:         key,
			Value:       []byte(value),
		})
	fmt.Printf("WriteWorkersKV[key: %s] with value[%s] to namespace[%s] Response: %+v\n", key, value, namespace, resp)
	if err != nil {
		fmt.Printf("WriteWorkersKV[key: %s] with value[%s] to namespace[%s] Error: %s\n", key, value, namespace, err)
	}

	return &resp, nil
}
