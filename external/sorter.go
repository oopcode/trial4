package external

import (
	"fmt"
	"os"
)

func Sort(fileName string, blockSize int) {
	sp := newSplitter()
	sp.split(fileName, blockSize)

	out, err := os.Create(fmt.Sprintf("sorted.%s", fileName))
	if err != nil {
		panic(err)
	}

	merge(sp.blocks, out)
	sp.cleanUp()
}
