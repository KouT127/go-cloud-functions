# 確認
```
gcloud config list
```

# プロジェクト確認
```bash
gcloud projects list
```

# プロジェクトセット
```bash
gcloud config set project PROJECT_NAME
```

# deploy

```bash
./DIR_NAME
go mod vendor
gcloud functions deploy FUNCTIONS_NAME --runtime go113 --trigger-http --allow-unauthenticatedgcloud functions deploy FUNCTIONS_NAME --runtime go113 --trigger-http --allow-unauthenticated
```