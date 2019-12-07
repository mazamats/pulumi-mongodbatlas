// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `.EncryptionAtRest` Atlas encrypts your data at rest using encrypted storage media. 
// Using keys you manage with AWS KMS, Atlas encrypts your data a second time when it writes it to the MongoDB encrypted storage engine. 
// You can use the following clouds: AWS CMK, AZURE KEY VAULT and GOOGLE KEY VAULT to encrypt the MongoDB master encryption keys.
// 
// > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-mongodbatlas/blob/master/website/docs/r/encryption_at_rest.html.markdown.
type EncryptionAtRest struct {
	s *pulumi.ResourceState
}

// NewEncryptionAtRest registers a new resource with the given unique name, arguments, and options.
func NewEncryptionAtRest(ctx *pulumi.Context,
	name string, args *EncryptionAtRestArgs, opts ...pulumi.ResourceOpt) (*EncryptionAtRest, error) {
	if args == nil || args.ProjectId == nil {
		return nil, errors.New("missing required argument 'ProjectId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["awsKms"] = nil
		inputs["azureKeyVault"] = nil
		inputs["googleCloudKms"] = nil
		inputs["projectId"] = nil
	} else {
		inputs["awsKms"] = args.AwsKms
		inputs["azureKeyVault"] = args.AzureKeyVault
		inputs["googleCloudKms"] = args.GoogleCloudKms
		inputs["projectId"] = args.ProjectId
	}
	s, err := ctx.RegisterResource("mongodbatlas:index/encryptionAtRest:EncryptionAtRest", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &EncryptionAtRest{s: s}, nil
}

// GetEncryptionAtRest gets an existing EncryptionAtRest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEncryptionAtRest(ctx *pulumi.Context,
	name string, id pulumi.ID, state *EncryptionAtRestState, opts ...pulumi.ResourceOpt) (*EncryptionAtRest, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["awsKms"] = state.AwsKms
		inputs["azureKeyVault"] = state.AzureKeyVault
		inputs["googleCloudKms"] = state.GoogleCloudKms
		inputs["projectId"] = state.ProjectId
	}
	s, err := ctx.ReadResource("mongodbatlas:index/encryptionAtRest:EncryptionAtRest", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &EncryptionAtRest{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *EncryptionAtRest) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *EncryptionAtRest) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
func (r *EncryptionAtRest) AwsKms() pulumi.Output {
	return r.s.State["awsKms"]
}

// Specifies Azure Key Vault configuration details and whether Encryption at Rest is enabled for an Atlas project.
func (r *EncryptionAtRest) AzureKeyVault() pulumi.Output {
	return r.s.State["azureKeyVault"]
}

// Specifies GCP KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
func (r *EncryptionAtRest) GoogleCloudKms() pulumi.Output {
	return r.s.State["googleCloudKms"]
}

// The unique identifier for the project.
func (r *EncryptionAtRest) ProjectId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["projectId"])
}

// Input properties used for looking up and filtering EncryptionAtRest resources.
type EncryptionAtRestState struct {
	// Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
	AwsKms interface{}
	// Specifies Azure Key Vault configuration details and whether Encryption at Rest is enabled for an Atlas project.
	AzureKeyVault interface{}
	// Specifies GCP KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
	GoogleCloudKms interface{}
	// The unique identifier for the project.
	ProjectId interface{}
}

// The set of arguments for constructing a EncryptionAtRest resource.
type EncryptionAtRestArgs struct {
	// Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
	AwsKms interface{}
	// Specifies Azure Key Vault configuration details and whether Encryption at Rest is enabled for an Atlas project.
	AzureKeyVault interface{}
	// Specifies GCP KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.
	GoogleCloudKms interface{}
	// The unique identifier for the project.
	ProjectId interface{}
}
