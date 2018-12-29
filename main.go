package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/kyokomi/emoji.v1"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"net/url"
	"os"
	"os/user"
	"strings"
	"syscall"
)

const (
	AppVersion   = "0.0.1"
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

var (
	argVersion    = flag.Bool("version", false, "バージョンを出力.")
	argRepository = flag.String("repo", "", "Git Repository URL を指定.")
	argDirectory  = flag.String("dir", "", "Git Repository を展開するディレクトリを指定.")
	argUsername   = flag.String("username", "", "Git Repository の Username を指定.")
	argBranch     = flag.String("branch", "", "Git Repository のブランチ名を指定.")
)

func gitClone(repo string, repo_username string, repo_directory string) *git.Repository {
	var repo_url string
	var clone_options *git.CloneOptions

	directory := repo_directory

	if strings.Contains(repo, "backlog") && repo_username != "" {
		fmt.Println("Please Input Your Backlog User Password:")
		password, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Printf("\x1b[31;1mError : %s\x1b[0m\n", err)
			os.Exit(1)
		}
		u, _ := url.Parse(repo)
		username_password := repo_username + ":" + string(password)
		repo_url = u.Scheme + "://" + username_password + "@" + u.Host + u.Path
		clone_options = &git.CloneOptions{
			URL:      repo_url,
			Progress: os.Stdout,
		}
	} else {
		repo_url = repo
		currentUser, err := user.Current()
		if err != nil {
			fmt.Printf("\x1b[31;1mError : %s\x1b[0m\n", err)
			os.Exit(1)
		}
		sshAuth, err := ssh.NewPublicKeysFromFile("git", currentUser.HomeDir+"/.ssh/id_rsa", "")
		if err != nil {
			fmt.Printf("\x1b[31;1mError : %s\x1b[0m\n", err)
			os.Exit(1)
		}
		clone_options = &git.CloneOptions{
			URL:      repo_url,
			Progress: os.Stdout,
			Auth:     sshAuth,
		}
	}

	emoji.Printf(":white_check_mark: 対象リポジトリ %s を %s に clone します.\n", repo, directory)
	r, err := git.PlainClone(directory, false, clone_options)

	if err != nil {
		fmt.Println(err)
		emoji.Println(":bangbang: 対象リポジトリの clone に失敗しました.")
		os.Exit(1)
	} else {
		emoji.Println(":white_check_mark: 対象リポジトリの clone に成功しました.")
	}

	return r
}

func gitCheckOut(r *git.Repository, branch string) {
	remote_branch_reference := "refs/remotes/origin/" + branch
	ref, err := r.Reference(plumbing.ReferenceName(remote_branch_reference), true)
	if err != nil {
		fmt.Printf("\x1b[31;1mError : %s\x1b[0m\n", err)
		os.Exit(1)
	}

	w, err := r.Worktree()
	if err != nil {
		fmt.Printf("\x1b[31;1mError : %s\x1b[0m\n", err)
		os.Exit(1)
	}

	local_branch_reference := "refs/heads/" + branch
	checkout_options := &git.CheckoutOptions{
		Hash:   plumbing.Hash(ref.Hash()),
		Branch: plumbing.ReferenceName(local_branch_reference),
		Create: true,
		Force:  true,
	}
	err = w.Checkout(checkout_options)
	if err != nil {
		emoji.Println(":bangbang: 対象ブランチの checkout に失敗しました.")
		os.Exit(1)
	} else {
		emoji.Println(":white_check_mark: 対象ブランチの checkout に成功しました.")
	}
}

func main() {
	flag.Parse()

	if *argVersion {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	if *argRepository != "" && *argDirectory != "" {
		r := gitClone(*argRepository, *argUsername, *argDirectory)
		if *argBranch != "" {
			gitCheckOut(r, *argBranch)
		}
	} else {
		emoji.Println(":bangbang: Git リポジトリと保存先のディレクトリを指定して下さい.")
		os.Exit(1)
	}
}
