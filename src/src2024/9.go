package src2024

import (
	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

type disk []file

type file struct {
	id     int // -1 for empty space
	blocks int
}

func (ds file) isEmpty() bool {
	return ds.id == -1
}

type fileBlock struct {
	fileId int
}

func (db fileBlock) isEmpty() bool {
	return db.fileId == -1
}

func constructDisk(line string) disk {
	disk := make(disk, 0)
	id := 0

	for i := 0; i < len(line); i++ {
		blocks := utils.ToInt(string(line[i]))
		if blocks == 0 {
			continue
		}

		fileId := -1 // represents empty space
		if i%2 == 0 {
			fileId = id
			id++
		}

		disk = append(disk, file{
			id:     fileId,
			blocks: blocks,
		})
	}
	logger.Debugf("Line: %s", line)
	logger.Infof("disk: %v", disk)

	return disk
}

func filesToBlocks(disk disk) []fileBlock {
	blocks := make([]fileBlock, 0)
	for _, comp := range disk {
		for _ = range comp.blocks {
			blocks = append(blocks, fileBlock{
				fileId: comp.id,
			})
		}
	}
	return blocks
}

func compactFiles(d disk) disk {

	// find the last file id
	lastFileId := -1
	for j := len(d) - 1; j >= 0; j-- {
		if !d[j].isEmpty() {
			lastFileId = d[j].id
			break
		}
	}

	handleFileSwaps := func(d disk, fileId int) disk {
		for j := len(d) - 1; j >= 0; j-- {
			if d[j].isEmpty() || d[j].id != fileId {
				continue
			}

			for i := 0; i < j; i++ {
				if d[i].isEmpty() && d[i].blocks == d[j].blocks {
					// replace the file and the empty space

					d[i] = d[j]
					d[j].id = -1
					break
				} else if d[i].isEmpty() && d[i].blocks > d[j].blocks {
					remainingSpaces := file{
						id:     -1,
						blocks: d[i].blocks - d[j].blocks,
					}
					// replace the empty space with the file ...
					d[i] = d[j]
					d[j].id = -1

					// ...but insert an empty space after the moved file with the remaining empty spaces
					d = append(d[:i+1], append([]file{remainingSpaces}, d[i+1:]...)...)
					break
				} else {
					continue
				}
			}
		}
		return d
	}

	// try to swap every file
	// O(N^2) but eh..
	for id := lastFileId; id >= 0; id-- {
		d = handleFileSwaps(d, id)
	}

	return d
}

func compactFileBlocks(d []fileBlock) []fileBlock {

	i, j := 0, len(d)-1
	for true {
		logger.Debugf("i=%d j=%d", i, j)

		if i == j {
			break
		}

		// find the next available spot to insert (starting from the beginning)
		// along with the next available position to be moved (starting from the end)
		if !d[i].isEmpty() {
			i++
			continue
		} else {
			// find the next available file block to insert
			if !d[j].isEmpty() {
				d[i] = d[j]
				d[j].fileId = -1
			}
			j--
		}
	}

	return d
}

func findDiskChecksum(d []fileBlock) int {
	checkSum := 0

	for i := 0; i < len(d); i++ {
		if d[i].isEmpty() {
			continue
		}
		logger.Debugf("%v %v", i, d[i].fileId)
		checkSum += i * d[i].fileId
	}
	return checkSum
}

func (s *Solver) Day_09(part int, reader utils.Reader) int {

	var disk disk
	for line := range reader.Stream() {
		disk = constructDisk(line)
	}

	switch part {
	case 1:
		fileBlocks := filesToBlocks(disk)
		logger.Infof("fileBlocks: %v", fileBlocks)

		compactedDisk := compactFileBlocks(fileBlocks)
		logger.Infof("Compacted disk: %v", compactedDisk)

		return findDiskChecksum(compactedDisk)

	case 2:
		compacteddisk := compactFiles(disk)
		logger.Debugf("Compacted disk: %v", compacteddisk)

		logger.Debugf("Compacted disk: %v", filesToBlocks(compacteddisk))

		return findDiskChecksum(filesToBlocks(compacteddisk))
	}
	return -1

}
