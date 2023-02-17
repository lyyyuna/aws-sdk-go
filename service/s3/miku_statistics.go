package s3

import "github.com/aws/aws-sdk-go/aws/request"

const opGetBucketStatistic = "GetBucketStatistic"

// 获取镜像回源
func (c *S3) GetBucketStatistic(input *GetBucketStatisticInput) (*GetBucketStatisticOutput, error) {
	req, out := c.GetBucketStatisticRequest(input)
	return out, req.Send()
}

func (c *S3) GetBucketStatisticRequest(input *GetBucketStatisticInput) (req *request.Request, output *GetBucketStatisticOutput) {
	op := &request.Operation{
		Name:       opGetBucketStatistic,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?statistic",
	}

	if input == nil {
		input = &GetBucketStatisticInput{}
	}

	output = &GetBucketStatisticOutput{}
	req = c.NewRequest(op, input, output)

	return
}

type GetBucketStatisticInput struct {
	_ struct{} `type:"structure"`

	Bucket *string `location:"uri" locationName:"Bucket" type:"string"`
}

func (s *GetBucketStatisticInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type GetBucketStatisticOutput struct {
	_ struct{} `type:"structure" payload:"BucketStatistic"`

	BucketStatistic *BucketStatistic `locationName:"MirrorConfiguration" type:"structure"`

	RequestId *string `location:"header" locationName:"x-amz-request-id" type:"string"`
}

type BucketStatistic struct {
	Storage *int64 `locationName:"Storage" type:"integer"`

	ObjectCount *int64 `locationName:"ObjectCount" type:"integer"`

	StandardStorage *int64 `locationName:"StandardStorage" type:"integer"`

	StandardObjectCount *int64 `locationName:"StandardObjectCount" type:"integer"`

	StandardIAStorage *int64 `locationName:"StandardIAStorage" type:"integer"`

	StandardIAObjectCount *int64 `locationName:"StandardIAObjectCount" type:"integer"`

	GlacierStorage *int64 `locationName:"GlacierStorage" type:"integer"`

	GlacierObjectCount *int64 `locationName:"GlacierObjectCount" type:"integer"`

	DeepArchiveStorage *int64 `locationName:"DeepArchiveStorage" type:"integer"`

	DeepArchiveObjectCount *int64 `locationName:"DeepArchiveObjectCount" type:"integer"`
}
