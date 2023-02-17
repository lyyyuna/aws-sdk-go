package s3

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/restxml"
)

const opMikuPutBucketLifeCycle = "MikuPutBucketLifeCycle"

// 设置生命周期
func (c *S3) PutMikuBucketLifeCycle(input *PutMikuBucketLifeCycleInput) (*PutMikuBucketLifeCycleOutput, error) {
	req, out := c.PutMikuBucketLifeCycleRequest(input)
	return out, req.Send()
}

func (c *S3) PutMikuBucketLifeCycleRequest(input *PutMikuBucketLifeCycleInput) (req *request.Request, output *PutMikuBucketLifeCycleOutput) {
	op := &request.Operation{
		Name:       opMikuPutBucketLifeCycle,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?lifecycle",
	}

	if input == nil {
		input = &PutMikuBucketLifeCycleInput{}
	}

	output = &PutMikuBucketLifeCycleOutput{}
	req = c.NewRequest(op, input, output)
	req.Handlers.Build.PushBack(contentMD5)

	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}

type PutMikuBucketLifeCycleInput struct {
	_ struct{} `type:"structure"`

	Bucket *string `location:"uri" locationName:"Bucket" type:"string"`

	LifecycleConfiguration *MikuLifecycleConfiguration `locationName:"LifecycleConfiguration" type:"structure"`
}

func (s *PutMikuBucketLifeCycleInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type PutMikuBucketLifeCycleOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `location:"header" locationName:"x-amz-request-id" type:"string"`
}

type MikuLifecycleConfiguration struct {
	_ struct{} `type:"structure"`

	Rules []*MikuRule `locationName:"Rule" type:"list" flattened:"true"`
}

type MikuRule struct {
	_ struct{} `type:"structure"`

	CreateTime *time.Time `locationName:"CreateTime" type:"timestamp"`

	ID *string `locationName:"ID" type:"string"`

	Filter *Filter `locationName:"Filter" type:"structure"`

	Transitions []*MikuTransition `locationName:"Transition" type:"list" flattened:"true"`

	Expiration *Expiration `locationName:"Expiration" type:"structure"`
}

// 过滤条件
type Filter struct {
	_ struct{} `type:"structure"`

	Prefix *string `locationName:"Prefix" type:"string"`
}

// 转换存储类型
type MikuTransition struct {
	_ struct{} `type:"structure"`

	Days *int64 `locationName:"Days" type:"integer"`

	StorageClass *string `locationName:"StorageClass" type:"string"`
}

// 到期删除时间
type Expiration struct {
	Days *int64 `locationName:"Days" type:"integer"`
}

const opMikuGetBucketLifeCycle = "MikuGetBucketLifeCycle"

// 查询生命周期
func (c *S3) GetMikuBucketLifeCycle(input *GetMikuBucketLifeCycleInput) (*GetMikuBucketLifeCycleOutput, error) {
	req, out := c.GetMikuBucketLifeCycleRequest(input)
	return out, req.Send()
}

func (c *S3) GetMikuBucketLifeCycleRequest(input *GetMikuBucketLifeCycleInput) (req *request.Request, output *GetMikuBucketLifeCycleOutput) {
	op := &request.Operation{
		Name:       opMikuGetBucketLifeCycle,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?lifecycle",
	}

	if input == nil {
		input = &GetMikuBucketLifeCycleInput{}
	}

	output = &GetMikuBucketLifeCycleOutput{}
	req = c.NewRequest(op, input, output)

	return
}

type GetMikuBucketLifeCycleInput struct {
	_ struct{} `type:"structure"`

	Bucket *string `location:"uri" locationName:"Bucket" type:"string"`
}

func (s *GetMikuBucketLifeCycleInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type GetMikuBucketLifeCycleOutput struct {
	_ struct{} `type:"structure" payload:"LifecycleConfiguration"`

	RequestId *string `location:"header" locationName:"x-amz-request-id" type:"string"`

	LifecycleConfiguration *MikuLifecycleConfiguration `locationName:"LifecycleConfiguration" type:"structure"`
}

const opMikuDeleteBucketLifeCycle = "MikuDeleteBucketLifeCycle"

// 删除生命周期
func (c *S3) DeleteMikuBucketLifeCycle(input *DeleteMikuBucketLifeCycleInput) (*DeleteMikuBucketLifeCycleOutput, error) {
	req, out := c.DeleteMikuBucketLifeCycleRequest(input)
	return out, req.Send()
}

func (c *S3) DeleteMikuBucketLifeCycleRequest(input *DeleteMikuBucketLifeCycleInput) (req *request.Request, output *DeleteMikuBucketLifeCycleOutput) {
	op := &request.Operation{
		Name:       opMikuDeleteBucketLifeCycle,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?lifecycle",
	}

	if input == nil {
		input = &DeleteMikuBucketLifeCycleInput{}
	}

	output = &DeleteMikuBucketLifeCycleOutput{}
	req = c.NewRequest(op, input, output)

	return
}

type DeleteMikuBucketLifeCycleInput struct {
	_ struct{} `type:"structure"`

	Bucket *string `location:"uri" locationName:"Bucket" type:"string"`
}

func (s *DeleteMikuBucketLifeCycleInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type DeleteMikuBucketLifeCycleOutput struct {
	_ struct{} `type:"structure" payload:"LifecycleConfiguration"`

	RequestId *string `location:"header" locationName:"x-amz-request-id" type:"string"`
}
