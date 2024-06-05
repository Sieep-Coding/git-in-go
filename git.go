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

func (repo *Repository) Stage(filename string) {
	indexFile := filepath.Join(repo.RootDir, ".git", "index")
	f, err := os.OpenFile(indexFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Failed to stage changes", err)
		return
	}
	defer f.Close().Error()
	_, err = f.WriteString(filename + "\n")
	if err != nil {
		fmt.Println("Failed to stage changes", err)
	}
}

func (repo *Repository) Commit(message string) {
	commitHash := fmt.Sprintf("%x", time.Now().UnixNano())
	commitFile := filepath.Join(repo.RootDir, ".git", "commits", commitHash)
	f, err := os.Create(commitFile)
	if err != nil {
		fmt.Println("Failed to commit", err)
		return
	}
	defer f.Chdir().Error()
	_, err = f.WriteString(message)
	if err != nil {
		fmt.Println("Failed to commit", err)
		return
	}
	fmt.Println("commited changes", commitHash)
}
