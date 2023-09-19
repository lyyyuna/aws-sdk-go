package s3

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/restxml"
)

const opPutBucketRemark = "PutBucketRemark"

// 设置空间备注
func (c *S3) PutBucketRemark(input *PutBucketRemarkInput) (*PutBucketRemarkOutput, error) {
	req, out := c.PutBucketRemarkRequest(input)
	return out, req.Send()
}

func (c *S3) PutBucketRemarkRequest(input *PutBucketRemarkInput) (req *request.Request, output *PutBucketRemarkOutput) {
	op := &request.Operation{
		Name:       opPutBucketRemark,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?remark",
	}

	if input == nil {
		input = &PutBucketRemarkInput{}
	}

	output = &PutBucketRemarkOutput{}
	req = c.NewRequest(op, input, output)
	req.Handlers.Build.PushBack(contentMD5)

	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}

type RemarkConfiguration struct {
	_ struct{} `type:"structure"`

	Remark *string `locationName:"Remark" type:"string"`
}

type PutBucketRemarkInput struct {
	_ struct{} `type:"structure"`

	Bucket *string `location:"uri" locationName:"Bucket" type:"string"`

	RemarkConfiguration *RemarkConfiguration `locationName:"RemarkConfiguration" type:"structure"`
}

func (s *PutBucketRemarkInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type PutBucketRemarkOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `location:"header" locationName:"x-amz-request-id" type:"string"`
}

const opGetBucketRemark = "GetBucketRemark"

// 查询空间备注
func (c *S3) GetBucketRemark(input *GetBucketRemarkInput) (*GetBucketRemarkOutput, error) {
	req, out := c.GetBucketRemarkRequest(input)
	return out, req.Send()
}

func (c *S3) GetBucketRemarkRequest(input *GetBucketRemarkInput) (req *request.Request, output *GetBucketRemarkOutput) {
	op := &request.Operation{
		Name:       opGetBucketRemark,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?remark",
	}

	if input == nil {
		input = &GetBucketRemarkInput{}
	}

	output = &GetBucketRemarkOutput{}
	req = c.NewRequest(op, input, output)

	return
}

type GetBucketRemarkInput struct {
	_ struct{} `type:"structure"`

	Bucket *string `location:"uri" locationName:"Bucket" type:"string"`
}

func (s *GetBucketRemarkInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type GetBucketRemarkOutput struct {
	_ struct{} `type:"structure" payload:"RemarkConfiguration"`

	RequestId *string `location:"header" locationName:"x-amz-request-id" type:"string"`

	RemarkConfiguration *RemarkConfiguration `locationName:"RemarkConfiguration" type:"structure"`
}

const opDeleteBucketRemark = "DeleteBucketRemark"

// 删除空间备注
func (c *S3) DeleteBucketRemark(input *DeleteBucketRemarkInput) (*DeleteBucketRemarkOutput, error) {
	req, out := c.DeleteBucketRemarkRequest(input)
	return out, req.Send()
}

func (c *S3) DeleteBucketRemarkRequest(input *DeleteBucketRemarkInput) (req *request.Request, output *DeleteBucketRemarkOutput) {
	op := &request.Operation{
		Name:       opDeleteBucketRemark,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?remark",
	}

	if input == nil {
		input = &DeleteBucketRemarkInput{}
	}

	output = &DeleteBucketRemarkOutput{}
	req = c.NewRequest(op, input, output)

	return
}

type DeleteBucketRemarkInput struct {
	_ struct{} `type:"structure"`

	Bucket *string `location:"uri" locationName:"Bucket" type:"string"`
}

func (s *DeleteBucketRemarkInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type DeleteBucketRemarkOutput struct {
	_ struct{} `type:"structure" payload:"StyleConfiguration"`

	RequestId *string `location:"header" locationName:"x-amz-request-id" type:"string"`
}
