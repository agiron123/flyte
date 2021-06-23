/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// An ExecutionSpec encompasses all data used to launch this execution. The Spec does not change over the lifetime of an execution as it progresses across phase changes.
type AdminExecutionSpec struct {
	LaunchPlan *CoreIdentifier `json:"launch_plan,omitempty"`
	Inputs *CoreLiteralMap `json:"inputs,omitempty"`
	Metadata *AdminExecutionMetadata `json:"metadata,omitempty"`
	// List of notifications based on Execution status transitions When this list is not empty it is used rather than any notifications defined in the referenced launch plan. When this list is empty, the notifications defined for the launch plan will be applied.
	Notifications *AdminNotificationList `json:"notifications,omitempty"`
	// This should be set to true if all notifications are intended to be disabled for this execution.
	DisableAll bool `json:"disable_all,omitempty"`
	// Labels to apply to the execution resource.
	Labels *AdminLabels `json:"labels,omitempty"`
	// Annotations to apply to the execution resource.
	Annotations *AdminAnnotations `json:"annotations,omitempty"`
	// Optional: security context override to apply this execution.
	SecurityContext *CoreSecurityContext `json:"security_context,omitempty"`
	// Optional: auth override to apply this execution.
	AuthRole *AdminAuthRole `json:"auth_role,omitempty"`
	// Indicates the runtime priority of the execution.
	QualityOfService *CoreQualityOfService `json:"quality_of_service,omitempty"`
	// Controls the maximum number of task nodes that can be run in parallel for the entire workflow. This is useful to achieve fairness. Note: MapTasks are regarded as one unit, and parallelism/concurrency of MapTasks is independent from this.
	MaxParallelism int32 `json:"max_parallelism,omitempty"`
}
