## Context

The teejay release script fails under jj-colocated repos. The cloudia-reader-aws project solved the same problem by using git commands that work in both pure-git and jj-colocated contexts. The key insight: jj-colocated mode keeps the git working tree in sync, so `git diff --quiet` catches uncommitted jj changes, and `git push origin HEAD:main` works with detached HEAD.

## Goals / Non-Goals

**Goals:**
- Make the release script work for both pure-git and jj-colocated repos
- Keep the script 100% pure git (no jj commands)
- Follow the proven pattern from cloudia-reader-aws

**Non-Goals:**
- Adding jj-specific commands or detection
- Changing the release workflow or GitHub Actions

## Decisions

### 1. Clean tree check: `git diff --quiet`
Replace `git status --porcelain` with `git diff --quiet && git diff --cached --quiet`. This checks for both unstaged and staged changes, and works reliably under jj-colocated mode because jj syncs changes to the git working tree.

**Reference**: cloudia-reader-aws `release.sh` line 13.

### 2. Remove branch check
Remove the `git branch --show-current` check entirely. Under jj, HEAD is always detached so this always fails. The branch check is unnecessary — if the working tree is clean and the tag doesn't exist, the release is valid regardless of branch. The push to `HEAD:main` ensures changes land on main.

**Alternative considered**: Check branch via `git rev-parse --abbrev-ref HEAD` — rejected because it also returns `HEAD` under jj detached state.

### 3. Push with `HEAD:main`
Replace `git push origin main` with `git push origin HEAD:main`. This pushes the current commit to the remote main branch regardless of whether a local branch is checked out. Works identically under both git and jj.

**Reference**: cloudia-reader-aws `release.sh` line 63.

## Risks / Trade-offs

- **Removing branch check means releases can be made from any commit** → Acceptable. The user explicitly confirms the release, and the tag + push provide the actual safety. The branch check was a convenience, not a critical guard.
