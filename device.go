package wasmgpu

import "syscall/js"

// NewDevice creates a new GPUDevice that uses the specified JavaScript
// reference of the device.
func NewDevice(jsValue js.Value) GPUDevice {
	return GPUDevice{
		jsValue: jsValue,
	}
}

// GPUDevice as described:
// https://gpuweb.github.io/gpuweb/#gpudevice
type GPUDevice struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUDevice) ToJS() any {
	return g.jsValue
}

// Queue as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-queue
func (g GPUDevice) Queue() GPUQueue {
	jsQueue := g.jsValue.Get("queue")
	return GPUQueue{
		jsValue: jsQueue,
	}
}

// CreateCommandEncoder as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createcommandencoder
func (g GPUDevice) CreateCommandEncoder() GPUCommandEncoder {
	jsEncoder := g.jsValue.Call("createCommandEncoder")
	return GPUCommandEncoder{
		jsValue: jsEncoder,
	}
}

// CreateBuffer as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createbuffer
func (g GPUDevice) CreateBuffer(descriptor GPUBufferDescriptor) GPUBuffer {
	jsBuffer := g.jsValue.Call("createBuffer", descriptor.ToJS())
	return GPUBuffer{
		jsValue: jsBuffer,
	}
}

// CreateShaderModule as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createshadermodule
func (g GPUDevice) CreateShaderModule(desc GPUShaderModuleDescriptor) GPUShaderModule {
	jsShader := g.jsValue.Call("createShaderModule", desc.ToJS())
	return GPUShaderModule{
		jsValue: jsShader,
	}
}

// CreateRenderPipeline as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createrenderpipeline
func (g GPUDevice) CreateRenderPipeline(descriptor GPURenderPipelineDescriptor) GPURenderPipeline {
	jsPipeline := g.jsValue.Call("createRenderPipeline", descriptor.ToJS())
	return GPURenderPipeline{
		jsValue: jsPipeline,
	}
}

// CreateBindGroup as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createbindgroup
func (g GPUDevice) CreateBindGroup(descriptor GPUBindGroupDescriptor) GPUBindGroup {
	jsBindGroup := g.jsValue.Call("createBindGroup", descriptor.ToJS())
	return GPUBindGroup{
		jsValue: jsBindGroup,
	}
}

// CreateBindGroupLayout as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createbindgrouplayout
func (g GPUDevice) CreateBindGroupLayout(descriptor GPUBindGroupLayoutDescriptor) GPUBindGroupLayout {
	jsLayout := g.jsValue.Call("createBindGroupLayout", descriptor.ToJS())
	return GPUBindGroupLayout{
		jsValue: jsLayout,
	}
}

// CreatePipelineLayout as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createpipelinelayout
func (g GPUDevice) CreatePipelineLayout(descriptor GPUPipelineLayoutDescriptor) GPUPipelineLayout {
	jsLayout := g.jsValue.Call("createPipelineLayout", descriptor.ToJS())
	return GPUPipelineLayout{
		jsValue: jsLayout,
	}
}

// CreateComputePipeline as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createcomputepipeline
func (g GPUDevice) CreateComputePipeline(descriptor GPUComputePipelineDescriptor) GPUComputePipeline {
	jsPipeline := g.jsValue.Call("createComputePipeline", descriptor.ToJS())
	return GPUComputePipeline{
		jsValue: jsPipeline,
	}
}
