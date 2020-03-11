/*
Copyright 2019 The xridge kubestone contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// S3BenchSpec defines the desired state of S3Bench
type S3BenchSpec struct {
	// Image defines the warp docker image used for the benchmark
	// +optional
	Image ImageSpec `json:"image,omitempty"`

	// PodConfig contains the configuration for the benchmark pod, including
	// pod labels and scheduling policies (affinity, toleration, node selector...)
	// +optional
	PodConfig PodConfigurationSpec `json:"podConfig,omitempty"`

	// Mode defines the operating mode of the benchmark test. See https://github.com/minio/warp#mixed for option definition.
	// Currently accepted values are: get, put, delete, mixed
	Mode string `json:"mode"`

	// Host defines the host to benchmark against.
	// Multiple hosts can be specified as a comma separated list. (default: "127.0.0.1:9000")
	Host string `json:"host"`

	// S3BenchOptions defines the runtime arguments for the benchmark test
	// +optional
	S3BenchOptions S3BenchOptions `json:"options"`

	// S3ObjectOptions defines options for the objects generated by the benchmark
	// +optional
	S3ObjectOptions S3ObjectOptions `json:"objects,omitempty"`

	// S3AutoTermOptions defines options for the auto terminate feature of warp.
	// +optional
	S3AutoTermOptions S3AutoTermOptions `json:"auto_term,omitempty"`

	// S3AnalysisOptions defines options for the analysis features of warp
	// +optional
	S3AnalysisOptions S3AnalysisOptions `json:"analysis,omitempty"`

	// MixedDistributionOptions defines the distribution of operation types if using the mixed mode
	// Will only be used in "mixed" mode.
	// +optional
	MixedDistributionOptions MixedDistributionOptions `json:"mixed_dist,omitempty"`
}

// S3BenchOptions defines the runtime arguments for the Warp cli
type S3BenchOptions struct {
	// NoColor will disable color theme (default: false)
	// +optional
	NoColor bool `json:"no_color,omitempty"`

	// Debug will enable debug output (default: false)
	// +optional
	Debug bool `json:"debug,omitempty"`

	// Insecure defines if to disable SSL certificate verification (default: false)
	// +optional
	Insecure bool `json:"insecure,omitempty"`

	// +optional
	AccessKey string `json:"access_key,omitempty"`

	// +optional
	SecretKey string `json:"secret_key,omitempty"`

	// Tls defines if to use TLS (HTTPS) for transport (default: false)
	// +optional
	Tls bool `json:"tls,omitempty"`

	// Region defines a custom region
	// +optional
	Region string `json:"region,omitempty"`

	// Encrypt defines if to encrypt/decrypt objects (using server-side encryption with random keys)
	// (default: false)
	// +optional
	Encrypt bool `json:"encrypt,omitempty"`

	// Bucket defines which bucket to use for benchmark data. ALL DATA WILL BE DELETED IN BUCKET!
	// (default: "warp-benchmark-bucket")
	// +optional
	Bucket string `json:"bucket,omitempty"`

	// HostSelect defines the host selection algorithm. Can be "weighed" or "roundrobin" (default: "weighed")
	// +optional
	HostSelect string `json:"host_select,omitempty"`

	// Concurrent defines how many concurrent operations to run (default: 6)
	// +optional
	Concurrent int32 `json:"concurrent,omitempty"`

	// NoPrefix defines if to NOT use separate prefix for each thread (default: false)
	// +optional
	NoPrefix bool `json:"no_prefix,omitempty"`

	// Output benchmark+profile data to this file. By default unique filename is generated.
	// +optional
	BenchOutput string `json:"bench_output,omitempty"`

	// Duration defines the time length to run the benchmark. Use 's' and 'm' to specify seconds and minutes.
	// (default: 5m0s)
	// +optional
	Duration string `json:"duration,omitempty"`

	// NoClear Do not clear bucket before or after running benchmarks. Use when running multiple clients.
	// (default: false)
	// +optional
	NoClear bool `json:"no_clear,omitempty"`

	// Specify a benchmark start time. Time format is 'hh:mm' where hours are specified in 24h format, server TZ.
	// +optional
	SyncStart string `json:"sync_start,omitempty"`

	// Requests Display individual request stats.
	// +optional
	Requests bool `json:"requests,omitempty"`
}

// S3ObjectOptions defines options for the objects generated by the benchmark
type S3ObjectOptions struct {
	// Count defines the number of objects to upload. (default: 2500)
	// +optional
	Count int32 `json:"count,omitempty"`

	// Size defines the size of each generated object. Can be a number or 10KiB/MiB/GiB.
	// All sizes are base 2 binary. (default: "10MiB")
	// +optional
	Size string `json:"size,omitempty"`

	// Generator defines if to use a specific data generator (default: "random")
	// +optional
	Generator string `json:"generator,omitempty"`

	// RandomSize defines if to randomize size of objects so they will be up to the specified size
	// (default: false)
	// +optional
	RandomSize bool `json:"random_size,omitempty"`
}

// S3AnalysisOptions defines options for the analysis features of warp
type S3AnalysisOptions struct {
	// Duration defines the time length to split analysis into durations of this length (default: "1s")
	// +optional
	Duration string `json:"duration,omitempty"`
	// Output aggregated data as to file
	// +optional
	Output string `json:"output,omitempty"`
	// OperationFilter Only output for this op. Can be GET/PUT/DELETE, etc.
	// +optional
	OperationFilter string `json:"operation_filter,omitempty"`
	// PrintErrors Print out errors (default: false)
	// +optional
	PrintErrors bool `json:"print_errors,omitempty"`
	// HostFilter Only output for this host.
	// +optional
	HostFilter string `json:"host_filter,omitempty"`
	// Skip Additional duration to skip when analyzing data. (default: 0s)
	// +optional
	Skip string `json:"skip,omitempty"`
	// HostDetails Do detailed time segmentation per host (default: false)
	// +optional
	HostDetails bool `json:"host_details,omitempty"`
}

// MixedDistributionOptions defines the distribution of operation types if using the mixed mode
type MixedDistributionOptions struct {
	// GetDist The amount of GET operations. (default: 45)
	// +optional
	GetDist int32 `json:"get,omitempty"`
	// StatDist The amount of STAT operations. (default: 30)
	// +optional
	StatDist int32 `json:"stat,omitempty"`
	// PutDist The amount of PUT operations. (default: 15)
	// +optional
	PutDist int32 `json:"put,omitempty"`
	// DeleteDist The amount of DELETE operations. Must be at least the same as PUT. (default: 10)
	// +optional
	DeleteDist int32 `json:"delete,omitempty"`
}

// S3AutoTermOptions defines options for the auto terminate feature of warp.
type S3AutoTermOptions struct {

	// Enabled defines if to auto terminate when benchmark is considered stable. (default: false)
	// +optional
	Enabled bool `json:"enabled,omitempty"`

	// Duration defines the minimum duration where output must have been stable to allow automatic termination. (default: 10s)
	// +optional
	Duration string `json:"duration,omitempty"`

	// Percent defines the percentage the last 6/25 time blocks must be within current speed to auto terminate. (default: 7.5)
	// +optional
	Percent string `json:"percent,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Running",type="boolean",JSONPath=".status.running"
// +kubebuilder:printcolumn:name="Completed",type="boolean",JSONPath=".status.completed"

// S3Bench is the Schema for the s3benches API
type S3Bench struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   S3BenchSpec     `json:"spec,omitempty"`
	Status BenchmarkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3BenchList contains a list of S3Bench
type S3BenchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3Bench `json:"items"`
}

func init() {
	SchemeBuilder.Register(&S3Bench{}, &S3BenchList{})
}
