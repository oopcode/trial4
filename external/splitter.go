package external

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type splitter struct {
	blocks     []*os.File
	blockNames []string
}

func newSplitter() *splitter {
	return &splitter{}
}

func (s *splitter) split(inputFileName string, blockSize int) {
	f, err := os.Open(inputFileName)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var blockNum int

LOOP:
	for {
		block := []string{}

		for i := 0; i < blockSize; i++ {
			line, _, err := reader.ReadLine()

			if err == io.EOF {
				s.writeBlock(block, blockNum)
				break LOOP
			}
			if err != nil {
				panic(err)
			}

			block = append(block, string(line))
		}

		s.writeBlock(block, blockNum)

		blockNum++
	}
}

func (s *splitter) writeBlock(block []string, blockNumber int) {
	if len(block) < 1 {
		return
	}

	sort.Strings(block)

	var (
		data     = strings.Join(block, "\n")
		fileName = fmt.Sprintf("block_%d", blockNumber)
	)
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	_, err = f.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	f.Seek(0, 0)

	s.blocks = append(s.blocks, f)
	s.blockNames = append(s.blockNames, fileName)
}

func (s *splitter) cleanUp() {
	for _, fileName := range s.blockNames {
		os.Remove(fileName)
	}
}
