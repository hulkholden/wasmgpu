package wasmgpu

import (
	"syscall/js"

	"github.com/mokiat/gog/opt"
)

// GPUCommandEncoder as described:
// https://gpuweb.github.io/gpuweb/#gpucommandencoder
type GPUCommandEncoder struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUCommandEncoder) ToJS() any {
	return g.jsValue
}

// BeginRenderPass as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-beginrenderpass
func (g GPUCommandEncoder) BeginRenderPass(descriptor GPURenderPassDescriptor) GPURenderPassEncoder {
	jsRenderPass := g.jsValue.Call("beginRenderPass", descriptor.ToJS())
	return GPURenderPassEncoder{
		jsValue: jsRenderPass,
	}
}

// BeginComputePass as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-begincomputepass
func (g GPUCommandEncoder) BeginComputePass(descriptor opt.T[GPUComputePassDescriptor]) GPUComputePassEncoder {
	params := make([]any, 1)
	if descriptor.Specified {
		params[0] = descriptor.Value.ToJS()
	} else {
		params[0] = js.Undefined()
	}
	jsComputePass := g.jsValue.Call("beginComputePass", params...)
	return GPUComputePassEncoder{
		jsValue: jsComputePass,
	}
}

// CopyBufferToBuffer as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-copybuffertobuffer
func (g GPUCommandEncoder) CopyBufferToBuffer(source GPUBuffer, sourceOffset GPUSize64, destination GPUBuffer, destinationOffset, size GPUSize64) {
	params := make([]any, 5)
	params[0] = source.ToJS()
	params[1] = sourceOffset.ToJS()
	params[2] = destination.ToJS()
	params[3] = destinationOffset.ToJS()
	params[4] = size.ToJS()
	g.jsValue.Call("copyBufferToBuffer", params...)
}

// ClearBuffer as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-clearbuffer
func (g GPUCommandEncoder) ClearBuffer(source GPUBuffer, offset, size GPUSize64) {
	params := make([]any, 3)
	params[0] = source.ToJS()
	if offset > 0 {
		params[1] = offset.ToJS()
	} else {
		params[1] = js.Undefined()
	}
	if size > 0 {
		params[2] = size.ToJS()
	} else {
		params[2] = js.Undefined()
	}
	g.jsValue.Call("clearBuffer", params...)
}

// Finish as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-finish
func (g GPUCommandEncoder) Finish() GPUCommandBuffer {
	jsBuffer := g.jsValue.Call("finish")
	return GPUCommandBuffer{
		jsValue: jsBuffer,
	}
}
