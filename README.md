aws --endpoint-url=http://localhost:4566 dynamodb create-table --cli-input-json file://src/infra/databases/dynamo/products_table_schema.json