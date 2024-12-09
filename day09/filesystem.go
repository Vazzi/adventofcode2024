package day09

type filesystem struct {
	data       []int
	movedFiles []bool
	size       int
}

func newFilesystem(blocks, spaces []int, size int) *filesystem {
	data := make([]int, size)
	movedFiles := make([]bool, size)
	fsIndex := 0

	for fileId, block := range blocks {
		for range block {
			data[fsIndex] = fileId
			fsIndex++
		}
		space := spaces[fileId]
		for range space {
			data[fsIndex] = freeSpace
			fsIndex++
		}
	}

	return &filesystem{data, movedFiles, size}
}

func (fs filesystem) computeChecksum() int {
	checksum := 0
	for i, id := range fs.data {
		if id < 0 {
			continue
		}
		checksum += i * id

	}
	return checksum
}

func (fs *filesystem) markMovedFile(id int) {
	fs.movedFiles[id] = true
}

const (
	freeSpace int = -1
	movedFile     = -2
)
