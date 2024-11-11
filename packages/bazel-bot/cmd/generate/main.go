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
	"bufio"
	"flag"
	"log/slog"
	"os"
	"os/exec"
)

var (
	tagsFile = flag.String("tags.txt", "", "")
	shas     = flag.String("shas.txt", "", "")
)

func main() {
}

func readTags(f string) (map[string]bool, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	output := map[string]bool{}
	for scanner.Scan() {
		output[scanner.Text()] = true
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return output, nil
}

func run(shas []string, tags []string) {
	var ungeneratedSHAs []string
	for _, sha := range shas {
	}
	// for sha in $shas; do
	//
	//	if grep $sha tags.txt; then
	//	    # Found a sha we already generated.
	//	    break
	//	else
	//	    ungenerated_shas+=($sha)
	//	fi
	//
	// done
	//
}

func rangeUngeneratedSHAs() {
	// # Iterate over the ungenerated_shas from oldest to newest.
	// for (( idx=${#ungenerated_shas[@]}-1 ; idx>=0 ; idx-- )) ; do
	//
	//	sha="${ungenerated_shas[idx]}"
	//
	//	# Collect the commit message from the commit in googleapis.  The commit
	//	# messages from broken shas will accumulate in commit-msg.txt.
	//	git -C "$GOOGLEAPIS" log -1 --format=%B $sha >> commit-msg.txt
	//
	//	# If as $sha is contained in a list of bad SHAs (SHAs that
	//	# will cause bazel to fail) skip the sha. The variable $BROKEN_SHAS
	//	# is defined in the Cloud Build Terraform config, with the intention that it
	//	# is only used for exceptional circumstances.
	//	if echo $BROKEN_SHAS | grep $sha; then
	//	    echo "skipping broken $sha"
	//	    continue
	//	fi
	//
	//	git -C "$GOOGLEAPIS" checkout "$sha"
	//	# Choose build targets.
	//	if [[ -z "$BUILD_TARGETS" ]] ; then
	//	    targets=$(cd "$GOOGLEAPIS" \
	//	    && bazelisk query $BAZEL_FLAGS  'filter("-(go|csharp|java|php|ruby|nodejs|py)$", kind("rule", //...:*))' \
	//	    | grep -v -E ":(proto|grpc|gapic)-.*-java$")
	//	else
	//	    targets="$BUILD_TARGETS"
	//	fi
	//	# Clean out all the source packages from the previous build.
	//
	//	BAZEL_BIN=$(cd "$GOOGLEAPIS" && bazel info bazel-bin)
	//	# The directory might not exist before the build, but its path is known
	//	if test -d "$BAZEL_BIN"; then
	//	  rm -f $(find "$BAZEL_BIN" -name "*.tar.gz")
	//	fi
	//	# Confirm that bazel can fetch remote build dependencies before building
	//	# with -k.  Otherwise, we can't distinguish a build failure due to a bad proto
	//	# vs. a build failure due to transient network issue.
	//	if [[ -z "$FETCH_TARGETS" ]] ; then
	//	    fetch_targets="$targets"
	//	else
	//	    fetch_targets="$FETCH_TARGETS"
	//	fi
	//	(cd "$GOOGLEAPIS" && bazelisk fetch $BAZEL_FLAGS $fetch_targets)
	//	# Some API always fails to build.  One failing API should not prevent all other
	//	# APIs from being updated.
	//	set +e
	//	# Invoke bazel build. Limiting job count helps to avoid memory error b/376777535.
	//	(cd "$GOOGLEAPIS" && bazelisk build --jobs=8 $BAZEL_FLAGS -k $targets)
	//
	//	# Clear out the existing contents of googleapis-gen before we copy back into it,
	//	# so that deleted APIs will be be removed.
	//	rm -rf "$GOOGLEAPIS_GEN/external" "$GOOGLEAPIS_GEN/google" "$GOOGLEAPIS_GEN/grafeas"
	//
	//	# Untar the generated source files into googleapis-gen.
	//	let target_count=0
	//	failed_targets=()
	//	for target in $targets ; do
	//	    let target_count++
	//	    tar_gz=$(echo "${target:2}.tar.gz" | tr ":" "/")
	//	    # Create the parent directory if it doesn't already exist.
	//	    parent_dir=$(dirname $tar_gz)
	//	    target_dir="$GOOGLEAPIS_GEN/$parent_dir"
	//	    mkdir -p "$target_dir"
	//	    tar -xf "$BAZEL_BIN/$tar_gz" -C "$target_dir" || {
	//	        failed_targets+=($target)
	//	        # Restore the original source code because bazel failed to generate
	//	        # the new source code.
	//	        git -C "$GOOGLEAPIS_GEN" checkout -- "$target_dir"
	//	        # TODO: report an issue with 'gh'
	//	    }
	//	done
	//
	//	# Report failures.
	//	let failed_percent="100 * ${#failed_targets[@]} / $target_count"
	//	set -e
	//	echo "$failed_percent% of targets failed to build."
	//	printf '%s\n' "${failed_targets[@]}"
	//
	//	# Tell git about the new source code we just copied into googleapis-gen.
	//	git -C "$GOOGLEAPIS_GEN" add -A
	//
	//	# Credentials only last 10 minutes, so install them right before git pushing.
	//	$INSTALL_CREDENTIALS
	//
	//	if git -C "$GOOGLEAPIS_GEN" diff-index --quiet HEAD ; then
	//	    # No changes to commit, so just push the tag.
	//	    git -C "$GOOGLEAPIS_GEN" tag "googleapis-$sha"
	//	    git -C "$GOOGLEAPIS_GEN" push origin "googleapis-$sha"
	//	else
	//	    # Determine the current branch so we can explicitly push to it
	//	    googleapis_gen_branch=$(git -C "$GOOGLEAPIS_GEN" rev-parse --abbrev-ref HEAD)
	//
	//	    echo "Source-Link: https://github.com/googleapis/googleapis/commit/$sha" >> commit-msg.txt
	//	    # Commit changes and push them.
	//	    git -C "$GOOGLEAPIS_GEN" commit -F "$(realpath commit-msg.txt)"
	//	    git -C "$GOOGLEAPIS_GEN" tag "googleapis-$sha"
	//	    git -C "$GOOGLEAPIS_GEN" push origin "$googleapis_gen_branch" "googleapis-$sha"
	//	fi
	//	# Clean out the commit message so it's not used for the next sha.
	//	rm commit-msg.txt
	//
	// done
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
