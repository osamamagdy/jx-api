// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/jenkins-x/jx-api/v1/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Apps returns a AppInformer.
	Apps() AppInformer
	// BuildPacks returns a BuildPackInformer.
	BuildPacks() BuildPackInformer
	// CommitStatuses returns a CommitStatusInformer.
	CommitStatuses() CommitStatusInformer
	// Environments returns a EnvironmentInformer.
	Environments() EnvironmentInformer
	// EnvironmentRoleBindings returns a EnvironmentRoleBindingInformer.
	EnvironmentRoleBindings() EnvironmentRoleBindingInformer
	// Extensions returns a ExtensionInformer.
	Extensions() ExtensionInformer
	// Facts returns a FactInformer.
	Facts() FactInformer
	// GitServices returns a GitServiceInformer.
	GitServices() GitServiceInformer
	// PipelineActivities returns a PipelineActivityInformer.
	PipelineActivities() PipelineActivityInformer
	// PipelineStructures returns a PipelineStructureInformer.
	PipelineStructures() PipelineStructureInformer
	// Plugins returns a PluginInformer.
	Plugins() PluginInformer
	// Releases returns a ReleaseInformer.
	Releases() ReleaseInformer
	// Schedulers returns a SchedulerInformer.
	Schedulers() SchedulerInformer
	// SourceRepositories returns a SourceRepositoryInformer.
	SourceRepositories() SourceRepositoryInformer
	// SourceRepositoryGroups returns a SourceRepositoryGroupInformer.
	SourceRepositoryGroups() SourceRepositoryGroupInformer
	// Teams returns a TeamInformer.
	Teams() TeamInformer
	// Users returns a UserInformer.
	Users() UserInformer
	// Workflows returns a WorkflowInformer.
	Workflows() WorkflowInformer
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

// Apps returns a AppInformer.
func (v *version) Apps() AppInformer {
	return &appInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BuildPacks returns a BuildPackInformer.
func (v *version) BuildPacks() BuildPackInformer {
	return &buildPackInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CommitStatuses returns a CommitStatusInformer.
func (v *version) CommitStatuses() CommitStatusInformer {
	return &commitStatusInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Environments returns a EnvironmentInformer.
func (v *version) Environments() EnvironmentInformer {
	return &environmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EnvironmentRoleBindings returns a EnvironmentRoleBindingInformer.
func (v *version) EnvironmentRoleBindings() EnvironmentRoleBindingInformer {
	return &environmentRoleBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Extensions returns a ExtensionInformer.
func (v *version) Extensions() ExtensionInformer {
	return &extensionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Facts returns a FactInformer.
func (v *version) Facts() FactInformer {
	return &factInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GitServices returns a GitServiceInformer.
func (v *version) GitServices() GitServiceInformer {
	return &gitServiceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PipelineActivities returns a PipelineActivityInformer.
func (v *version) PipelineActivities() PipelineActivityInformer {
	return &pipelineActivityInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PipelineStructures returns a PipelineStructureInformer.
func (v *version) PipelineStructures() PipelineStructureInformer {
	return &pipelineStructureInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Plugins returns a PluginInformer.
func (v *version) Plugins() PluginInformer {
	return &pluginInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Releases returns a ReleaseInformer.
func (v *version) Releases() ReleaseInformer {
	return &releaseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Schedulers returns a SchedulerInformer.
func (v *version) Schedulers() SchedulerInformer {
	return &schedulerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SourceRepositories returns a SourceRepositoryInformer.
func (v *version) SourceRepositories() SourceRepositoryInformer {
	return &sourceRepositoryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SourceRepositoryGroups returns a SourceRepositoryGroupInformer.
func (v *version) SourceRepositoryGroups() SourceRepositoryGroupInformer {
	return &sourceRepositoryGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Teams returns a TeamInformer.
func (v *version) Teams() TeamInformer {
	return &teamInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Users returns a UserInformer.
func (v *version) Users() UserInformer {
	return &userInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Workflows returns a WorkflowInformer.
func (v *version) Workflows() WorkflowInformer {
	return &workflowInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
