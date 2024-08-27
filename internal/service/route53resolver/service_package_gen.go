// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package route53resolver

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	route53resolver_sdkv2 "github.com/aws/aws-sdk-go-v2/service/route53resolver"
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
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceEndpoint,
			TypeName: "aws_route53_resolver_endpoint",
			Name:     "Endpoint",
		},
		{
			Factory:  dataSourceFirewallConfig,
			TypeName: "aws_route53_resolver_firewall_config",
			Name:     "Firewall Config",
		},
		{
			Factory:  dataSourceFirewallDomainList,
			TypeName: "aws_route53_resolver_firewall_domain_list",
			Name:     "Firewall Domain List",
		},
		{
			Factory:  dataSourceFirewallRuleGroup,
			TypeName: "aws_route53_resolver_firewall_rule_group",
			Name:     "Firewall Rule Group",
		},
		{
			Factory:  dataSourceFirewallRuleGroupAssociation,
			TypeName: "aws_route53_resolver_firewall_rule_group_association",
			Name:     "Firewall Rule Group Association",
		},
		{
			Factory:  dataSourceResolverFirewallRules,
			TypeName: "aws_route53_resolver_firewall_rules",
			Name:     "Firewall Rules",
		},
		{
			Factory:  dataSourceQueryLogConfig,
			TypeName: "aws_route53_resolver_query_log_config",
			Name:     "Query Log Config",
		},
		{
			Factory:  dataSourceRule,
			TypeName: "aws_route53_resolver_rule",
			Name:     "Rule",
		},
		{
			Factory:  dataSourceRules,
			TypeName: "aws_route53_resolver_rules",
			Name:     "Rules",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceConfig,
			TypeName: "aws_route53_resolver_config",
			Name:     "Config",
		},
		{
			Factory:  resourceDNSSECConfig,
			TypeName: "aws_route53_resolver_dnssec_config",
			Name:     "DNSSEC Config",
		},
		{
			Factory:  resourceEndpoint,
			TypeName: "aws_route53_resolver_endpoint",
			Name:     "Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceFirewallConfig,
			TypeName: "aws_route53_resolver_firewall_config",
			Name:     "Firewall Config",
		},
		{
			Factory:  resourceFirewallDomainList,
			TypeName: "aws_route53_resolver_firewall_domain_list",
			Name:     "Firewall Domain List",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceFirewallRule,
			TypeName: "aws_route53_resolver_firewall_rule",
			Name:     "Firewall Rule",
		},
		{
			Factory:  resourceFirewallRuleGroup,
			TypeName: "aws_route53_resolver_firewall_rule_group",
			Name:     "Firewall Rule Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceFirewallRuleGroupAssociation,
			TypeName: "aws_route53_resolver_firewall_rule_group_association",
			Name:     "Firewall Rule Group Association",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceQueryLogConfig,
			TypeName: "aws_route53_resolver_query_log_config",
			Name:     "Query Log Config",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceQueryLogConfigAssociation,
			TypeName: "aws_route53_resolver_query_log_config_association",
			Name:     "Query Log Config Association",
		},
		{
			Factory:  resourceRule,
			TypeName: "aws_route53_resolver_rule",
			Name:     "Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceRuleAssociation,
			TypeName: "aws_route53_resolver_rule_association",
			Name:     "Rule Association",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Route53Resolver
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*route53resolver_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return route53resolver_sdkv2.NewFromConfig(cfg,
		route53resolver_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
