# Product Catalog Service

## Things TODO
- rename src folder by internal [X]
- Add loggin system []
- settings on viper or .env []
- move prompt files to cmd like controller on web api []
- swagger documentation []
- Improve tests []
- CI []
- Terraform for dynamodb and maybe EC2 instance []
- deploy []

## Features TODO
- Add list products use case on controller/prompt []
- validate product before insert []
- send email when insert products []
- import products by file (XLSX, CSV, TXT) []

```
aws --endpoint-url=http://localhost:4566 dynamodb create-table --cli-input-json file://src/infra/databases/dynamo/products_table_schema.json
```

```
go test -v -coverprofile cover.out ./YOUR_CODE_FOLDER/...
go tool cover -html=cover.out -o cover.html
open cover.html
```