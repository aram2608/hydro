package buffer

type GapBuffer struct {
	buff []byte
}

// Method to create a new buffer and initialize in to a valid state
// Takes a string as a parameter and returns a pointer to the created buffer
func New(source []byte) *GapBuffer {
	buffer := &GapBuffer{
		buff: source,
	}
	return buffer
}

// Method to return the buffers contents
// Takes no parameters and returns a string
func (b *GapBuffer) ToString() string {
	return string(b.buff)
}
