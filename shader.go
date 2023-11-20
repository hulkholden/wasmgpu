package wasmgpu

import (
	"syscall/js"

	"github.com/mokiat/gog/opt"
)

// GPUProgrammableStage as described:
// https://gpuweb.github.io/gpuweb/#gpuprogrammablestage
type GPUProgrammableStage struct {
	Module     GPUShaderModule
	EntryPoint string
	Constants  opt.T[GPUProgrammableStageConstants]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUProgrammableStage) ToJS() any {
	result := map[string]any{
		"module":     g.Module.ToJS(),
		"entryPoint": g.EntryPoint,
	}
	if g.Constants.Specified {
		result["constants"] = g.Constants.Value.ToJS()
	}
	return result
}

// GPUShaderModuleDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpushadermoduledescriptor
type GPUShaderModuleDescriptor struct {
	Code string
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUShaderModuleDescriptor) ToJS() any {
	return map[string]any{
		"code": g.Code,
	}
}

// GPUShaderModule as described:
// https://gpuweb.github.io/gpuweb/#gpushadermodule
type GPUShaderModule struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUShaderModule) ToJS() any {
	return g.jsValue
}

// GPUProgrammableStageConstants as described:
// https://developer.mozilla.org/en-US/docs/Web/API/GPUDevice/createComputePipeline#constants
type GPUProgrammableStageConstants map[any]any

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUProgrammableStageConstants) ToJS() any {
	o := objectCtor.New()
	for k, v := range g {
		o.Set(k, v)
	}
	return o
}
