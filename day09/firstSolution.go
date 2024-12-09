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
	optimize(fs)

	return fs.computeChecksum()
}

func optimize(fs *filesystem) {
	li := 0
	ri := fs.size - 1

	for li < ri {

		if fs.data[li] != freeSpace {
			li++
			continue
		}
		if fs.data[ri] == freeSpace {
			ri--
			continue
		}
		fs.data[li] = fs.data[ri]
		fs.data[ri] = freeSpace
		li++
		ri--
	}
}
