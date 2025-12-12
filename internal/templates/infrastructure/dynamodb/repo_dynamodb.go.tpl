package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"{{ .Module }}/internal/domain/entities"
	{{- if .Features.ErrorHandling }}
	"{{ .Module }}/internal/domain/errors"
	{{- end }}
	"{{ .Module }}/internal/repository"
)

type {{ .Entity }}DynamoDBRepo struct {
	Client    *dynamodb.Client
	TableName string
}

func New{{ .Entity }}DynamoDBRepo(client *dynamodb.Client, tableName string) repository.{{ .Entity }}Repository {
	return &{{ .Entity }}DynamoDBRepo{
		Client:    client,
		TableName: tableName,
	}
}

func (r *{{ .Entity }}DynamoDBRepo) Save(ctx context.Context, e *entities.{{ .Entity }}) error {
	item, err := attributevalue.MarshalMap(e)
	if err != nil {
		{{- if .Features.ErrorHandling }}
		return errors.ErrInternalServer
		{{- else }}
		return err
		{{- end }}
	}

	_, err = r.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.TableName),
		Item:      item,
	})

	{{- if .Features.ErrorHandling }}
	if err != nil {
		return errors.ErrInternalServer
	}
	{{- end }}

	return err
}

func (r *{{ .Entity }}DynamoDBRepo) FindByID(ctx context.Context, id string) (*entities.{{ .Entity }}, error) {
	result, err := r.Client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(r.TableName),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		{{- if .Features.ErrorHandling }}
		return nil, errors.ErrInternalServer
		{{- else }}
		return nil, err
		{{- end }}
	}

	if result.Item == nil {
		{{- if .Features.ErrorHandling }}
		return nil, errors.NewNotFoundError("{{ .Entity }}")
		{{- else }}
		return nil, nil
		{{- end }}
	}

	var entity entities.{{ .Entity }}
	if err := attributevalue.UnmarshalMap(result.Item, &entity); err != nil {
		{{- if .Features.ErrorHandling }}
		return nil, errors.ErrInternalServer
		{{- else }}
		return nil, err
		{{- end }}
	}

	return &entity, nil
}
