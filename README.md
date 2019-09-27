# apiEchoGAE

- secret.yamlファイルを作成
  - app.yaml同じ場所に配置

```yaml
env_variables:
  ID: 'xxx'
  PW: 'yyy'
```

- deploy
  - `gcloud auth login`
  - `gcloud app deploy --project ${GCP_PROJECT_ID}`

- api
  - `curl https://${GCP_PROJECT_ID}.appspot.com/metal`
