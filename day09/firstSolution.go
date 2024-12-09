package day09

func firstSolution(data string) int {
	dataSize := len(data)

	blocks := make([]int, dataSize/2+1)
	spaces := make([]int, dataSize/2+1)
	fsSize := 0
	fileId := 0

	for i := range data {
		num := int(data[i] - '0')
		fsSize += num
		if i%2 == 0 {
			blocks[fileId] = num
			fileId++
		} else {
			spaces[fileId-1] = num
		}
	}

	fs := newFilesystem(blocks, spaces, fsSize)
	fs.optimize()

	return fs.computeChecksum()
}

type filesystem struct {
	data []int
	size int
}

func newFilesystem(blocks, spaces []int, size int) *filesystem {
	data := make([]int, size)
	fsIndex := 0

	for fileId, block := range blocks {
		for range block {
			data[fsIndex] = fileId
			fsIndex++
		}
		space := spaces[fileId]
		for range space {
			data[fsIndex] = -1
			fsIndex++
		}
	}

	return &filesystem{data, size}
}

func (fs *filesystem) optimize() {
	li := 0
	ri := fs.size - 1

	for li < ri {

		if fs.data[li] != -1 {
			li++
			continue
		}
		if fs.data[ri] == -1 {
			ri--
			continue
		}
		fs.data[li] = fs.data[ri]
		fs.data[ri] = -1
		li++
		ri--
	}
}

func (fs filesystem) computeChecksum() int {
	checksum := 0
	for i, id := range fs.data {
		if id == -1 {
			break
		}
		checksum += i * id

	}
	return checksum
}
