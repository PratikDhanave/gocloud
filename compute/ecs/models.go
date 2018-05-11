package ecs

//CreateInstance to store all attribute to create Ali-cloud ECS instance
type CreateInstance struct {
	RegionID                string
	ZoneID                  string
	ImageID                 string
	InstanceType            string
	SecurityGroupID         string
	InstanceName            string
	Description             string
	InternetChargeType      string
	InternetMaxBandwidthIn  int
	InternetMaxBandwidthOut int
	HostName                string
	Password                string
	IoOptimized             string
	SystemDiskCategory      string
	SystemDiskSize          string
	SystemDiskName          string
	SystemDiskDescription   string
}

//CreateInstance to store all attribute to start Ali-cloud ECS instance
type StartInstance struct {
	InstanceId string
}

//CreateInstance to store all attribute to Stop Ali-cloud ECS instance
type StopInstance struct {
	InstanceId string
}


//RebootInstance to store all attribute to Reboot Ali-cloud ECS instance
type RebootInstance struct {
	InstanceId string
}


//DeleteInstance to store all attribute to Delete Ali-cloud ECS instance
type DeleteInstance struct {
	InstanceId string
}
