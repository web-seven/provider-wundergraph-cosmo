package cosmo

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const (
	// shortGroup is the API group shared by all Cosmo resources.
	shortGroup = "graph"

	// extractName extracts the name parameter of the referenced resource.
	// Cosmo resources are referenced by name within a namespace, not by ID.
	extractName = `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("name",false)`
)

// Configure configures the Cosmo resources.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("cosmo_namespace", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CosmoNamespace"
	})

	p.AddResourceConfigurator("cosmo_federated_graph", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "FederatedGraph"
		namespaceRef(r)
	})

	p.AddResourceConfigurator("cosmo_subgraph", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Subgraph"
		namespaceRef(r)
	})

	p.AddResourceConfigurator("cosmo_feature_subgraph", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "FeatureSubgraph"
		namespaceRef(r)
		r.References["base_subgraph_name"] = ujconfig.Reference{
			TerraformName:     "cosmo_subgraph",
			Extractor:         extractName,
			RefFieldName:      "BaseSubgraphRef",
			SelectorFieldName: "BaseSubgraphSelector",
		}
	})

	p.AddResourceConfigurator("cosmo_feature_flag", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "FeatureFlag"
		namespaceRef(r)
		r.References["feature_subgraphs"] = ujconfig.Reference{
			TerraformName: "cosmo_feature_subgraph",
			Extractor:     extractName,
		}
	})

	p.AddResourceConfigurator("cosmo_monograph", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Monograph"
		namespaceRef(r)
		r.TerraformResource.Schema["admission_webhook_secret"].Sensitive = true
	})

	p.AddResourceConfigurator("cosmo_contract", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Contract"
		namespaceRef(r)
		r.References["source"] = ujconfig.Reference{
			TerraformName:     "cosmo_federated_graph",
			Extractor:         extractName,
			RefFieldName:      "SourceRef",
			SelectorFieldName: "SourceSelector",
		}
		r.TerraformResource.Schema["admission_webhook_secret"].Sensitive = true
	})

	p.AddResourceConfigurator("cosmo_router_token", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RouterToken"
		namespaceRef(r)
		r.References["graph_name"] = ujconfig.Reference{
			TerraformName:     "cosmo_federated_graph",
			Extractor:         extractName,
			RefFieldName:      "GraphRef",
			SelectorFieldName: "GraphSelector",
		}
	})

	p.AddResourceConfigurator("cosmo_persisted_operations", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PersistedOperations"
		namespaceRef(r)
		r.References["federated_graph_name"] = ujconfig.Reference{
			TerraformName:     "cosmo_federated_graph",
			Extractor:         extractName,
			RefFieldName:      "FederatedGraphRef",
			SelectorFieldName: "FederatedGraphSelector",
		}
	})
}

// namespaceRef makes the namespace argument required and referenceable to a
// Namespace. The Terraform default is not applied before creation, and some
// resources cannot be read without a namespace.
func namespaceRef(r *ujconfig.Resource) {
	s := r.TerraformResource.Schema["namespace"]
	s.Optional = false
	s.Computed = false
	s.Required = true
	r.References["namespace"] = ujconfig.Reference{
		TerraformName:     "cosmo_namespace",
		Extractor:         extractName,
		RefFieldName:      "NamespaceRef",
		SelectorFieldName: "NamespaceSelector",
	}
}
