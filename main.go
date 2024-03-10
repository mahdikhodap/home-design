package main

import (
	"flag"
	"fmt"
	"home/config"
	"home/requester"
	"net/http"
	"os"
	"sync"

	"github.com/rs/zerolog"
)

type Job struct {
	URL string
}

const (
	StatusOK                  = 200
	StatusNotFound            = 404
	StatusInternalServerError = 500
	//TODO add more
)

var statusMessages = map[int]string{
	StatusOK:                  "OK",
	StatusNotFound:            "Not Found",
	StatusInternalServerError: "Internal Server Error",
	//TODO add more status codes as needed
}

func getStatusMessage(code int) string {
	if msg, ok := statusMessages[code]; ok {
		return msg
	}
	return "Unknown"
}

func worker(request *requester.Requester, jobs <-chan Job, results chan<- string, wg *sync.WaitGroup) {
	for job := range jobs {
		res, err := request.Get(requester.RequestEntity{Endpoint: job.URL})
		wg.Done()
		if err != nil {
			results <- fmt.Sprintf("%s is not reachable (error: %v)", job.URL, err)
			continue
		}

		if res.StatusCode == http.StatusOK {
			results <- fmt.Sprintf("%s is reachable", job.URL)
		} else {
			statusMsg := getStatusMessage(res.StatusCode)
			results <- fmt.Sprintf("%s is not reachable (status code: %d - %s)", job.URL, res.StatusCode, statusMsg)
		}
	}
}

var (
	configFile     string
	verbosityLevel int
)

func showHelp() {
	fmt.Println("      -c {config file}")
	fmt.Println("      -v {0-10} (verbosity level, default 0)")
	fmt.Println("      -h (show help info)")
}

func parse() ([]string, bool) {

	flag.StringVar(&configFile, "c", "home.yaml", "config file")
	flag.IntVar(&verbosityLevel, "v", -1, "verbosity level, higher value - more logs")
	help := flag.Bool("h", false, "help info")
	flag.Parse()

	if err := config.ReadConf(configFile); err != nil {
		fmt.Println(err)
		return nil, false
	}
	if *help {
		return nil, false
	}
	if len(os.Args) < 1 {
		fmt.Println("Usage: go run url_checker.go <url1> ...")
		os.Exit(1)
	}
	return os.Args[1:], true
}
func main() {

	urls, pareComplete := parse()
	if !pareComplete {
		showHelp()
		os.Exit(-1)
	}

	zerolog.TimeFieldFormat = "2006-01-02T15:04:05-0700"

	if verbosityLevel < 0 {
		verbosityLevel = config.Conf.Log.Level
	}

	zerolog.SetGlobalLevel(zerolog.Level(verbosityLevel))

	requester := requester.NewRequester().Load()

	numWorkers := config.Conf.WorkerPool.Number

	jobs := make(chan Job, len(urls))
	results := make(chan string, len(urls))
	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		go worker(requester, jobs, results, &wg)
	}

	wg.Add(len(urls))
	for _, url := range urls {
		jobs <- Job{URL: url}
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
