package dynamo

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/pedropazello/lojinha-product-catalog-service/src/domain/entities"
)

const tableName = "Products"

type ProductRepository struct {
}

func (p *ProductRepository) Create(product *entities.Product) (*entities.Product, error) {
	product.ID = uuid.New().String()

	svc := NewClient()

	av, err := dynamodbattribute.MarshalMap(product)
	if err != nil {
		return product, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	if _, err = svc.PutItem(input); err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductRepository) GetById(id string) (*entities.Product, error) {
	svc := NewClient()

	product := &entities.Product{}

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return product, err
	}

	if result.Item == nil {
		return nil, errors.New("Could not find '" + id + "'")
	}

	if err = dynamodbattribute.UnmarshalMap(result.Item, &product); err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductRepository) Save(product *entities.Product) error {
	svc := NewClient()

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":n": {
				S: aws.String(product.Name),
			},
			":d": {
				S: aws.String(product.Description),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#name": aws.String("Name"),
			"#desc": aws.String("Description"),
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(product.ID),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("SET #name = :n, #desc = :d"),
	}

	if _, err := svc.UpdateItem(input); err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) Delete(id string) error {
	svc := NewClient()

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(tableName),
	}

	if _, err := svc.DeleteItem(input); err != nil {
		return err
	}

	return nil
}
