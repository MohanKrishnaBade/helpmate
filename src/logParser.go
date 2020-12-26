package src

import (
	"bufio"
	"fmt"
	"github.com/alecthomas/chroma/quick"
	"github.com/helpmate/util"
	"github.com/mattn/go-colorable"
	"github.com/stoicperlman/fls"
	"io"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var mutex = &sync.Mutex{}

func LogParser(filepath, query string, limit int64, show bool) {
	file := util.HandleErrors(os.Open(filepath)).(*os.File)
	defer file.Close()
	channel, count := make(chan string), make(chan int64, 1) // unbuffered and buffered channel
	util.HandleErrors(file.Seek(getPosByLimit(file, limit), io.SeekStart))
	bf := bufio.NewReader(file)

	// iterate over the file line by line and parse it.
	var i int64
	for ; ; i++ {
		wg.Add(1)
		line, err := bf.ReadString('\n')
		if err == io.EOF || err != nil || (limit > 0 && limit == i) {
			wg.Done()
			break
		}
		go func(line string, ch chan string, wg *sync.WaitGroup) {
			lineParser(line, query, ch)
			defer wg.Done()
		}(line, channel, &wg)
	}

	// readout from the channel and stdout the output.
	go func(ch chan string, count chan int64, limit int64) {
		var wg1 sync.WaitGroup
		var noOfIterations int64
		for s := range ch {
			if show {
				wg1.Add(1)
				go func(s string, wg1 *sync.WaitGroup) {
					defer wg1.Done()
					printer(s)
				}(s, &wg1)
			}
			noOfIterations += 1
		}
		wg1.Wait()
		count <- noOfIterations

	}(channel, count, limit)

	// wait until all goroutines finished their job.
	wg.Wait()
	close(channel)
	printer(fmt.Sprintf("[%d] => results found!!!", <-count))
	close(count)

}

func getPosByLimit(file *os.File, limit int64) int64 {
	whence := io.SeekStart
	if limit < 0 {
		whence = io.SeekEnd
	} else {
		limit = 0
	}
	return util.HandleErrors(fls.LineFile(file).SeekLine(limit, whence)).(int64)
}

func printer(str string) {
	// get a cross-platform bash escape sequences emulator
	stdout := colorable.NewColorableStdout()
	// colorize and print the documentation
	mutex.Lock() // lock io before using it.
	//re := regexp.MustCompile(`\{(?:[^{}]|(R))*\}`)
	if err := quick.Highlight(stdout, strings.ReplaceAll(str, "\\", ""), "go", "terminal256", "dracula"); err != nil {
		panic(err)
	}
	mutex.Unlock() // release io after done.
}
func lineParser(str string, searchKey string, ch chan<- string) {
	if strings.TrimSpace(str) != "" {
		if strings.Contains(str, searchKey) {
			ch <- str
		}
	}
}
