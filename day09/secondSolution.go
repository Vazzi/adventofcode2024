package day09

func secondSolution(data string) int {
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

	optimizeWholeFiles(fs)

	return fs.computeChecksum()
}

func optimizeWholeFiles(fs *filesystem) {
	li := 0
	movedSomething := false

	for {
		for li < fs.size-1 {
			if fs.data[li] != freeSpace {
				li++
				continue
			}
			spaceSize := fs.readSpaceSize(li)
			ridIndex := fs.findRightFileWithMaxSize(spaceSize, li)
			if ridIndex != -1 {
				movingFileId := fs.data[ridIndex]
				movedSomething = true
				fs.markMovedFile(movingFileId)

				for i := li; fs.data[ridIndex] == movingFileId; ridIndex++ {
					fs.data[i] = fs.data[ridIndex]
					i++
					fs.data[ridIndex] = movedFile

					if ridIndex+1 >= fs.size {
						break
					}
				}
			}

			li += spaceSize
		}
		if movedSomething == false {
			break
		}
		movedSomething = false
		li = 0
	}

}

func (fs filesystem) readSpaceSize(index int) int {
	size := 0
	for i := index; i < fs.size && fs.data[i] == freeSpace; i++ {
		size++
	}
	return size
}

func (fs *filesystem) findRightFileWithMaxSize(size, lowestI int) int {
	fileSize := 0
	id := -1
	i := fs.size - 1

	for i > lowestI {
		curr := fs.data[i]

		if curr <= 0 || fs.movedFiles[curr] {
			if id != -1 && fileSize <= size {
				return i + 1
			}
			id = -1
			fileSize = 0
			i--
			continue
		}
		if id == -1 {
			id = curr
		} else if id != curr {
			if fileSize <= size {
				return i + 1
			}
			id = curr
			fileSize = 0
		}
		i--
		fileSize++
	}

	return -1
}
