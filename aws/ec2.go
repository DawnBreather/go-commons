package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type EC2 struct {
	id string
	region string
}

func (e *EC2) SetId(id string) *EC2{
	e.id = id
	return e
}

func (e *EC2) SetRegion(region string) *EC2{
	e.region = region
	return e
}

func (e *EC2) GetInstanceDescription() *types.Instance{
	client := ec2.NewFromConfig(newConfig())
	input := &ec2.DescribeInstancesInput{
		InstanceIds: []string{
			e.id,
		},
	}

	resp, err := client.DescribeInstances(context.TODO(), input)

	if err != nil {
		_logger.Errorf("Unable to describe instance {%s}: %v", e.id, err)
		return nil
	}

	if len(resp.Reservations) > 0{
		if len(resp.Reservations[0].Instances) > 0 {
			return &resp.Reservations[0].Instances[0]
		}
	}

	return nil
}