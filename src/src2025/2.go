package src2025

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

type repeatedBlockReq struct {
	blockSize, blocks int
}

// generateRepeated produces all numbers with repeated blocks of digits, given a block size.
// A block of size blockSize is repeated `block` times in order to generate the resulting number.
// The function generates all available blocks and then concatenates them.
// Each block is defined as a integer j in the range [10^(blockSize-1), 10^blockSize)
// Example for blockSize = 2, blocks = 2:
//
//	j: 10 - 99
//	10 -> 1010
//	23 -> 2323
//
// Example for blockSize = 3, blocks = 3:
//
//	j: 100 - 999
//	100 -> 100100100
//	234 -> 234234234
func generateRepeated(ctx context.Context, req repeatedBlockReq) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		start := int64(math.Pow10(req.blockSize - 1))
		end := int64(math.Pow10(req.blockSize))
		for j := start; j < end; j++ {
			id := strings.Repeat(fmt.Sprintf("%d", j), req.blocks)
			select {
			case <-ctx.Done():
				return
			case ch <- utils.ToInt(id):
			}
		}
	}()
	return ch
}

func (s *Solver) Day_02(part int, reader utils.Reader) int {
	idRanges := make([]utils.IdRange, 0)

	for line := range reader.Stream() {
		ranges := strings.Split(line, ",")
		for i := range ranges {
			parts := strings.Split(ranges[i], "-")
			if len(parts) != 2 {
				continue
			}

			a := utils.ToInt(parts[0])
			b := utils.ToInt(parts[1])
			idRanges = append(idRanges, utils.IdRange{
				Min:       a,
				Max:       b,
				MinLength: len(parts[0]),
				MaxLength: len(parts[1]),
			})
		}
	}

	logger.Debugf("Id ranges: %+v", idRanges)

	var invalidIdsSum int

	for _, idRange := range idRanges {

		var reqs []repeatedBlockReq

		switch part {
		case 1:
			logger.Debugf("Id range: %+v %d %d", idRange, utils.CeilingDivision(idRange.MinLength, 2), utils.CeilingDivision(idRange.MaxLength, 2))

			// - find the number of digits in min and max
			// - do ceiling division by 2 for both to find the number size that needs to be duplicated/repeated
			for i := utils.CeilingDivision(idRange.MinLength, 2); i <= utils.CeilingDivision(idRange.MaxLength, 2); i++ {
				// we want 2 blocks of i size
				// eg. received number 222233 -> ceilingDiv = 3 so we target 2 blocks of size 3
				// 4 blocks of 4/2=2 -> o
				reqs = append(reqs, repeatedBlockReq{
					blockSize: i,
					blocks:    2,
				})
			}

		case 2:
			// logger.Debugf("Id range: %+v %v %v", idRange, utils.FindFactors(idRange.MinLength), utils.FindFactors(idRange.MaxLength))

			// The factor defines the block size and the original number is used to calculate how many blocks should be generated.
			// eg. if the number is 10, we can generate:
			//  - 10 blocks of size 1
			//  - 5 blocks of size 2
			//  - 2 blocks of size 5
			// X  1 block of size 10 X (we want repeated)
			for _, f := range utils.FindFactors(idRange.MinLength) {
				if f == idRange.MinLength {
					continue
				}
				reqs = append(reqs, repeatedBlockReq{
					blockSize: f,
					blocks:    idRange.MinLength / f,
				})
			}

			// If the lengths of the two ranges are different, we have to consider all cases.
			if idRange.MinLength != idRange.MaxLength {
				for _, f := range utils.FindFactors(idRange.MaxLength) {
					if f == idRange.MaxLength {
						continue
					}
					reqs = append(reqs, repeatedBlockReq{
						blockSize: f,
						blocks:    idRange.MaxLength / f,
					})
				}
			}

		}

		logger.Debugf("Reqs: %+v", reqs)

		// Keep only the unique ids per idRange
		uniqueIds := map[int]struct{}{}
		for _, req := range reqs {

			// Generate repeated blocks of a fixed size
			ch := generateRepeated(context.Background(), req)
			for id := range ch {
				if id >= idRange.Min && id <= idRange.Max {
					logger.Debugf("(temp) invalid id: %d", id)
					uniqueIds[id] = struct{}{}
				}
			}
		}

		for id := range uniqueIds {
			logger.Debugf("Invalid id: %d", id)
			invalidIdsSum += id
		}
	}

	return invalidIdsSum
}
