// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package ivschat

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	ivschat_sdkv2 "github.com/aws/aws-sdk-go-v2/service/ivschat"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) EphemeralResources(ctx context.Context) []*types.ServicePackageEphemeralResource {
	return []*types.ServicePackageEphemeralResource{}
}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceLoggingConfiguration,
			TypeName: "aws_ivschat_logging_configuration",
			Name:     "Logging Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceRoom,
			TypeName: "aws_ivschat_room",
			Name:     "Room",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.IVSChat
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*ivschat_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return ivschat_sdkv2.NewFromConfig(cfg,
		ivschat_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
