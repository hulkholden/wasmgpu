package wasmgpu

import (
	"syscall/js"

	"github.com/mokiat/gog/opt"
)

// GPUComputePipelineDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpucomputepipelinedescriptor
type GPUComputePipelineDescriptor struct {
	Layout  opt.T[GPUPipelineLayout]
	Compute GPUProgrammableStage
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUComputePipelineDescriptor) ToJS() any {
	result := make(map[string]any)
	if g.Layout.Specified {
		result["layout"] = g.Layout.Value.ToJS()
	} else {
		result["layout"] = "auto"
	}
	result["compute"] = g.Compute.ToJS()
	return result
}

// GPUComputePipeline as described:
// https://gpuweb.github.io/gpuweb/#gpucomputepipeline
type GPUComputePipeline struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUComputePipeline) ToJS() any {
	return g.jsValue
}

// GetBindGroupLayout as described:
// https://gpuweb.github.io/gpuweb/#dom-gpupipelinebase-getbindgrouplayout
func (g GPUComputePipeline) GetBindGroupLayout(index uint32) GPUBindGroupLayout {
	params := make([]any, 1)
	params[0] = index
	return GPUBindGroupLayout{
		jsValue: g.jsValue.Call("getBindGroupLayout", params...),
	}
}
