package buffer

type GapBuffer struct {
	buff      []byte
	gapSize   int
	buffStart int
	buffEnd   int
}

// Method to create a new buffer and initialize in to a valid state
// Takes a string as a parameter and returns a pointer to the created buffer
func New(source []byte) *GapBuffer {
	size := 20
	contents := append(source, make([]byte, size)...)

	buffer := &GapBuffer{
		buff:      contents,
		gapSize:   size,
		buffStart: len(source),
		buffEnd:   len(contents),
	}
	return buffer
}

// Method to return the buffers contents
// Takes no parameters and returns a string
func (b *GapBuffer) ToString() string {
	left := string(b.buff[:b.buffStart])
	right := string(b.buff[b.buffEnd:])
	return left + right
}

// Method to insert characters into the buffer
func (b *GapBuffer) InsertChar(char byte) {
	if b.gapSize <= 1 {
		b.grow()
	}

	b.buff[b.buffStart] = char
	b.buffStart += 1
	b.gapSize -= 1
}

func (b *GapBuffer) InsertString(text string) {
	for idx := range text {
		b.InsertChar(text[idx])
	}
}

// Method to grow the gap size
// Takes no arguments and return nothing
// A bit janky but I think it does the job
func (b *GapBuffer) grow() {
	tmpLeft := b.buff[:b.buffStart]
	tmpRight := b.buff[b.buffEnd:]
	newGap := make([]byte, (b.size() * 2))
	b.buff = append(b.buff, tmpRight...)
	b.buff = append(b.buff, newGap...)
	b.buff = append(b.buff, tmpLeft...)
	b.gapSize = len(newGap)
}

func (b *GapBuffer) size() int {
	return len(b.buff) - b.gapSize
}
