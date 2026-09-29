package config

import (
	"context"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Cosmo assigns IDs on creation; reads are done by name and namespace.
	"cosmo_namespace":            config.IdentifierFromProvider,
	"cosmo_federated_graph":      identifierWithPlaceholderID(),
	"cosmo_subgraph":             config.IdentifierFromProvider,
	"cosmo_feature_subgraph":     config.IdentifierFromProvider,
	"cosmo_feature_flag":         config.IdentifierFromProvider,
	"cosmo_monograph":            config.IdentifierFromProvider,
	"cosmo_contract":             identifierWithPlaceholderID(),
	"cosmo_router_token":         config.IdentifierFromProvider,
	"cosmo_persisted_operations": config.IdentifierFromProvider,
}

// placeholderID is used as Terraform ID before the resource is created.
const placeholderID = "00000000-0000-0000-0000-000000000000"

// identifierWithPlaceholderID is IdentifierFromProvider for resources whose
// Terraform Read fails on an empty ID instead of reporting it as not found.
// The Read looks the resource up by name, so the placeholder is never sent.
func identifierWithPlaceholderID() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
		if externalName == "" {
			return placeholderID, nil
		}
		return externalName, nil
	}
	return e
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
