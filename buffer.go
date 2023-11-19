package wasmgpu

import "syscall/js"

// GPUBufferDescriptor as described:
// https://gpuweb.github.io/gpuweb/#gpubufferdescriptor
type GPUBufferDescriptor struct {
	Size  GPUSize64
	Usage GPUBufferUsageFlags
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUBufferDescriptor) ToJS() any {
	return map[string]any{
		"size":  g.Size.ToJS(),
		"usage": g.Usage.ToJS(),
	}
}

// GPUBuffer as described:
// https://gpuweb.github.io/gpuweb/#gpubuffer
type GPUBuffer struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUBuffer) ToJS() any {
	return g.jsValue
}

// MapAsync as described:
// https://gpuweb.github.io/gpuweb/#dom-gpubuffer-mapasync
func (g GPUBuffer) MapAsync(mode GPUMapModeFlags, offset, size GPUSize64) js.Value {
	params := make([]any, 3)
	params[0] = mode.ToJS()
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
	return g.jsValue.Call("mapAsync", params...)
}

// GetMappedRange as described:
// https://gpuweb.github.io/gpuweb/#dom-gpubuffer-getmappedrange
func (g GPUBuffer) GetMappedRange(offset, size GPUSize64) js.Value {
	params := make([]any, 2)
	if offset > 0 {
		params[0] = offset.ToJS()
	} else {
		params[0] = js.Undefined()
	}
	if size > 0 {
		params[1] = size.ToJS()
	} else {
		params[1] = js.Undefined()
	}
	return g.jsValue.Call("getMappedRange", params...)
}

// Unmap as described:
// https://gpuweb.github.io/gpuweb/#dom-gpubuffer-unmap
func (g GPUBuffer) Unmap() {
	g.jsValue.Call("unmap")
}

// Destroy as described:
// https://gpuweb.github.io/gpuweb/#dom-gpubuffer-destroy
func (g GPUBuffer) Destroy() {
	g.jsValue.Call("destroy")
}
