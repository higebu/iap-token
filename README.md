# iap-token

Generate Google Cloud Identity-Aware Proxy token with [Impersonate Service Account](https://cloud.google.com/iam/docs/impersonating-service-accounts)

# Install

```shell
go install github.com/higebu/iap-token@latest
```

# Usage

```shell
export CLOUDSDK_AUTH_IMPERSONATE_SERVICE_ACCOUNT=user-sa-name@project-id.iam.gserviceaccount.com
export IAP_CLIENT_ID=xxxxx.apps.googleusercontent.com
curl -H "Authorization: Bearer $(iap-token)" https://example.com
# curl -H "Proxy-Authorization: Bearer $(iap-token)" https://example.com
```
