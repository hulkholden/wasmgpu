package wasmgpu

import (
	"syscall/js"

	"github.com/mokiat/gog"
)

// GPUQueue as described:
// https://gpuweb.github.io/gpuweb/#gpuqueue
type GPUQueue struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUQueue) ToJS() any {
	return g.jsValue
}

// Submit as described:
// https://gpuweb.github.io/gpuweb/#dom-gpuqueue-submit
func (g GPUQueue) Submit(commandBuffers []GPUCommandBuffer) {
	jsSequence := gog.Map(commandBuffers, func(buffer GPUCommandBuffer) any {
		return buffer.ToJS()
	})
	g.jsValue.Call("submit", jsSequence)
}

// WriteBuffer as described:
// https://gpuweb.github.io/gpuweb/#dom-gpuqueue-writebuffer
func (g GPUQueue) WriteBuffer(buffer GPUBuffer, offset uint64, data []byte) {
	dataSize := stageBufferData(data)
	g.jsValue.Call("writeBuffer", buffer.jsValue, offset, uint8Array, uint64(0), dataSize)
}
