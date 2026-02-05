# `github.com/lestrrat-go/jwx/v3` Manual Refresh Bug

Output:

```sh
$ go run ./main.go
{ "time": "2026-02-05T14:42:31.479631012+01:00", "level": "DEBUG", "msg": "received JWK set request" }
{ "time": "2026-02-05T14:42:31.479735048+01:00", "level": "WARN", "msg": "failed to refresh JWK cache automatically", "error": "httprc.Resource.Sync: unexpected status code (status code=500, url="http://127.0.0.1:41699")" }
{ "time": "2026-02-05T14:42:32.480176686+01:00", "level": "DEBUG", "msg": "received JWK set request" }
{ "time": "2026-02-05T14:42:32.480329082+01:00", "level": "WARN", "msg": "failed to refresh JWK cache automatically", "error": "httprc.Resource.Sync: unexpected status code (status code=500, url="http://127.0.0.1:41699")" }
{ "time": "2026-02-05T14:42:33.480751385+01:00", "level": "DEBUG", "msg": "received JWK set request" }
{ "time": "2026-02-05T14:42:33.480988109+01:00", "level": "WARN", "msg": "failed to refresh JWK cache automatically", "error": "httprc.Resource.Sync: unexpected status code (status code=500, url="http://127.0.0.1:41699")" }
{ "time": "2026-02-05T14:42:34.481411555+01:00", "level": "DEBUG", "msg": "received JWK set request" }
{ "time": "2026-02-05T14:42:34.481601282+01:00", "level": "WARN", "msg": "failed to refresh JWK cache automatically", "error": "httprc.Resource.Sync: unexpected status code (status code=500, url="http://127.0.0.1:41699")" }
{ "time": "2026-02-05T14:42:35.482246455+01:00", "level": "DEBUG", "msg": "received JWK set request" }
{ "time": "2026-02-05T14:42:35.482448084+01:00", "level": "WARN", "msg": "failed to refresh JWK cache automatically", "error": "httprc.Resource.Sync: unexpected status code (status code=500, url="http://127.0.0.1:41699")" }
{ "time": "2026-02-05T14:42:36.480431841+01:00", "level": "DEBUG", "msg": "received JWK set request" }
{ "time": "2026-02-05T14:42:36.480523613+01:00", "level": "ERROR", "msg": "failed to refresh JWK cache manually" }
{ "time": "2026-02-05T14:42:38.483941131+01:00", "level": "DEBUG", "msg": "received JWK set request" }
{ "time": "2026-02-05T14:42:38.484107603+01:00", "level": "WARN", "msg": "failed to refresh JWK cache automatically", "error": "httprc.Resource.Sync: unexpected status code (status code=500, url="http://127.0.0.1:41699")" }
{ "time": "2026-02-05T14:42:39.48531158+01:00", "level": "DEBUG", "msg": "received JWK set request" }
{ "time": "2026-02-05T14:42:39.485406058+01:00", "level": "WARN", "msg": "failed to refresh JWK cache automatically", "error": "httprc.Resource.Sync: unexpected status code (status code=500, url="http://127.0.0.1:41699")" }
{ "time": "2026-02-05T14:42:40.486493427+01:00", "level": "DEBUG", "msg": "received JWK set request" }
{ "time": "2026-02-05T14:42:40.486561415+01:00", "level": "WARN", "msg": "failed to refresh JWK cache automatically", "error": "httprc.Resource.Sync: unexpected status code (status code=500, url="http://127.0.0.1:41699")" }
{ "time": "2026-02-05T14:42:41.481145832+01:00", "level": "DEBUG", "msg": "received JWK set request" }
{ "time": "2026-02-05T14:42:41.481278391+01:00", "level": "ERROR", "msg": "failed to refresh JWK cache manually" }
{ "time": "2026-02-05T14:42:43.48785339+01:00", "level": "DEBUG", "msg": "received JWK set request" }
{ "time": "2026-02-05T14:42:43.488120581+01:00", "level": "WARN", "msg": "failed to refresh JWK cache automatically", "error": "httprc.Resource.Sync: unexpected status code (status code=500, url="http://127.0.0.1:41699")" }
{ "time": "2026-02-05T14:42:44.489282154+01:00", "level": "DEBUG", "msg": "received JWK set request" }
{ "time": "2026-02-05T14:42:44.489475687+01:00", "level": "WARN", "msg": "failed to refresh JWK cache automatically", "error": "httprc.Resource.Sync: unexpected status code (status code=500, url="http://127.0.0.1:41699")" }
{ "time": "2026-02-05T14:42:45.489842879+01:00", "level": "DEBUG", "msg": "received JWK set request" }
{ "time": "2026-02-05T14:42:45.489985898+01:00", "level": "WARN", "msg": "failed to refresh JWK cache automatically", "error": "httprc.Resource.Sync: unexpected status code (status code=500, url="http://127.0.0.1:41699")" }
{ "time": "2026-02-05T14:42:46.482261434+01:00", "level": "DEBUG", "msg": "received JWK set request" }
{ "time": "2026-02-05T14:42:46.482406767+01:00", "level": "ERROR", "msg": "failed to refresh JWK cache manually" }

# Should continue to show DEBUG and WARN logs, but it's completely stuck.
```
