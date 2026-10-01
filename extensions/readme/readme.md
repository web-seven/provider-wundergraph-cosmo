# Provider WunderGraph Cosmo

[![CI](https://github.com/web-seven/provider-wundergraph-cosmo/actions/workflows/ci.yml/badge.svg)](https://github.com/web-seven/provider-wundergraph-cosmo/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/web-seven/provider-wundergraph-cosmo?include_prereleases)](https://github.com/web-seven/provider-wundergraph-cosmo/releases)
[![Upbound Marketplace](https://img.shields.io/badge/Upbound-Marketplace-6D64F5)](https://marketplace.upbound.io/providers/web7/provider-wundergraph-cosmo)
[![License](https://img.shields.io/github/license/web-seven/provider-wundergraph-cosmo)](https://github.com/web-seven/provider-wundergraph-cosmo/blob/main/LICENSE)

`provider-wundergraph-cosmo` is a [Crossplane](https://crossplane.io/) provider
that manages [WunderGraph Cosmo](https://cosmo-docs.wundergraph.com/)
namespaces, federated graphs, subgraphs and related resources as Kubernetes
objects.

## Resources

All resources are available as cluster-scoped (`graph.cosmo.crossplane.io`)
and namespaced (`graph.cosmo.m.crossplane.io`) kinds.

| Kind                  | Description          |
|-----------------------|----------------------|
| `CosmoNamespace`      | Cosmo namespace      |
| `FederatedGraph`      | Federated graph      |
| `Subgraph`            | Subgraph             |
| `FeatureSubgraph`     | Feature subgraph     |
| `FeatureFlag`         | Feature flag         |
| `Monograph`           | Monograph            |
| `Contract`            | Contract graph       |
| `RouterToken`         | Router token         |
| `PersistedOperations` | Persisted operations |

The Cosmo namespace kind is `CosmoNamespace` because a cluster-scoped kind with
the plural `namespaces` cannot be served by the Kubernetes API.

`namespace` is required on all resources except `CosmoNamespace`. It can be set
directly or resolved with `namespaceRef`/`namespaceSelector`. References to
graphs and subgraphs (`sourceRef`, `graphRef`, `federatedGraphRef`,
`baseSubgraphRef`, `featureSubgraphsRefs`) resolve to Cosmo names.

`RouterToken` writes the generated `token` to its connection secret.

## Getting Started

### Install the provider

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-wundergraph-cosmo
spec:
  package: xpkg.upbound.io/web7/provider-wundergraph-cosmo:v0.1.2
```

### Configure credentials

Create a secret with a Cosmo API key (`api_url` is optional and defaults to
`https://cosmo-cp.wundergraph.com`):

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: cosmo-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "api_key": "cosmo_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
      "api_url": "https://cosmo-cp.wundergraph.com"
    }
```

Then create a `ProviderConfig` that points to it. The examples below use the
namespaced kinds, so the `ProviderConfig` lives in the same namespace as the
resources:

```yaml
apiVersion: cosmo.m.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
  namespace: crossplane-system
spec:
  credentials:
    source: Secret
    secretRef:
      name: cosmo-creds
      namespace: crossplane-system
      key: credentials
```

A `ClusterProviderConfig` (`cosmo.m.crossplane.io/v1beta1`) with the same spec
can be shared by resources in all namespaces.

### Create a Cosmo namespace

Federated graphs and subgraphs belong to a Cosmo namespace:

```yaml
apiVersion: graph.cosmo.m.crossplane.io/v1alpha1
kind: CosmoNamespace
metadata:
  name: production
  namespace: crossplane-system
spec:
  forProvider:
    name: production
  providerConfigRef:
    kind: ProviderConfig
    name: default
```

### Create a federated graph

A federated graph composes every subgraph whose labels match its
`labelMatchers`. `routingUrl` is the URL of the router that serves the graph:

```yaml
apiVersion: graph.cosmo.m.crossplane.io/v1alpha1
kind: FederatedGraph
metadata:
  name: shop
  namespace: crossplane-system
spec:
  forProvider:
    name: shop
    routingUrl: https://router.example.com/graphql
    labelMatchers:
      - team=shop
    namespaceRef:
      name: production
  providerConfigRef:
    kind: ProviderConfig
    name: default
```

### Create subgraphs

Each subgraph is published with its schema and labelled so the federated graph
picks it up:

```yaml
apiVersion: graph.cosmo.m.crossplane.io/v1alpha1
kind: Subgraph
metadata:
  name: products
  namespace: crossplane-system
spec:
  forProvider:
    name: products
    routingUrl: https://products.example.com/graphql
    labels:
      team: shop
    schema: |
      type Query {
        products: [Product!]!
      }

      type Product @key(fields: "id") {
        id: ID!
        name: String!
      }
    namespaceRef:
      name: production
  providerConfigRef:
    kind: ProviderConfig
    name: default
---
apiVersion: graph.cosmo.m.crossplane.io/v1alpha1
kind: Subgraph
metadata:
  name: reviews
  namespace: crossplane-system
spec:
  forProvider:
    name: reviews
    routingUrl: https://reviews.example.com/graphql
    labels:
      team: shop
    schema: |
      type Product @key(fields: "id") {
        id: ID!
        reviews: [Review!]!
      }

      type Review {
        rating: Int!
        body: String
      }
    namespaceRef:
      name: production
  providerConfigRef:
    kind: ProviderConfig
    name: default
```

### Create a router token

The router needs a token to fetch the federated graph configuration. The token
is written to the connection secret:

```yaml
apiVersion: graph.cosmo.m.crossplane.io/v1alpha1
kind: RouterToken
metadata:
  name: shop-router
  namespace: crossplane-system
spec:
  forProvider:
    name: shop-router
    graphRef:
      name: shop
    namespaceRef:
      name: production
  providerConfigRef:
    kind: ProviderConfig
    name: default
  writeConnectionSecretToRef:
    name: shop-router-token
```

Check that everything is synced and ready:

```console
kubectl get managed
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/web-seven/provider-wundergraph-cosmo/issues).
