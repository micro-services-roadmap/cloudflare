package worker

import (
	"github.com/micro-services-roadmap/cloudflare/kv"
	"github.com/micro-services-roadmap/cloudflare/kv/cf"
	"github.com/micro-services-roadmap/cloudflare/util"
	"github.com/spf13/cast"
	"strconv"
)

var (
	KeyUidsWorker = "uids-worker"
	NamespaceKey  = "CLOUDFLARE_KV_NAMESPACE"
	Namespace     = "ENC(4jS+nQOD3uyEklfxGq/bBTZHv6ehn4PNDhOYExOIvwQr+Tju+pat+hQGRWQJfTh1/vvrRnjORr3D0GZxz4dlZpjlSzBZEvoZFZtVbNgtPZM=)"
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
	if _, err = cf.WriteWorkersKV(kv.AccountID, Namespace, KeyUidsWorker, cast.ToString(nextId)); err != nil {
		return 0, err
	}

	return nextId, nil
}
