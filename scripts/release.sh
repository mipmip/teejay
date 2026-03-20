#!/usr/bin/env bash
set -e

# =============================================================================
# Teejay Release Script
# Automates semantic versioning releases with changelog updates
# =============================================================================

# --- Colors and Styles ---
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
BOLD='\033[1m'
NC='\033[0m' # No Color

# --- Helper Functions ---
info() {
    echo -e "${BLUE}ℹ${NC} $1"
}

success() {
    echo -e "${GREEN}✓${NC} $1"
}

warn() {
    echo -e "${YELLOW}⚠${NC} $1"
}

error() {
    echo -e "${RED}✗${NC} $1"
    exit 1
}

# --- Dependency Check ---
if ! command -v gum &> /dev/null; then
    error "gum is not installed. Install it with:

    brew install gum          # macOS
    nix-env -iA nixpkgs.gum   # Nix
    go install github.com/charmbracelet/gum@latest  # Go

See https://github.com/charmbracelet/gum for more options."
fi

# --- Safety Checks ---
info "Running safety checks..."

# Check for clean git working directory (works with both git and jj-colocated repos)
if ! git diff --quiet || ! git diff --cached --quiet; then
    error "Working tree is not clean. Commit or stash changes first."
fi
success "Working directory is clean"

# Check CHANGELOG.md exists and contains [Unreleased]
if [[ ! -f "CHANGELOG.md" ]]; then
    error "CHANGELOG.md not found"
fi

if ! grep -q "\[Unreleased\]" CHANGELOG.md; then
    error "CHANGELOG.md does not contain [Unreleased] section"
fi
success "CHANGELOG.md has [Unreleased] section"

# --- Read Current Version ---
VERSION_FILE="cmd/tj/VERSION"
if [[ ! -f "$VERSION_FILE" ]]; then
    error "VERSION file not found at $VERSION_FILE"
fi

CURRENT_VERSION=$(cat "$VERSION_FILE" | tr -d '\n')
info "Current version: ${BOLD}$CURRENT_VERSION${NC}"

# Parse version components
IFS='.' read -r MAJOR MINOR PATCH <<< "$CURRENT_VERSION"

# --- Version Selection ---
echo ""
echo -e "${BOLD}Select version bump type:${NC}"
BUMP_TYPE=$(gum choose "patch" "minor" "major")

# Calculate new version
case $BUMP_TYPE in
    major)
        NEW_VERSION="$((MAJOR + 1)).0.0"
        ;;
    minor)
        NEW_VERSION="$MAJOR.$((MINOR + 1)).0"
        ;;
    patch)
        NEW_VERSION="$MAJOR.$MINOR.$((PATCH + 1))"
        ;;
esac

info "New version will be: ${BOLD}$NEW_VERSION${NC}"

# Check if tag already exists
if git tag -l "v$NEW_VERSION" | grep -q "v$NEW_VERSION"; then
    error "Tag v$NEW_VERSION already exists"
fi
success "Tag v$NEW_VERSION is available"

# --- Confirmation ---
echo ""
echo -e "${BOLD}Release Summary:${NC}"
echo -e "  Current version: $CURRENT_VERSION"
echo -e "  New version:     $NEW_VERSION"
echo -e "  Bump type:       $BUMP_TYPE"
echo ""

if ! gum confirm "Proceed with release?"; then
    warn "Release cancelled"
    exit 0
fi

# --- Nix vendorHash Update ---
update_nix_vendor_hash() {
    if ! command -v nix &> /dev/null; then
        warn "nix is not installed — skipping vendorHash update in flake.nix"
        return 0
    fi

    info "Updating vendorHash in flake.nix..."

    # Save the current vendorHash
    OLD_HASH=$(grep 'vendorHash' flake.nix | sed 's/.*"\(.*\)".*/\1/')

    # Temporarily set vendorHash to empty to force nix to compute the correct one
    sed -i "s|vendorHash = \".*\"|vendorHash = \"\"|" flake.nix

    # Run nix build and capture the expected hash from the error output
    NIX_OUTPUT=$(nix build .#default 2>&1 || true)
    NEW_HASH=$(echo "$NIX_OUTPUT" | grep "got:" | sed 's/.*got: *//')

    if [[ -z "$NEW_HASH" ]]; then
        # Restore old hash if we couldn't determine the new one
        sed -i "s|vendorHash = \".*\"|vendorHash = \"$OLD_HASH\"|" flake.nix
        warn "Could not determine new vendorHash — restored previous hash"
        return 0
    fi

    # Update flake.nix with the correct hash
    sed -i "s|vendorHash = \".*\"|vendorHash = \"$NEW_HASH\"|" flake.nix
    success "Updated vendorHash to $NEW_HASH"
}

# --- Execute Release ---
echo ""
info "Creating release..."

# Update VERSION file
echo "$NEW_VERSION" > "$VERSION_FILE"
success "Updated VERSION file"

# Update CHANGELOG.md
TODAY=$(date +%Y-%m-%d)
sed -i "s/## \[Unreleased\]/## [Unreleased]\n\n## [$NEW_VERSION] - $TODAY/" CHANGELOG.md
success "Updated CHANGELOG.md with version $NEW_VERSION"

# Update nix flake vendorHash
update_nix_vendor_hash

# Commit changes
git add "$VERSION_FILE" CHANGELOG.md flake.nix
git commit -m "chore: release v$NEW_VERSION"
success "Created release commit"

# Create tag
git tag -a "v$NEW_VERSION" -m "Release v$NEW_VERSION"
success "Created tag v$NEW_VERSION"

# Push to remote
info "Pushing to remote..."
git push origin HEAD:main
git push origin "v$NEW_VERSION"
success "Pushed commit and tag to remote"

# --- Done ---
echo ""
echo -e "${GREEN}${BOLD}Release v$NEW_VERSION complete!${NC}"
echo ""
echo "GitHub Actions will now build and publish the release."
echo "Check the progress at: https://github.com/mipmip/teejay/actions"
