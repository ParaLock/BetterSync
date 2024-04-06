package synchronizers

import (
    "fmt"
    "os"
    "time"
	"os/exec"
	"strings"
)

type GitCLI struct {
    test int
}

func (m *GitCLI) Execute() {

	repoDir := "/home/linuxdev/SyncTest/"
	email := "nathanael.mercaldo@gmail.com"
    username := "Paralock"
    sshKeyPath := os.Getenv("HOME") + "/.ssh/id_rsa_personal_github"
	remotes := make(map[string]string)
	remotes["origin"] = "git@github.com:ParaLock/SyncTest.git"
    
    var err = runSync(
        repoDir,
        email,
        username,
        sshKeyPath,
		remotes)

    fmt.Println(err)
}

func remoteExists(remoteName string) bool {
	cmd := exec.Command("git", "remote", "get-url", remoteName)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func workingDirectoryClean() bool {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error checking working directory:", err)
		return false
	}
	return strings.TrimSpace(string(output)) == ""
}

func localIsNotAheadOfRemote() bool {
	cmd := exec.Command("git", "rev-list", "origin/main..HEAD")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error checking if local is ahead of remote:", err)
		return false
	}
	return strings.TrimSpace(string(output)) == ""
}


func runSync(
	repoDir string,
	email string,
	username string,
	sshKeyPath string,
	remotes map[string]string) error {

	sshCommand := fmt.Sprintf("ssh -i %s -o IdentitiesOnly=yes", sshKeyPath)
	os.Setenv("GIT_SSH_COMMAND", sshCommand)

	err := os.Chdir(repoDir)
    if err != nil {
        fmt.Println("Failed to change directory:", err)
        return nil
    }

	if _, err := exec.Command("test", "-d", repoDir).Output(); err != nil {
		fmt.Println("Cloning repository...")
		if out, err := exec.Command("git", "clone", remotes["origin"]).CombinedOutput(); err != nil {
			return fmt.Errorf("failed to clone repository: %s, %v", string(out), err)
		}
	}

	if out, err := exec.Command("git", "-C", repoDir, "config", "user.name", username).CombinedOutput(); err != nil {
		return fmt.Errorf("failed to set git username: %s, %v", string(out), err)
	}
	if out, err := exec.Command("git", "-C", repoDir, "config", "user.email", email).CombinedOutput(); err != nil {
		return fmt.Errorf("failed to set git email: %s, %v", string(out), err)
	}

	for remoteName, remoteUrl := range remotes {

		if(!remoteExists(remoteName)) {
			if out, err := exec.Command("git", "remote", "add", remoteName, remoteUrl).CombinedOutput(); err != nil {
				fmt.Printf("Remote %s might already exist or error: %s\n", remoteName, string(out))
			}
		}

	}

	changesStashed := false

	if(!workingDirectoryClean() && !localIsNotAheadOfRemote()) {

		if out, err := exec.Command("git", "stash").CombinedOutput(); err != nil {
			return fmt.Errorf("failed to stash changes: %s, %v", string(out), err)
		}
		changesStashed = true
	}

	if out, err := exec.Command("git", "pull", "origin").CombinedOutput(); err != nil {
		return fmt.Errorf("failed to pull from origin: %s, %v", string(out), err)
	}


	if(changesStashed) {
		if out, err := exec.Command("git", "stash", "pop").CombinedOutput(); err != nil {
			fmt.Printf("Error or merge conflict while unstashing: %s\n", string(out))
			return fmt.Errorf("merge conflict or error while unstashing: %v", err)
		}
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	commitMessage := fmt.Sprintf("Sync Commit on %s", currentTime)
	if out, err := exec.Command("git", "commit", "-am", commitMessage).CombinedOutput(); err != nil {
		fmt.Printf("Nothing to commit or error: %s\n", string(out))
	}

	if out, err := exec.Command("git", "push", "origin").CombinedOutput(); err != nil {
		return fmt.Errorf("failed to push to origin: %s, %v", string(out), err)
	}

	return nil
}
