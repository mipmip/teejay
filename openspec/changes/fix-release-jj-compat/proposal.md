## Why

The release script (`scripts/release.sh`) uses git commands that break when the repo is jj-colocated. Specifically: `git branch --show-current` returns empty (detached HEAD), `git status --porcelain` may report jj state as dirty, and `git push origin main` fails without a checked-out branch. The cloudia-reader-aws project solved this same problem with pure git commands that work in both contexts.

## What Changes

- Replace `git status --porcelain` with `git diff --quiet && git diff --cached --quiet` for clean tree check (works with jj-colocated)
- Remove the `git branch --show-current` main branch check (unnecessary — if the tree is clean and the tag is correct, the branch doesn't matter)
- Replace `git push origin main` with `git push origin HEAD:main` (works even with detached HEAD under jj)
- Keep the script 100% pure git — no jj commands needed

## Capabilities

### New Capabilities
_None_

### Modified Capabilities
- `release-automation`: Release script safety checks and push commands become jj-colocated compatible

## Impact

- `scripts/release.sh`: Three targeted fixes to git commands
- No code changes — purely release tooling
