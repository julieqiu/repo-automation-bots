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
	"os"
	"os/exec"

	"github.com/julieqiu/derrors"
)

func main() {
	if err := run("secret", "appID", "appInstallationID"); err != nil {
		log.Fatal(err)
	}
}

func run(secret, appID, installationID string) error {
	jwt, err := generateJWT(secret, appID)
	if err != nil {
		return err
	}
	tokens, err := generateAccessTokens(jwt, installationID)
	if err != nil {
		return err
	}
	fmt.Println(tokens)
	return nil

}

func generateJWT(secret, appID string) (_ string, err error) {
	defer derrors.Wrap(&err, "generateJWT")
	cmd := exec.Command(
		"jwt", "encode", "--secret", secret, "--iss", appID, "--exp", "+10 min", "--alg", "RS256",
	)
	cmd.Stderr = os.Stderr
	jwt, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(jwt), nil
}

func generateAccessTokens(jwt, installationID string) (_ string, err error) {
	defer derrors.Wrap(&err, "generateAccessTokens")
	cmd := exec.Command(
		"curl", "-X", "POST",
		"-H", fmt.Sprintf("Authorization: Bearer %s", jwt),
		"-H", "Accept: application/vnd.github.v3+json",
		fmt.Sprintf("https://api.github.com/app/installations/%s/access_tokens", installationID),
	)
	accessTokens, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(accessTokens), nil
}
