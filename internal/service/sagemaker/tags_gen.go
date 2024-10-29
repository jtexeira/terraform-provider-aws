// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package sagemaker

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	awstypes "github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/logging"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/types/option"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// listTags lists sagemaker service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func listTags(ctx context.Context, conn *sagemaker.Client, identifier string, optFns ...func(*sagemaker.Options)) (tftags.KeyValueTags, error) {
	input := &sagemaker.ListTagsInput{
		ResourceArn: aws.String(identifier),
	}
	var output []awstypes.Tag

	pages := sagemaker.NewListTagsPaginator(conn, input)
	for pages.HasMorePages() {
		page, err := pages.NextPage(ctx, optFns...)

		if err != nil {
			return tftags.New(ctx, nil), err
		}

		for _, v := range page.Tags {
			output = append(output, v)
		}
	}

	return KeyValueTags(ctx, output), nil
}

// ListTags lists sagemaker service tags and set them in Context.
// It is called from outside this package.
func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier string) error {
	tags, err := listTags(ctx, meta.(*conns.AWSClient).SageMakerClient(ctx), identifier)

	if err != nil {
		return err
	}

	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = option.Some(tags)
	}

	return nil
}

// []*SERVICE.Tag handling

// Tags returns sagemaker service tags.
func Tags(tags tftags.KeyValueTags) []awstypes.Tag {
	result := make([]awstypes.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := awstypes.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from sagemaker service tags.
func KeyValueTags(ctx context.Context, tags []awstypes.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.ToString(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// getTagsIn returns sagemaker service tags from Context.
// nil is returned if there are no input tags.
func getTagsIn(ctx context.Context) []awstypes.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// setTagsOut sets sagemaker service tags in Context.
func setTagsOut(ctx context.Context, tags []awstypes.Tag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = option.Some(KeyValueTags(ctx, tags))
	}
}

// updateTags updates sagemaker service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func updateTags(ctx context.Context, conn *sagemaker.Client, identifier string, oldTagsMap, newTagsMap any, optFns ...func(*sagemaker.Options)) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	ctx = tflog.SetField(ctx, logging.KeyResourceId, identifier)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.SageMaker)
	if len(removedTags) > 0 {
		input := &sagemaker.DeleteTagsInput{
			ResourceArn: aws.String(identifier),
			TagKeys:     removedTags.Keys(),
		}

		_, err := conn.DeleteTags(ctx, input, optFns...)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.SageMaker)
	if len(updatedTags) > 0 {
		input := &sagemaker.AddTagsInput{
			ResourceArn: aws.String(identifier),
			Tags:        Tags(updatedTags),
		}

		_, err := conn.AddTags(ctx, input, optFns...)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// UpdateTags updates sagemaker service tags.
// It is called from outside this package.
func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
	return updateTags(ctx, meta.(*conns.AWSClient).SageMakerClient(ctx), identifier, oldTags, newTags)
}
