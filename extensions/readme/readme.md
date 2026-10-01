# Provider WunderGraph Cosmo

[![CI](https://github.com/web-seven/provider-wundergraph-cosmo/actions/workflows/ci.yml/badge.svg)](https://github.com/web-seven/provider-wundergraph-cosmo/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/web-seven/provider-wundergraph-cosmo?include_prereleases)](https://github.com/web-seven/provider-wundergraph-cosmo/releases)
[![Upbound Marketplace](https://img.shields.io/badge/Upbound-Marketplace-6D64F5)](https://marketplace.upbound.io/providers/web7/provider-wundergraph-cosmo)
[![License](https://img.shields.io/github/license/web-seven/provider-wundergraph-cosmo)](https://github.com/web-seven/provider-wundergraph-cosmo/blob/main/LICENSE)

`provider-wundergraph-cosmo` is a [Crossplane](https://crossplane.io/) provider
built using [Upjet](https://github.com/crossplane/upjet) code generation tools.
It exposes XRM-conformant managed resources for
[WunderGraph Cosmo](https://cosmo-docs.wundergraph.com/).

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

Install the provider:

```console
kubectl apply -f examples/install.yaml
```

Create a secret with the Cosmo API key (`api_url` is optional and defaults to
`https://cosmo-cp.wundergraph.com`):

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: example-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "api_key": "cosmo_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
      "api_url": "https://cosmo-cp.wundergraph.com"
    }
```

Then create a `ProviderConfig` (see `examples/*/providerconfig`) and resources
(see `examples/*/graph`).

## Developing

Run code-generation pipeline:

```console
make generate
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/web-seven/provider-wundergraph-cosmo/issues).
