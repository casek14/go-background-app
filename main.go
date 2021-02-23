package main

import (
	"bufio"
	"github.com/kardianos/service"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

var logger service.Logger

type program struct {
	FileName string
	LogPath  string
}

func (p *program) run() {
	// do smthng here

	f, err := os.Create(p.FileName)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	_, err = w.WriteString("Counter started ...\n")
	w.Flush()

	var i int
	i = 0
	for {
		str := "Sleeping for " + strconv.Itoa(i) + " seconds .\n"
		_, err = w.WriteString(str)
		w.Flush()
		time.Sleep(time.Second * time.Duration(i))
		i++
	}

	_, err = w.WriteString("WORK DONE")
	w.Flush()

	return
}

func (p *program) Start(s service.Service) error {
	logFileName := "go-service-example-simple.log"
	pathSeparator := "/"
	workDir := "/var/log"
	if runtime.GOOS == "windows" {
		pathSeparator = "\\"
		workDir = `C:\Program Files\Go-service-example`
		if _, err := os.Stat(workDir); os.IsNotExist(err) {
			err = os.MkdirAll(workDir, 0755)
			if err != nil {
				log.Fatalf("Error creating working dir: %s", err.Error())
			}
		}
	}
	fileName := workDir + pathSeparator + logFileName
	p.FileName = filepath.FromSlash(fileName)
	log.Printf("<=== outputing to %s ...\n", p.FileName)
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	var cmd string
	args := os.Args
	if len(args) > 1 {
		cmd = args[1]
	}
	svcConfig := &service.Config{
		Name:        "GoServiceExampleSimple",
		DisplayName: "Go Service Example",
		Description: "This is an example Go service.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	switch cmd {
	case "install":
		install(s)
	case "start":
		start(s)
	default:
		logger, err = s.Logger(nil)
		if err != nil {
			log.Fatal(err)
		}
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
	}
}

func install(s service.Service) {
	log.Println("Installing ...")
	err := s.Install()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Installed !")
}

func start(s service.Service) {
	log.Println("Starting ...")
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Started !")
}
