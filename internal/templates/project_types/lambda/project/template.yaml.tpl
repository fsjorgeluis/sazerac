AWSTemplateFormatVersion: '2010-09-09'
Transform: AWS::Serverless-2016-10-31
Description: {{ .ProjectName }} - Serverless application

Globals:
  Function:
    Timeout: 30
    Runtime: provided.al2
    Architectures:
      - x86_64
    Environment:
      Variables:
        {{- if eq .Features.Database "dynamodb" }}
        TABLE_NAME: !Ref {{ .Entity }}Table
        {{- end }}

Resources:
  {{ .UseCase }}Function:
    Type: AWS::Serverless::Function
    Properties:
      CodeUri: .
      Handler: bootstrap
      {{- if .Features.APIGateway }}
      Events:
        ApiEvent:
          Type: Api
          Properties:
            Path: /{{ .UseCaseRoute }}
            Method: POST
      {{- end }}
      {{- if eq .Features.Database "dynamodb" }}
      Policies:
        - DynamoDBCrudPolicy:
            TableName: !Ref {{ .Entity }}Table
      {{- end }}

  {{- if eq .Features.Database "dynamodb" }}
  {{ .Entity }}Table:
    Type: AWS::DynamoDB::Table
    Properties:
      TableName: {{ .ProjectName }}-{{ .Entity | ToLower }}-table
      AttributeDefinitions:
        - AttributeName: ID
          AttributeType: S
      KeySchema:
        - AttributeName: ID
          KeyType: HASH
      BillingMode: PAY_PER_REQUEST
  {{- end }}

Outputs:
  {{- if .Features.APIGateway }}
  ApiEndpoint:
    Description: "API Gateway endpoint URL"
    Value: !Sub "https://${ServerlessRestApi}.execute-api.${AWS::Region}.amazonaws.com/Prod/{{ .UseCaseRoute }}/"
  {{- end }}
  {{ .UseCase }}FunctionArn:
    Description: "Lambda Function ARN"
    Value: !GetAtt {{ .UseCase }}Function.Arn
  {{- if eq .Features.Database "dynamodb" }}
  {{ .Entity }}TableName:
    Description: "DynamoDB Table Name"
    Value: !Ref {{ .Entity }}Table
  {{- end }}
