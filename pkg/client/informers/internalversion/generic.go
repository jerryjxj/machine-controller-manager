// This file was automatically generated by informer-gen

package internalversion

import (
	"fmt"

	machine "github.com/gardener/machine-controller-manager/pkg/apis/machine"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=machine.sapcloud.io, Version=internalVersion
	case machine.SchemeGroupVersion.WithResource("awsmachineclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Machine().InternalVersion().AWSMachineClasses().Informer()}, nil
	case machine.SchemeGroupVersion.WithResource("aliyunmachineclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Machine().InternalVersion().AliyunMachineClasses().Informer()}, nil
	case machine.SchemeGroupVersion.WithResource("azuremachineclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Machine().InternalVersion().AzureMachineClasses().Informer()}, nil
	case machine.SchemeGroupVersion.WithResource("gcpmachineclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Machine().InternalVersion().GCPMachineClasses().Informer()}, nil
	case machine.SchemeGroupVersion.WithResource("machines"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Machine().InternalVersion().Machines().Informer()}, nil
	case machine.SchemeGroupVersion.WithResource("machinedeployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Machine().InternalVersion().MachineDeployments().Informer()}, nil
	case machine.SchemeGroupVersion.WithResource("machinesets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Machine().InternalVersion().MachineSets().Informer()}, nil
	case machine.SchemeGroupVersion.WithResource("machinetemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Machine().InternalVersion().MachineTemplates().Informer()}, nil
	case machine.SchemeGroupVersion.WithResource("openstackmachineclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Machine().InternalVersion().OpenStackMachineClasses().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
