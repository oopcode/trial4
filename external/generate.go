package external

import (
	"math/rand"
	"os"
)

func GenerateFile(fileName string, numLines, lineLen int) {
	w, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	for i := 0; i < numLines; i++ {
		w.Write([]byte(randString(lineLen)))
	}
}

func randString(n int) string {
	var (
		letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		b           = make([]rune, n)
	)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b) + "\n"
}
