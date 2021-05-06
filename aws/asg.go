package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type ASG struct{
	name string
	region string
	instances []*types.Instance
}

func (a *ASG) SetName(name string) *ASG {
	a.name = name
	return a
}

func (a *ASG) SetRegion(region string) *ASG {
	a.region = region
	return a
}

func (a *ASG) GetName() string {
	return a.name
}
func (a *ASG) GetRegion() string {
	return a.region
}
func (a *ASG) GetInstances() []*types.Instance {
	return a.instances
}

func (a *ASG) CollectInstances() *ASG{

	client := autoscaling.NewFromConfig(newConfig())
	
	input := &autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: []string{
			a.name,
		},
	}

	var instances []*types.Instance

	resp, err := client.DescribeAutoScalingGroups(context.TODO(), input)
	if err != nil {
		_logger.Errorf("Unable to describe ASG {%s} in region {%s}: %v", a.name, a.region, err)
		return a
	}
	for _, g := range resp.AutoScalingGroups{
		for _, i := range g.Instances{
			e := EC2{}
			e.SetId(*i.InstanceId).
				SetRegion(a.region)
			instance := e.GetInstanceDescription()
			instances = append(instances, instance)
		}
	}
	a.instances = instances

	return a
}