// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `.ProjectIpWhitelist` provides an IP Whitelist entry resource. The whitelist grants access from IPs or CIDRs to clusters within the Project.
// 
// > **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-mongodbatlas/blob/master/website/docs/r/project_ip_whitelist.html.markdown.
type ProjectIpWhitelist struct {
	s *pulumi.ResourceState
}

// NewProjectIpWhitelist registers a new resource with the given unique name, arguments, and options.
func NewProjectIpWhitelist(ctx *pulumi.Context,
	name string, args *ProjectIpWhitelistArgs, opts ...pulumi.ResourceOpt) (*ProjectIpWhitelist, error) {
	if args == nil || args.ProjectId == nil {
		return nil, errors.New("missing required argument 'ProjectId'")
	}
	if args == nil || args.Whitelists == nil {
		return nil, errors.New("missing required argument 'Whitelists'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["projectId"] = nil
		inputs["whitelists"] = nil
	} else {
		inputs["projectId"] = args.ProjectId
		inputs["whitelists"] = args.Whitelists
	}
	s, err := ctx.RegisterResource("mongodbatlas:index/projectIpWhitelist:ProjectIpWhitelist", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ProjectIpWhitelist{s: s}, nil
}

// GetProjectIpWhitelist gets an existing ProjectIpWhitelist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectIpWhitelist(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ProjectIpWhitelistState, opts ...pulumi.ResourceOpt) (*ProjectIpWhitelist, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["projectId"] = state.ProjectId
		inputs["whitelists"] = state.Whitelists
	}
	s, err := ctx.ReadResource("mongodbatlas:index/projectIpWhitelist:ProjectIpWhitelist", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ProjectIpWhitelist{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ProjectIpWhitelist) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ProjectIpWhitelist) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The ID of the project in which to add the whitelist entry.
func (r *ProjectIpWhitelist) ProjectId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["projectId"])
}

func (r *ProjectIpWhitelist) Whitelists() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["whitelists"])
}

// Input properties used for looking up and filtering ProjectIpWhitelist resources.
type ProjectIpWhitelistState struct {
	// The ID of the project in which to add the whitelist entry.
	ProjectId interface{}
	Whitelists interface{}
}

// The set of arguments for constructing a ProjectIpWhitelist resource.
type ProjectIpWhitelistArgs struct {
	// The ID of the project in which to add the whitelist entry.
	ProjectId interface{}
	Whitelists interface{}
}
