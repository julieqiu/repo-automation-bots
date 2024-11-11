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
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

var repo = flag.String("repo", "", "")

func main() {
	flag.Parse()
	if err := run(*repo); err != nil {
		log.Fatal(err)
	}
}

func run(repoDir string) error {
	cmd := exec.Command("bazelisk", "query", `filter("-(go|csharp|java|php|ruby|nodejs|py)$", kind("rule", //...:*))`)
	cmd.Dir = repoDir
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	targets := strings.Fields(string(output))

	// Confirm that bazel can fetch remote build dependencies before building
	// with -k.  Otherwise, we can't distinguish a build failure due to a bad proto
	// vs. a build failure due to transient network issue.
	/*
		if err := runCommand(repoDir, "bazelisk", append([]string{"fetch"}, targets...)...); err != nil {
			return err
		}
	*/

	for _, target := range targets {
		if strings.Contains(target, "csharp") || strings.Contains(target, "ruby") {
			continue
		}
		if strings.Contains(target, "secretmanager/v1:gapi-cloud-secretmanager-v1-go") {
			// Invoke bazel build. Limiting job count helps to avoid memory error b/376777535.
			/*
				if err := runCommand(repoDir, "bazelisk", "build", "--jobs=8", "-k", target); err != nil {
					return err
				}
			*/
			parts := strings.SplitN(target, ":", 2)
			parts[0] = strings.TrimPrefix(parts[0], "//")
			tarFile := fmt.Sprintf("%s/bazel-bin/%s/%s.tar.gz", repoDir, parts[0], parts[1])
			if err := runCommand(".", "tar", "-xf", tarFile); err != nil {
				return err
			}
			if err := runCommand(".", "go", "run", "./postprocessor",
				"--client-root", "cloud.google.com/go",
				"--googleapis-dir", repoDir,
				"--branch", "julie-pr1",
				"--dirs", "cloud.google.com/go/secretmanager",
				"--pr-file", "prfile.txt",
			); err != nil {
				return err
			}
		}
	}
	return nil
}

func runCommand(dir, c string, args ...string) error {
	cmd := exec.Command(c, args...)
	cmd.Dir = dir
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	slog.Info(strings.Repeat("-", 80))
	slog.Info(cmd.String())
	return cmd.Run()
}
