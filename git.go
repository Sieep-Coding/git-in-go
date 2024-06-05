package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Repository struct {
	RootDir string
}

type Commit struct {
	Hash       string
	Message    string
	Author     string
	Timestampe time.Time
}

func (repo *Repository) Init() {
	gitDir := filepath.Join(repo.RootDir, ".git")
	err := os.MkdirAll(gitDir, 0755)
	if err != nil {
		fmt.Println("Failed to initialize repository", err)
		os.Exit(0)
	}
	fmt.Println("Initialized repository in:", gitDir)
}

func (repo )
