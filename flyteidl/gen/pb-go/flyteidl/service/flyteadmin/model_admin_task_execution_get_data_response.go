/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Response structure for TaskExecutionGetDataRequest which contains inputs and outputs for a task execution.
type AdminTaskExecutionGetDataResponse struct {
	// Signed url to fetch a core.LiteralMap of task execution inputs. Deprecated: Please use full_inputs instead.
	Inputs *AdminUrlBlob `json:"inputs,omitempty"`
	// Signed url to fetch a core.LiteralMap of task execution outputs. Deprecated: Please use full_outputs instead.
	Outputs *AdminUrlBlob `json:"outputs,omitempty"`
	// Full_inputs will only be populated if they are under a configured size threshold.
	FullInputs *CoreLiteralMap `json:"full_inputs,omitempty"`
	// Full_outputs will only be populated if they are under a configured size threshold.
	FullOutputs *CoreLiteralMap `json:"full_outputs,omitempty"`
}
