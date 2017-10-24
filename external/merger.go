package external

import (
	"bufio"
	"io"
	"os"
)

func merge(files []*os.File, outputFile *os.File) {
	var (
		readers   []*bufio.Reader
		stackTops []string
	)

	for _, file := range files {
		readers = append(readers, bufio.NewReader(file))
	}

	for _, reader := range readers {
		line, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}

		stackTops = append(stackTops, string(line))
	}

	for len(stackTops) > 0 {
		minIdx, minElem := getMin(stackTops)

		_, err := outputFile.Write([]byte(minElem + "\n"))
		if err != nil {
			panic(err)
		}

		nextElem, _, err := readers[minIdx].ReadLine()
		if err != io.EOF && err != nil {
			panic(err)
		}

		if len(nextElem) > 0 {
			stackTops[minIdx] = string(nextElem)
		} else {
			files[minIdx].Close()
			os.Remove(files[minIdx].Name())

			files = append(files[:minIdx], files[minIdx+1:]...)
			readers = append(readers[:minIdx], readers[minIdx+1:]...)
			stackTops = append(stackTops[:minIdx], stackTops[minIdx+1:]...)
		}
	}
}

func getMin(elems []string) (minIdx int, minElem string) {
	minIdx, minElem = 0, elems[0]

	for idx, elem := range elems {
		if elem < minElem {
			minIdx, minElem = idx, elem
		}
	}

	return
}
