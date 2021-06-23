/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Defines a set of overridable task resource attributes set during task registration.
type AdminTaskResourceSpec struct {
	Cpu string `json:"cpu,omitempty"`
	Gpu string `json:"gpu,omitempty"`
	Memory string `json:"memory,omitempty"`
	Storage string `json:"storage,omitempty"`
}
