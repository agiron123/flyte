/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Request to register a launch plan. The included LaunchPlanSpec may have a complete or incomplete set of inputs required to launch a workflow execution. By default all launch plans are registered in state INACTIVE. If you wish to set the state to ACTIVE, you must submit a LaunchPlanUpdateRequest, after you have successfully created a launch plan.
type AdminLaunchPlanCreateRequest struct {
	// Uniquely identifies a launch plan entity.
	Id *CoreIdentifier `json:"id,omitempty"`
	// User-provided launch plan details, including reference workflow, inputs and other metadata.
	Spec *AdminLaunchPlanSpec `json:"spec,omitempty"`
}
