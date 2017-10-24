package external

import (
	"fmt"
	"os"
)

const (
	fanIn = 128
)

func Sort(fileName string, blockSize int) {
	sp := newSplitter()
	sp.split(fileName, blockSize)

	blockNames := sp.blockNames
	mergeNames := []string{}

	var (
		cycleCount int
		mergeCount int
	)

	for {
		for len(blockNames) > 0 {
			var (
				blocks []*os.File
				offset int
			)
			if len(blockNames) < fanIn {
				offset = len(blockNames)
			} else {
				offset = fanIn
			}

			for _, blockName := range blockNames[:offset] {
				f, err := os.Open(blockName)
				if err != nil {
					panic(err)
				}

				blocks = append(blocks, f)
			}

			mergedName := fmt.Sprintf("merged_%d_%d", cycleCount, mergeCount)
			w, err := os.Create(mergedName)
			if err != nil {
				panic(err)
			}

			merge(blocks, w)
			mergeNames = append(mergeNames, mergedName)
			mergeCount++
			blockNames = blockNames[offset:]
		}

		if len(mergeNames) < 2 {
			outName := fmt.Sprintf("sorted.%s", fileName)
			os.Rename(mergeNames[0], outName)
			return
		}

		blockNames = mergeNames
		mergeNames = []string{}
		cycleCount++
	}
}
