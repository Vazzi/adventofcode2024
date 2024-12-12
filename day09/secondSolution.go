package day09

func secondSolution(data string) int {
	dataSize := len(data)
	blocks := make([]int, dataSize/2+1)
	spaces := make([]int, dataSize/2+1)
	fsSize := 0
	fileId := 0

	for i := range data {
		num := int(data[i] - '0') // convert to number
		fsSize += num
		if i%2 == 0 {
			blocks[fileId] = num
			fileId++
		} else {
			spaces[fileId-1] = num
		}
	}

	fs := newFilesystem(blocks, spaces, fsSize)

	optimizeWholeFiles(fs)

	return fs.computeChecksum()
}

func optimizeWholeFiles(fs *filesystem) {
	for ri := len(fs.data) - 1; ri > 0; ri-- {
		if fs.data[ri] <= 0 {
			continue
		}

		// Get whole file
		fileId := fs.data[ri]
		fileSize := 0
		for fileId == fs.data[ri] {
			ri--
			fileSize++
		}
		ri++ // Return back to first index of file

		// Find space from left to store the file
		for li := 1; li < ri; li++ {
			if fs.data[li] == freeSpace {
				spaceSize := 0
				// Compute size of the free space
				for i := li; fs.data[i] == freeSpace; i++ {
					spaceSize++
				}

				if spaceSize >= fileSize {
					// Space founded, move file
					for i := 0; i < fileSize; i++ {
						fs.data[li+i] = fileId
						fs.data[ri+i] = movedFile
					}
					break
				}
			}
		}
	}
}
