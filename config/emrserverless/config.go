package emrserverless

import (
	"context"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/upbound/upjet/pkg/config"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	fieldPathApplicationName = "spec.forProvider.name"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_emrserverless_application", func(r *config.Resource) {
		r.InitializerFns = append(r.InitializerFns, func(c client.Client) managed.Initializer {
			return &initializeNameParameter{client: c}
		})
		if t, ok := r.TerraformResource.Schema["name"]; ok {
			t.Optional = true
		}
	})
}

type initializeNameParameter struct {
	client client.Client
}

func (i *initializeNameParameter) Initialize(ctx context.Context, mg xpresource.Managed) error {
	paved, err := fieldpath.PaveObject(mg)
	if err != nil {
		return err
	}
	// If the value isn't already set, set it to the name of the
	// managed resource
	name, err := paved.GetString(fieldPathApplicationName)
	if ignoreNotFound(err) != nil {
		return err
	}

	if name == "" {
		if err := paved.SetString(fieldPathApplicationName, mg.GetName()); err != nil {
			return err
		}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(paved.UnstructuredContent(), mg); err != nil {
			return err
		}
		if err := i.client.Update(ctx, mg); err != nil {
			return err
		}
	}
	return nil
}

func ignoreNotFound(err error) error {
	if fieldpath.IsNotFound(err) {
		return nil
	}
	return err
}
