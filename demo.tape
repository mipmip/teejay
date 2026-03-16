# VHS Showcase for Teejay - Terminal Junky
# A tmux activity monitor for vibe coders
#
# Requirements:
# - VHS (charmbracelet/vhs) v0.7.0+
# - tmux
# - teejay binary (tj)
# - claude CLI
# - opencode CLI
#
# Run with: ./scripts/record-demo.sh (not directly with vhs)
# This ensures proper setup/cleanup of demo tmux sessions

# Output settings
Output demo.gif
Set FontSize 18
Set Width 1200
Set Height 700
Set Theme "Catppuccin Mocha"
Set Padding 20
Set Framerate 15
Set Shell "bash"
Set TypingSpeed 50ms

# Start teejay with the demo watchlist (pre-populated with 2 panes)
Type "./tj -w /tmp/tj-demo-watchlist.json"
Enter
Sleep 3s

# Show the pre-existing watched panes (claude-1 busy, opencode waiting)
# Give extra time for status detection to run
Sleep 3s

# Navigate between the pre-existing panes
Down
Sleep 800ms
Up
Sleep 1s

# Now demonstrate adding a pane using the browser
Type "a"
Sleep 1500ms

# Browser shows sessions with unwatched panes.
# aaa-claude-1 and aaa-opencode are already watched, so they don't appear.
# aaa-claude-3 should be the only session shown.
# Select aaa-claude-3 session
Enter
Sleep 1500ms

# Select the pane to add it via browser
Enter
Sleep 2s

# Now we have 3 panes - show the list
Sleep 1s

# Navigate to show all three panes
Down
Sleep 800ms
Down
Sleep 800ms
Up
Sleep 800ms

# Open configure popup with 'c'
Type "c"
Sleep 1500ms

# Toggle notification setting
Down
Sleep 500ms
Enter
Sleep 800ms

# Close configure popup
Type "q"
Sleep 1500ms

# Final view
Sleep 2s

# Clean exit
Type "q"
Sleep 500ms
