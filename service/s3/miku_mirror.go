package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/restxml"
)

type Source struct {
	_ struct{} `type:"structure"`

	Addr *string `type:"string"`
}

type Backup struct {
	_ struct{} `type:"structure"`

	Addr *string `type:"string"`
}

type PutBucketMirrorInput struct {
	_ struct{} `type:"structure"`

	Bucket *string `location:"uri" locationName:"Bucket" type:"string"`

	MirrorConfiguration *MirrorConfiguration `locationName:"MirrorConfiguration" type:"structure"`
}

func (s *PutBucketMirrorInput) getBucket() string {
	return *s.Bucket
}

type MirrorConfiguration struct {
	_ struct{} `type:"structure"`

	Sources        []*Source `locationName:"Sources" type:"list" flattened:"true"`
	Backups        []*Backup `locationName:"Backups" type:"list" flattened:"true"`
	Host           *string   `locationName:"Host" type:"string"`
	ForwardQuery   *bool     `locationNme:"ForwardQuery" type:"boolean"`
	ForwardHeaders []*string `locationNme:"ForwardHeaders" type:"list" flattened:"true"`
}

type PutBucketMirrorOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `location:"header" locationName:"x-amz-request-id" type:"string"`
}

const opPutBucketMirror = "PutBucketMirror"

func (c *S3) PutBucketMirrorRequest(input *PutBucketMirrorInput) (req *request.Request, output *PutBucketMirrorOutput) {
	op := &request.Operation{
		Name:       opPutBucketMirror,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?mirror",
	}

	if input == nil {
		input = &PutBucketMirrorInput{}
	}

	output = &PutBucketMirrorOutput{}
	req = c.NewRequest(op, input, output)
	req.Handlers.Build.PushBack(contentMD5)

	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}

// 设置镜像回源
func (c *S3) PutBucketMirror(input *PutBucketMirrorInput) (*PutBucketMirrorOutput, error) {
	req, out := c.PutBucketMirrorRequest(input)
	return out, req.Send()
}

func (c *S3) PutBucketMirrorWithContext(ctx aws.Context, input *PutBucketMirrorInput, opts ...request.Option) (*PutBucketMirrorOutput, error) {
	req, out := c.PutBucketMirrorRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetBucketMirror = "GetBucketMirror"

// 获取镜像回源
func (c *S3) GetBucketMirror(input *GetBucketMirrorInput) (*GetBucketMirrorOutput, error) {
	req, out := c.GetBucketMirrorRequest(input)
	return out, req.Send()
}

func (c *S3) GetBucketMirrorRequest(input *GetBucketMirrorInput) (req *request.Request, output *GetBucketMirrorOutput) {
	op := &request.Operation{
		Name:       opGetBucketMirror,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?mirror",
	}

	if input == nil {
		input = &GetBucketMirrorInput{}
	}

	output = &GetBucketMirrorOutput{}
	req = c.NewRequest(op, input, output)

	return
}

type GetBucketMirrorInput struct {
	_ struct{} `type:"structure"`

	Bucket *string `location:"uri" locationName:"Bucket" type:"string"`
}

func (s *GetBucketMirrorInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type GetBucketMirrorOutput struct {
	_ struct{} `type:"structure" payload:"MirrorConfiguration"`

	MirrorConfiguration *MirrorConfiguration `locationName:"MirrorConfiguration" type:"structure"`

	RequestId *string `location:"header" locationName:"x-amz-request-id" type:"string"`
}

const opDeleteBucketMirror = "DeleteBucketMirror"

// 删除镜像回源
func (c *S3) DeleteBucketMirror(input *DeleteBucketMirrorInput) (*DeleteBucketMirrorOutput, error) {
	req, out := c.DeleteBucketMirrorRequest(input)
	return out, req.Send()
}

func (c *S3) DeleteBucketMirrorRequest(input *DeleteBucketMirrorInput) (req *request.Request, output *DeleteBucketMirrorOutput) {
	op := &request.Operation{
		Name:       opDeleteBucketMirror,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?mirror",
	}

	if input == nil {
		input = &DeleteBucketMirrorInput{}
	}

	output = &DeleteBucketMirrorOutput{}
	req = c.NewRequest(op, input, output)

	return
}

type DeleteBucketMirrorInput struct {
	_ struct{} `type:"structure"`

	Bucket *string `location:"uri" locationName:"Bucket" type:"string"`
}

func (s *DeleteBucketMirrorInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type DeleteBucketMirrorOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `location:"header" locationName:"x-amz-request-id" type:"string"`
}
