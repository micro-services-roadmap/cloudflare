# cloudflare

```shell
go get github.com/micro-services-roadmap/cloudflare@v0.0.2
```

## cloudflare-kv

1. set follow env

   | env key                 |     desc      | sample |
   | :---------------------- | :-----------: | :----: |
   | CLOUDFLARE_API_KEY_KV   | kv api token  |   xx   |
   | CLOUDFLARE_ACCOUNT_ID   | cf account id |   xx   |
   | CLOUDFLARE_KV_NAMESPACE | kv namespace  |   xx   |

2. obtain worker-id for uid-generator by kv.worker

   - `KeyUidsWorker = "uids-worker"`

   ```go
   // 自增
   id, err := NextWorkerID()
   ```
