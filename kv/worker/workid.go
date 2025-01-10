package worker

import (
	"github.com/micro-services-roadmap/cloudflare/kv"
	"github.com/micro-services-roadmap/cloudflare/kv/cf"
	"github.com/micro-services-roadmap/cloudflare/util"
	"strconv"
)

var (
	KeyUidsWorker = "uids-worker"
	NamespaceKey  = "CLOUDFLARE_KV_NAMESPACE"
	Namespace     = "ENC(9aNc9rnyBiNExymn87ObqhkVI4EuojUgzYYOY8POsySMb2ESlC3XiIAqxNYZDUR3bEOEKGXkCBI+n8n5jDv8eLn1sbF3kipyqZXFjeZN1Ws=)"
	// NamespaceName in cloudflare: prod-uid-worker
)

func init() {
	// 1. parse kv storage key
	Namespace = util.GetEnvDefault(NamespaceKey, Namespace)
}

func NextWorkerID() (int64, error) {
	readResp, err := cf.GetWorkersKV(kv.AccountID, Namespace, KeyUidsWorker)
	if err != nil {
		return 0, err
	}

	currentID, err := strconv.Atoi(string(readResp))
	if err != nil {
		return 0, err
	}

	nextId := int64(currentID + 1)
	if _, err = cf.WriteWorkersKV(kv.AccountID, Namespace, KeyUidsWorker, strconv.FormatInt(nextId, 10)); err != nil {
		return 0, err
	}
	return nextId, nil
}
