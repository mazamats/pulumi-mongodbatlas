module github.com/pulumi/pulumi-mongodbatlas

go 1.12

replace (
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
)

require (
	github.com/hashicorp/terraform v0.12.9
	github.com/hashicorp/terraform-plugin-sdk v1.4.0
	github.com/hashicorp/vault v1.2.3 // indirect
	github.com/pulumi/pulumi v1.6.0
	github.com/pulumi/pulumi-terraform v0.18.4-0.20191202134852-87cfb4dc8ae1
	github.com/pulumi/pulumi-terraform-bridge v1.5.0
	github.com/pulumi/scripts v0.0.0-20191031001615-ae6f56d5f37f // indirect
	github.com/stretchr/testify v1.4.0
	github.com/terraform-providers/terraform-provider-mongodbatlas v0.3.1
)
