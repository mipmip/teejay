## 1. Fix clean tree check

- [x] 1.1 Replace `git status --porcelain` check with `git diff --quiet && git diff --cached --quiet`

## 2. Remove branch check

- [x] 2.1 Remove the `git branch --show-current` / main branch check block

## 3. Fix push command

- [x] 3.1 Replace `git push origin main` with `git push origin HEAD:main`

## 4. Verification

- [x] 4.1 Verify script syntax with `bash -n scripts/release.sh`
