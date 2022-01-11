/*
Copyright 2015-2021 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"context"

	"github.com/gravitational/teleport-plugins/terraform/tfschema"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/trace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceTeleportAuthPreference returns Teleport cluster auth preference
func dataSourceTeleportAuthPreference() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTeleportAuthPreferenceRead,
		Schema:      tfschema.SchemaAuthPreferenceV2,
	}
}

// dataSourceTeleportAuthPreferenceRead reads Teleport cluster auth preference
func dataSourceTeleportAuthPreferenceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, err := getClient(m)
	if err != nil {
		return diagFromErr(err)
	}

	g, err := c.GetAuthPreference(ctx)
	if err != nil {
		return diagFromErr(describeErr(err, "cluster_auth_preference"))
	}

	g3, ok := g.(*types.AuthPreferenceV2)
	if !ok {
		return diagFromErr(trace.Errorf("can not convert %T to *types.AuthPreferenceV2", g))
	}

	err = tfschema.ToTerraformAuthPreferenceV2(g3, d)
	if err != nil {
		return diagFromErr(err)
	}

	return diag.Diagnostics{}
}
