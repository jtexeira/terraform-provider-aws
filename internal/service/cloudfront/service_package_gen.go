// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	cloudfront_sdkv2 "github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) EphemeralResources(ctx context.Context) []*types.ServicePackageEphemeralResource {
	return []*types.ServicePackageEphemeralResource{}
}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newDataSourceOriginAccessControl,
			Name:    "Origin Access Control",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newContinuousDeploymentPolicyResource,
			Name:    "Continuous Deployment Policy",
		},
		{
			Factory: newKeyValueStoreResource,
			Name:    "Key Value Store",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceCachePolicy,
			TypeName: "aws_cloudfront_cache_policy",
			Name:     "Cache Policy",
		},
		{
			Factory:  dataSourceDistribution,
			TypeName: "aws_cloudfront_distribution",
			Name:     "Distribution",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  dataSourceFunction,
			TypeName: "aws_cloudfront_function",
			Name:     "Function",
		},
		{
			Factory:  dataSourceLogDeliveryCanonicalUserID,
			TypeName: "aws_cloudfront_log_delivery_canonical_user_id",
			Name:     "Log Delivery Canonical User ID",
		},
		{
			Factory:  dataSourceOriginAccessIdentities,
			TypeName: "aws_cloudfront_origin_access_identities",
			Name:     "Origin Access Identities",
		},
		{
			Factory:  dataSourceOriginAccessIdentity,
			TypeName: "aws_cloudfront_origin_access_identity",
			Name:     "Origin Access Identity",
		},
		{
			Factory:  dataSourceOriginRequestPolicy,
			TypeName: "aws_cloudfront_origin_request_policy",
			Name:     "Origin Request Policy",
		},
		{
			Factory:  dataSourceRealtimeLogConfig,
			TypeName: "aws_cloudfront_realtime_log_config",
			Name:     "Real-time Log Config",
		},
		{
			Factory:  dataSourceResponseHeadersPolicy,
			TypeName: "aws_cloudfront_response_headers_policy",
			Name:     "Response Headers Policy",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceCachePolicy,
			TypeName: "aws_cloudfront_cache_policy",
			Name:     "Cache Policy",
		},
		{
			Factory:  resourceDistribution,
			TypeName: "aws_cloudfront_distribution",
			Name:     "Distribution",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceFieldLevelEncryptionConfig,
			TypeName: "aws_cloudfront_field_level_encryption_config",
			Name:     "Field-level Encryption Config",
		},
		{
			Factory:  resourceFieldLevelEncryptionProfile,
			TypeName: "aws_cloudfront_field_level_encryption_profile",
			Name:     "Field-level Encryption Profile",
		},
		{
			Factory:  resourceFunction,
			TypeName: "aws_cloudfront_function",
			Name:     "Function",
		},
		{
			Factory:  resourceKeyGroup,
			TypeName: "aws_cloudfront_key_group",
			Name:     "Key Group",
		},
		{
			Factory:  resourceMonitoringSubscription,
			TypeName: "aws_cloudfront_monitoring_subscription",
			Name:     "Monitoring Subscription",
		},
		{
			Factory:  resourceOriginAccessControl,
			TypeName: "aws_cloudfront_origin_access_control",
			Name:     "Origin Access Control",
		},
		{
			Factory:  resourceOriginAccessIdentity,
			TypeName: "aws_cloudfront_origin_access_identity",
			Name:     "Origin Access Identity",
		},
		{
			Factory:  resourceOriginRequestPolicy,
			TypeName: "aws_cloudfront_origin_request_policy",
			Name:     "Origin Request Policy",
		},
		{
			Factory:  resourcePublicKey,
			TypeName: "aws_cloudfront_public_key",
			Name:     "Public Key",
		},
		{
			Factory:  resourceRealtimeLogConfig,
			TypeName: "aws_cloudfront_realtime_log_config",
			Name:     "Real-time Log Config",
		},
		{
			Factory:  resourceResponseHeadersPolicy,
			TypeName: "aws_cloudfront_response_headers_policy",
			Name:     "Response Headers Policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CloudFront
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*cloudfront_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return cloudfront_sdkv2.NewFromConfig(cfg,
		cloudfront_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
