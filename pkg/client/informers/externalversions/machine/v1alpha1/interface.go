// This file was automatically generated by informer-gen

package v1alpha1

import (
	internalinterfaces "github.com/gardener/machine-controller-manager/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AWSMachineClasses returns a AWSMachineClassInformer.
	AWSMachineClasses() AWSMachineClassInformer
	// AliyunMachineClasses returns a AliyunMachineClassInformer.
	AliyunMachineClasses() AliyunMachineClassInformer
	// AzureMachineClasses returns a AzureMachineClassInformer.
	AzureMachineClasses() AzureMachineClassInformer
	// GCPMachineClasses returns a GCPMachineClassInformer.
	GCPMachineClasses() GCPMachineClassInformer
	// Machines returns a MachineInformer.
	Machines() MachineInformer
	// MachineDeployments returns a MachineDeploymentInformer.
	MachineDeployments() MachineDeploymentInformer
	// MachineSets returns a MachineSetInformer.
	MachineSets() MachineSetInformer
	// MachineTemplates returns a MachineTemplateInformer.
	MachineTemplates() MachineTemplateInformer
	// OpenStackMachineClasses returns a OpenStackMachineClassInformer.
	OpenStackMachineClasses() OpenStackMachineClassInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AWSMachineClasses returns a AWSMachineClassInformer.
func (v *version) AWSMachineClasses() AWSMachineClassInformer {
	return &aWSMachineClassInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AliyunMachineClasses returns a AliyunMachineClassInformer.
func (v *version) AliyunMachineClasses() AliyunMachineClassInformer {
	return &aliyunMachineClassInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AzureMachineClasses returns a AzureMachineClassInformer.
func (v *version) AzureMachineClasses() AzureMachineClassInformer {
	return &azureMachineClassInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GCPMachineClasses returns a GCPMachineClassInformer.
func (v *version) GCPMachineClasses() GCPMachineClassInformer {
	return &gCPMachineClassInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Machines returns a MachineInformer.
func (v *version) Machines() MachineInformer {
	return &machineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MachineDeployments returns a MachineDeploymentInformer.
func (v *version) MachineDeployments() MachineDeploymentInformer {
	return &machineDeploymentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MachineSets returns a MachineSetInformer.
func (v *version) MachineSets() MachineSetInformer {
	return &machineSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MachineTemplates returns a MachineTemplateInformer.
func (v *version) MachineTemplates() MachineTemplateInformer {
	return &machineTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OpenStackMachineClasses returns a OpenStackMachineClassInformer.
func (v *version) OpenStackMachineClasses() OpenStackMachineClassInformer {
	return &openStackMachineClassInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
