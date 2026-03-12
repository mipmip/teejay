## Tmon - tmux activity monitor

## Why

I want to be free in my tmux use, but it's obvious that in this vibecoding age
we want to use tmux for running parallel agent sessions. I do not want to force
a coding workflow, and I do not want to integrate deeply with git etc. This TUI
app will just serve as a watch list for panes the user has added to the watch
list.

## Basic features

- All watching panes in a list in the left sidepanel
- Preview pane contents in main body
- Show status (busy/waiting for input)
- Detect proces (agent) and show as extra info
- open tmux pane from list for user input
- switch to tmux session and window and focus pane
- The use can add a pane to the watch list:
  - by running `[appname] add` from within the current pane which should be added
  - by browsing through sessions->windows->panes
- The user can remove a pane from the watch list
- When activity in a terminal session stops an notification can be configured per pane

## Techstack descisions

- go + lipgloss
- python + textual
- rust + ??
