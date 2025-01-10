package cf

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/kv"
	"github.com/micro-services-roadmap/cloudflare/kvs"
	"io"
)

type NamespaceValueGetResponse struct {
	// Arbitrary JSON to be associated with a key/value pair.
	Metadata string `json:"metadata"`
	// A byte sequence to be stored, up to 25 MiB in length.
	Value any `json:"value"`
}

func GetWorkersKV(account, namespace, key string) (*NamespaceValueGetResponse, error) {
	if kvs.Client == nil {
		return nil, fmt.Errorf("KV Client not initialized")
	}

	resp, err := kvs.Client.KV.Namespaces.Values.Get(context.Background(), namespace, key,
		kv.NamespaceValueGetParams{AccountID: cloudflare.F(account)})
	if err != nil {
		fmt.Printf("GetWorkersKV[key: %s] from namespace[%s] Error: %s\n", key, namespace, err)
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			fmt.Println(string(apierr.DumpRequest(true)))
		}
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	fmt.Printf("GetWorkersKV[key: %s] from namespace[%s] Response: %+v\n", key, namespace, string(bytes))
	if err != nil {
		return nil, err
	}

	kvResponse := &NamespaceValueGetResponse{}
	if err = json.Unmarshal(bytes, kvResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal GetWorkersKV response: %w", err)
	}

	return kvResponse, nil
}

func WriteWorkersKV(account, namespace, key, value string) (*kv.NamespaceValueUpdateResponse, error) {
	if kvs.Client == nil {
		return nil, fmt.Errorf("KV Client not initialized")
	}

	// Expiration is in seconds since the epoch
	resp, err := kvs.Client.KV.Namespaces.Values.Update(context.Background(), namespace, key,
		kv.NamespaceValueUpdateParams{
			AccountID: cloudflare.F(account),
			Value:     cloudflare.F(value),
			//Metadata:      cloudflare.F("{\"someMetadataKey\": \"someMetadataValue\"}"),
			//Expiration:    cloudflare.F(1578435000.000000),
			//ExpirationTTL: cloudflare.F(300.000000),
		})
	if err != nil {
		fmt.Printf("WriteWorkersKV[key: %s] with value[%s] to namespace[%s] Error: %s\n", key, value, namespace, err)
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			fmt.Println(string(apierr.DumpRequest(true)))
		}
		return nil, err
	}

	fmt.Printf("WriteWorkersKV[key: %s] with value[%s] to namespace[%s] Response: %+v\n", key, value, namespace, resp.JSON.RawJSON())
	return resp, nil
}
