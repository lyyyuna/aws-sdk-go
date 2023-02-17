package s3

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/restxml"
)

const opPutBucketStyleConfig = "PutBucketStyleConfig"

// 设置分隔符和原图保护
func (c *S3) PutBucketStyleConfig(input *PutBucketStyleConfigInput) (*PutBucketStyleConfigOutput, error) {
	req, out := c.PutBucketStyleConfigRequest(input)
	return out, req.Send()
}

func (c *S3) PutBucketStyleConfigRequest(input *PutBucketStyleConfigInput) (req *request.Request, output *PutBucketStyleConfigOutput) {
	op := &request.Operation{
		Name:       opPutBucketStyleConfig,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?styleConfig",
	}

	if input == nil {
		input = &PutBucketStyleConfigInput{}
	}

	output = &PutBucketStyleConfigOutput{}
	req = c.NewRequest(op, input, output)
	req.Handlers.Build.PushBack(contentMD5)

	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return
}

type PutBucketStyleConfigInput struct {
	_ struct{} `type:"structure"`

	Bucket *string `location:"uri" locationName:"Bucket" type:"string"`

	StyleInfoConfiguration *StyleInfoConfiguration `locationName:"StyleInfoConfiguration" type:"structure"`
}

func (s *PutBucketStyleConfigInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type PutBucketStyleConfigOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `location:"header" locationName:"x-amz-request-id" type:"string"`
}

type StyleInfoConfiguration struct {
	_ struct{} `type:"structure"`

	Separator *string `locationName:"Separator" type:"string"`

	OriginalResourceProtect *OriginalResourceProtect `locationName:"OriginalResourceProtect" type:"structure"`
}

type OriginalResourceProtect struct {
	_ struct{} `type:"structure"`

	Enable *bool `locationName:"Enable" type:"bool"`
}

const opGetBucketStyleConfig = "GetBucketStyleConfig"

// 查询分隔符和原图保护
func (c *S3) GetBucketStyleConfig(input *GetBucketStyleConfigInput) (*GetBucketStyleConfigOutput, error) {
	req, out := c.GetBucketStyleConfigRequest(input)
	return out, req.Send()
}

func (c *S3) GetBucketStyleConfigRequest(input *GetBucketStyleConfigInput) (req *request.Request, output *GetBucketStyleConfigOutput) {
	op := &request.Operation{
		Name:       opGetBucketStyleConfig,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?styleConfig",
	}

	if input == nil {
		input = &GetBucketStyleConfigInput{}
	}

	output = &GetBucketStyleConfigOutput{}
	req = c.NewRequest(op, input, output)

	return
}

type GetBucketStyleConfigInput struct {
	_ struct{} `type:"structure"`

	Bucket *string `location:"uri" locationName:"Bucket" type:"string"`
}

func (s *GetBucketStyleConfigInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type GetBucketStyleConfigOutput struct {
	_ struct{} `type:"structure" payload:"StyleInfoConfiguration"`

	RequestId *string `location:"header" locationName:"x-amz-request-id" type:"string"`

	StyleInfoConfiguration *StyleInfoConfiguration `locationName:"name" type:"structure"`
}
