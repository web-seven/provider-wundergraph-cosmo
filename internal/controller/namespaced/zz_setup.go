// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	contract "github.com/web-seven/provider-wundergraph-cosmo/internal/controller/namespaced/graph/contract"
	cosmonamespace "github.com/web-seven/provider-wundergraph-cosmo/internal/controller/namespaced/graph/cosmonamespace"
	featureflag "github.com/web-seven/provider-wundergraph-cosmo/internal/controller/namespaced/graph/featureflag"
	featuresubgraph "github.com/web-seven/provider-wundergraph-cosmo/internal/controller/namespaced/graph/featuresubgraph"
	federatedgraph "github.com/web-seven/provider-wundergraph-cosmo/internal/controller/namespaced/graph/federatedgraph"
	monograph "github.com/web-seven/provider-wundergraph-cosmo/internal/controller/namespaced/graph/monograph"
	persistedoperations "github.com/web-seven/provider-wundergraph-cosmo/internal/controller/namespaced/graph/persistedoperations"
	routertoken "github.com/web-seven/provider-wundergraph-cosmo/internal/controller/namespaced/graph/routertoken"
	subgraph "github.com/web-seven/provider-wundergraph-cosmo/internal/controller/namespaced/graph/subgraph"
	providerconfig "github.com/web-seven/provider-wundergraph-cosmo/internal/controller/namespaced/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		contract.Setup,
		cosmonamespace.Setup,
		featureflag.Setup,
		featuresubgraph.Setup,
		federatedgraph.Setup,
		monograph.Setup,
		persistedoperations.Setup,
		routertoken.Setup,
		subgraph.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		contract.SetupGated,
		cosmonamespace.SetupGated,
		featureflag.SetupGated,
		featuresubgraph.SetupGated,
		federatedgraph.SetupGated,
		monograph.SetupGated,
		persistedoperations.SetupGated,
		routertoken.SetupGated,
		subgraph.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		contract.SetupWebhookWithManager,
		cosmonamespace.SetupWebhookWithManager,
		featureflag.SetupWebhookWithManager,
		featuresubgraph.SetupWebhookWithManager,
		federatedgraph.SetupWebhookWithManager,
		monograph.SetupWebhookWithManager,
		persistedoperations.SetupWebhookWithManager,
		routertoken.SetupWebhookWithManager,
		subgraph.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
