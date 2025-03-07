package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type arrayFlag []string

func (i *arrayFlag) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *arrayFlag) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type config struct {
	ext     arrayFlag
	size    int64
	list    bool
	del     bool
	wLog    io.Writer
	archive string
}

func main() {
	var ext arrayFlag

	root := flag.String("root", ".", "Root directory to start walking in")
	list := flag.Bool("list", false, "List files only")
	del := flag.Bool("del", false, "Delete matched files")
	flag.Var(&ext, "", "File extensions to filter for")
	size := flag.Int64("size", 0, "Minimum file size")
	logFile := flag.String("log", "", "Log deletions to this file")
	archive := flag.String("archive", "", "Archive directory")
	flag.Parse()

	var (
		f   = os.Stdout
		err error
	)

	if *logFile != "" {
		f, err = os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		defer f.Close()
	}

	c := config{
		ext:     ext,
		size:    *size,
		list:    *list,
		del:     *del,
		wLog:    f,
		archive: *archive,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(root string, out io.Writer, cfg config) error {
	delLogger := log.New(cfg.wLog, "Deleted File: ", log.LstdFlags)

	return filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			for _, e := range cfg.ext {
				if filterOut(path, e, cfg.size, info) {
					return nil
				}
			}

			if cfg.list {
				return listFile(path, out)
			}

			if cfg.archive != "" {
				if err := archiveFile(cfg.archive, root, path); err != nil {
					return err
				}
			}

			if cfg.del {
				return delFile(path, delLogger)
			}

			return listFile(path, out)
		})
}
