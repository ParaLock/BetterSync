package synchronizers

import (
    "fmt"
    "os"
    "time"
	git "github.com/go-git/go-git/v5"
	// "github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
    "github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

type Git struct {
    test int
}

func (m *Git) Execute() {
    
	repoURL := "git@github.com:ParaLock/SyncTest.git"
	directory := "/home/linuxdev/SyncTest/"
	email := "nathanael.mercaldo@gmail.com"
    username := "Paralock"
    sshKeyPath := os.Getenv("HOME") + "/.ssh/id_rsa_personal_github"

    remotes := make(map[string]string)
	remotes["origin"] = "git@github.com:ParaLock/SyncTest.git"
    
    
    var err = executeSync(
        repoURL,
        directory,
        email,
        username,
        sshKeyPath,
        remotes)

    fmt.Println("test123")
    fmt.Println(err)
}

func executeSync(
    repoURL string,
    directory string,
    email string,
    username string,
    sshKeyPath string,
	remotes map[string]string) error {

	var repo *git.Repository
	var err error

    sshKey, err := ssh.NewPublicKeysFromFile("git", sshKeyPath, "")
	if err != nil {
		return fmt.Errorf("could not create ssh auth: %w", err)
	}

	if _, err := os.Stat(directory + "/.git"); os.IsNotExist(err) {

		repo, err = git.PlainClone(directory, false, &git.CloneOptions{
			URL:           repoURL,
			Auth:          sshKey,
		})
		if err != nil {
			return err
		}
	} else {

		repo, err = git.PlainOpen(directory)
		if err != nil {
			return err
		}
	}

	w, err := repo.Worktree()
	if err != nil {
		return err
	}

    // _, err = w.Stash(git.StashOptions{
    //     KeepIndex: true,
    //     Message:   "Pre-pull stash on " + time.Now().Format(time.RFC1123),
    // })
    // if err != nil {
    //     return fmt.Errorf("could not stash changes: %w", err)
    // }

	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
		Auth: sshKey,
	})

	if err != nil && err != git.NoErrAlreadyUpToDate {
		return err
	}

    // err = w.StashPop(0)
    // if err != nil {
    //     return fmt.Errorf("could not pop stash: %w", err)
    // }

	err = w.AddWithOptions(&git.AddOptions{All: true})
	if err != nil {
		return err
	}

	_, err = w.Commit("Sync Commit on "+time.Now().Format(time.RFC1123), &git.CommitOptions{
		All: true,
		Author: &object.Signature{
			Name:  username,
			Email: email,
			When:  time.Now(),
		},
	})
	if err != nil {
		return err
	}

	err = repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: sshKey,
	})
	if err != nil {
		return err
	}

	return nil
}