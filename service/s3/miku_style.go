package s3

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/restxml"
)

const opPutBucketStyle = "PutBucketStyle"

// 设置样式
func (c *S3) PutBucketStyle(input *PutBucketStyleInput) (*PutBucketStyleOutput, error) {
	req, out := c.PutBucketStyleRequest(input)
	return out, req.Send()
}

func (c *S3) PutBucketStyleRequest(input *PutBucketStyleInput) (req *request.Request, output *PutBucketStyleOutput) {
	op := &request.Operation{
		Name:       opPutBucketStyle,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?style",
	}

	if input == nil {
		input = &PutBucketStyleInput{}
	}

	output = &PutBucketStyleOutput{}
	req = c.NewRequest(op, input, output)
	req.Handlers.Build.PushBack(contentMD5)

	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}

type PutBucketStyleInput struct {
	_ struct{} `type:"structure"`

	Bucket *string `location:"uri" locationName:"Bucket" type:"string"`

	StyleConfiguration *StyleConfiguration `locationName:"StyleConfiguration" type:"structure"`
}

func (s *PutBucketStyleInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type PutBucketStyleOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `location:"header" locationName:"x-amz-request-id" type:"string"`
}

type StyleConfiguration struct {
	_ struct{} `type:"structure"`

	Styles []*Style `locationName:"Styles" type:"list" flattened:"true"`
}

type Style struct {
	_ struct{} `type:"structure"`

	Name *string `locationName:"Name" type:"string"`

	PersistentEnable *bool `locationName:"PersistentEnable" type:"boolean"`

	Command *string `locationName:"Command" type:"string"`
}

const opGetBucketStyle = "GetBucketStyle"

// 查询样式
func (c *S3) GetBucketStyle(input *GetBucketStyleInput) (*GetBucketStyleOutput, error) {
	req, out := c.GetBucketStyleRequest(input)
	return out, req.Send()
}

func (c *S3) GetBucketStyleRequest(input *GetBucketStyleInput) (req *request.Request, output *GetBucketStyleOutput) {
	op := &request.Operation{
		Name:       opGetBucketStyle,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?style",
	}

	if input == nil {
		input = &GetBucketStyleInput{}
	}

	output = &GetBucketStyleOutput{}
	req = c.NewRequest(op, input, output)

	return
}

type GetBucketStyleInput struct {
	_ struct{} `type:"structure"`

	Bucket *string `location:"uri" locationName:"Bucket" type:"string"`
}

func (s *GetBucketStyleInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type GetBucketStyleOutput struct {
	_ struct{} `type:"structure" payload:"StyleConfiguration"`

	RequestId *string `location:"header" locationName:"x-amz-request-id" type:"string"`

	StyleConfiguration *StyleConfiguration `locationName:"StyleConfiguration" type:"structure"`
}

const opDeleteBucketStyle = "DeleteBucketStyle"

// 删除样式
func (c *S3) DeleteBucketStyle(input *DeleteBucketStyleInput) (*DeleteBucketStyleOutput, error) {
	req, out := c.DeleteBucketStyleRequest(input)
	return out, req.Send()
}

func (c *S3) DeleteBucketStyleRequest(input *DeleteBucketStyleInput) (req *request.Request, output *DeleteBucketStyleOutput) {
	op := &request.Operation{
		Name:       opDeleteBucketStyle,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?style",
	}

	if input == nil {
		input = &DeleteBucketStyleInput{}
	}

	output = &DeleteBucketStyleOutput{}
	req = c.NewRequest(op, input, output)

	return
}

type DeleteBucketStyleInput struct {
	_ struct{} `type:"structure"`

	Bucket *string `location:"uri" locationName:"Bucket" type:"string"`

	Name *string `location:"querystring" locationName:"name" type:"string" required:"true"`
}

func (s *DeleteBucketStyleInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type DeleteBucketStyleOutput struct {
	_ struct{} `type:"structure" payload:"StyleConfiguration"`

	RequestId *string `location:"header" locationName:"x-amz-request-id" type:"string"`
}
