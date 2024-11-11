// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"strconv"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	crypto "golang.org/x/crypto/ssh"
)

const sshKeyFile = "/Users/julieqiu/.ssh/id_ed25519"

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	const (
		generateGoogleapisGen = "docker-image/generate-googleapis-gen.sh"
		commit                = ""

		googleapisURL    = "git@github.com:googleapis/googleapis"
		googleapisGenURL = "git@github.com:googleapis/googleapis-gen"

		googleapisDir         = "/tmp/generatetest/googleapis"
		googleapisGenDir      = "/tmp/generatetest/googleapis-gen"
		googleapisGenCloneDir = "/tmp/generatetest/googleapis-gen-clone"
		garbageFile           = "googleapis-gen/google/cloud/vision/v1/vision-v1-nodejs/garbage.js"
		bogusFile             = "googleapis-gen/google/bogus/api/keepme.java"
		helloFile             = "googleapis-gen/hello.txt"
	)

	if err := os.RemoveAll("/tmp/generatetest"); err != nil {
		return err
	}

	for _, repo := range [][]string{
		{googleapisURL, googleapisDir},
		{googleapisGenURL, googleapisGenDir},
		{googleapisGenURL, googleapisGenCloneDir},
	} {
		_, err := cloneRepo(repo[0], repo[1], sshKeyFile)
		if err != nil {
			return err
		}
		slog.Info("Done!")
	}

	shaGoogleapis, err := computeSHATwoAgo(googleapisDir)
	if err != nil {
		return err
	}
	err = modifyGoogleapisGen(googleapisGenDir, shaGoogleapis, garbageFile, bogusFile, helloFile)
	if err != nil {
		return err
	}
	err = runTest(googleapisGenDir)
	if err != nil {
		return err
	}
	confirmOutput(garbageFile, bogusFile, helloFile)
	return nil
}

func cloneRepo(url, dir, sshKeyFile string) (*git.Repository, error) {
	slog.Info(fmt.Sprintf("Cloning %s to %s", url, dir))
	sshKey, err := os.ReadFile(sshKeyFile)
	if err != nil {
		return nil, err
	}

	signer, err := crypto.ParsePrivateKey([]byte(sshKey))
	if err != nil {
		return nil, err
	}
	publicKey := &ssh.PublicKeys{User: "git", Signer: signer}
	repo, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:          url,
		Auth:         publicKey,
		Progress:     os.Stdout,
		SingleBranch: true,
		Depth:        5,
	})
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func computeSHATwoAgo(dir string) (string, error) {
	// # Select the sha for HEAD~2
	// sha=$(git -C googleapis log -3 --format=%H | tail -1)
	// cmd := exec.Command("git", "-C", dir, "log", "-3", "--format=%H", "|", "tail", "-1")
	return runCommandWithOutput(dir, "git", "rev-parse", "--short", "HEAD~2")
}

func modifyGoogleapisGen(dir, shaGoogleapis, garbageFile, bogusFile, helloFile string) error {
	commands := [][]string{
		// Create a git repository in the googleapis-gen directory, setting the
		// initial branch name to master
		{"git", "init", "--initial-branch=master"},

		// # Hello.txt lives in the root directory and should not be removed.
		{"echo", "hello", ">", helloFile},
		// # keepme.java fails to build and therefore should not be removed.
		{"mkdir", "-p", ">", "googleapis-gen/google/bogus/api"},
		{"echo", `"import *;"`, ">", bogusFile},

		// # garbage.js should be wiped out by newly generated code.
		{"mkdir", "-p", "googleapis-gen/google/cloud/vision/v1/vision-v1-nodejs"},
		{"echo", `"garbage"`, ">", garbageFile},

		{"git", "config", "user.email", "test@example.com"},
		{"git", "add", "user.name", "Test McTestFace"},

		{"git", "add", "-A"},
		{"git", "commit", "-m", "Hello world."},
		{"git", "tag", fmt.Sprintf("googleapis-%s", shaGoogleapis)},
		{"git", "checkout", "-b", "other"},
	}
	for _, c := range commands {
		if err := runCommand(dir, c[0], c[1:]...); err != nil {
			return err
		}
	}
	return nil
}

/*
func createGoogleapisGenClone(url, dir string) error {
	commands := [][]string{
		{"git", "config", "user.email", "test@example.com"},
		{"git", "add", "user.name", "Test McTestFace"},
	}
	for _, c := range commands {
		if err := runCommand(dir, c[0], c[1:]...); err != nil {
			return err
		}
	}
	return nil
}
*/

func runTest(dir string) error {
	cmd := exec.Command("bash", "-x", "docker-image/generate-googleapis-gen.sh")
	slog.Info(cmd.String())
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env,
		"GOOGLEAPIS_GEN=`realpath googleapis-gen-clone`",
		`INSTALL_CREDENTIALS="echo 'Pretending to install credentials...'"`,
		`BUILD_TARGETS="//google/cloud/vision/v1:vision-v1-nodejs //google/bogus:api"`,
		`FETCH_TARGETS="//google/cloud/vision/v1:vision-v1-nodejs"`,
	)

	// # Display the state of googleapis-gen
	commands := [][]string{
		{"git", "checkout", "master"},
		{"git", "log", "--name-only"},
	}
	for _, c := range commands {
		if err := runCommand(dir, c[0], c[1:]...); err != nil {
			return err
		}
	}
	return nil
}

func confirmOutput(garbageFile, bogusFile, helloFile string) error {
	// # Confirm that we added at least one commit.
	// commit_count=$(git -C googleapis-gen log --pretty=%H | wc -l)
	// if [ $commit_count -lt 2 ] ; then
	//     echo "ERROR: There should be a new commit in googlapis-gen"
	//     exit 99
	// fi

	s, err := runCommandWithOutput("git", "log", "--pretty=%H", "|", "wc", "-l")
	if err != nil {
		return err
	}
	commitCount, err := strconv.Atoi(string(s))
	if err != nil {
		return err
	}
	if commitCount < 2 {
		return fmt.Errorf("ERROR: There should be a new commit in googlapis-gen")
	}

	// # Confirm the garbage file was removed because it was replaced by new code.
	// if [ -e "$garbage_file_path" ] ; then
	//     echo "ERROR: $garbage_file_path should have been removed"
	//     exit 98
	// fi
	_, err = os.Stat(garbageFile)
	if err == nil || err != os.ErrNotExist {
		return fmt.Errorf("ERROR: $garbage_file_path should have been removed")
	}

	// # Confirm that the source code for the target that failed to build still exists.
	// if [ ! -e "$bogus_file_path" ] ; then
	//     echo "ERROR: $bogus_file_path should not been removed"
	//     exit 97
	// fi
	_, err = os.Stat(bogusFile)
	if err != nil {
		return fmt.Errorf("ERROR: $bogus_file_path should not have been removed")
	}

	// # Confirm that hello.txt still exists because it's not in a source code directory.
	// if [ ! -e "$hello_path" ] ; then
	//     echo "ERROR: $hello_path should not been removed"
	//     exit 96
	// fi
	_, err = os.Stat(helloFile)
	if err != nil {
		return fmt.Errorf("ERROR: $hello_path should not have been removed")
	}
	return nil
}

func runCommand(dir, c string, args ...string) error {
	cmd := exec.Command(c, args...)
	slog.Info(cmd.String())
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runCommandWithOutput(dir, c string, args ...string) (string, error) {
	cmd := exec.Command(c, args...)
	slog.Info(cmd.String())
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
