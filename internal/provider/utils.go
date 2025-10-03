package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// https://developer.hashicorp.com/terraform/plugin/framework/handling-data/types/list#accessing-values
func tfListToStringArray(ctx context.Context, tflist types.List) []string {
	elements := make([]types.String, 0, len(tflist.Elements()))
	tflist.ElementsAs(ctx, &elements, false)

	array := make([]string, len(elements))
	for i, element := range elements {
		array[i] = element.ValueString()
	}

	return array
}
