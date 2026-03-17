#!/usr/bin/env bash
# Record the Teejay demo GIF using VHS
#
# This script:
# 1. Creates isolated demo tmux sessions (aaa-* prefix for alphabetical sorting)
# 2. Sets up a demo watchlist with pre-existing monitored panes
# 3. Runs VHS to record the demo
# 4. Cleans up the demo sessions and watchlist
#
# Usage: ./scripts/record-demo.sh

set -e

cd "$(dirname "$0")/.."

# Session names with aaa- prefix to ensure they sort first alphabetically
SESSIONS=("aaa-claude-1" "aaa-claude-3" "aaa-opencode")
DEMO_WATCHLIST="/tmp/tj-demo-watchlist.json"

cleanup() {
    for session in "${SESSIONS[@]}"; do
        tmux kill-session -t "$session" 2>/dev/null || true
    done
    rm -f "$DEMO_WATCHLIST"
}

# Always cleanup on exit
trap cleanup EXIT

# Clean any existing demo sessions first
cleanup

# Session 1: Claude session - short task that will finish during recording
tmux new-session -d -s aaa-claude-1 -n agent
tmux send-keys -t aaa-claude-1 "claude 'Write a function in Go that reverses a string. Include a brief explanation.'" Enter

# Session 2: OpenCode session - just start opencode without a task (shows waiting screen)
tmux new-session -d -s aaa-opencode -n main
tmux send-keys -t aaa-opencode "opencode" Enter

# Session 3: Another Claude session - this one we'll add via browser
# Task should take ~15-20 seconds so it's still busy when we add it, then finishes during the demo
tmux new-session -d -s aaa-claude-3 -n coding
tmux send-keys -t aaa-claude-3 "claude 'Write a complete REST API in Go with net/http that has endpoints for GET /users (returns a JSON list), POST /users (creates a user), and GET /users/{id} (returns a single user). Include a User struct with ID, Name, and Email fields, an in-memory store using a map, proper JSON encoding/decoding, error handling for missing users, and a brief explanation of the design choices.'" Enter

# Give sessions time to start and begin processing
# Wait longer to ensure claude sessions have started and are showing activity
sleep 5

# Get pane IDs for our demo sessions (claude-1 and opencode will be pre-populated)
PANE1=$(tmux list-panes -t aaa-claude-1 -F '#{pane_id}' | head -1)
PANE2=$(tmux list-panes -t aaa-opencode -F '#{pane_id}' | head -1)

# Create a pre-populated watchlist with claude-1 (busy) and opencode (waiting)
# Format must be an array of panes, not a map
cat > "$DEMO_WATCHLIST" << EOF
{
  "panes": [
    {
      "id": "$PANE1",
      "name": "claude-1"
    },
    {
      "id": "$PANE2",
      "name": "opencode"
    }
  ]
}
EOF

# Run VHS with a clean bash environment
# - PS1: minimal prompt (just $)
# - HISTFILE: no history
# - HOME: use temp dir to avoid loading .bashrc
# - Keep TMUX socket so tj can access the demo sessions
export PS1='$ '
export HISTFILE=
export HOME=/tmp
# Don't unset TMUX - tj needs it to access the sessions
vhs demo.tape
