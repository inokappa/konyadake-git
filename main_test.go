package main

import (
	_ "bytes"
	_ "fmt"
	_ "os/exec"
	_ "strings"
	"testing"
)

// func TestHelpFlag(t *testing.T) {
// 	cmd := exec.Command("gom", "run", "main.go", "-help")
// 	stdout := new(bytes.Buffer)
// 	cmd.Stderr = stdout
// 	msg := "バージョンを出力."
// 	_ = cmd.Run()
// 	if !strings.Contains(stdout.String(), msg) {
// 		t.Fatal("Failed Test")
// 	}
// }

func TestRepoDirectory(t *testing.T) {
	repo := "git@github.com:inokappa/example.git"
	result := repoDirectory(repo)
	if result != "example" {
		t.Fatalf("Failed test")
	}
}

func TestIsGitUrlTrueGitAt(t *testing.T) {
	repo := "git@github.com:inokappa/example.git"
	result := isGitUrl(repo)
	if result != true {
		t.Fatalf("Failed test")
	}
}

func TestIsGitUrlTrueHttps(t *testing.T) {
	repo := "https://github.com/inokappa/example.git"
	result := isGitUrl(repo)
	if result != true {
		t.Fatalf("Failed test")
	}
}

func TestIsGitUrlFalse(t *testing.T) {
	repo := "github.com:inokappa/example.git"
	result := isGitUrl(repo)
	if result != false {
		t.Fatalf("Failed test")
	}
}
