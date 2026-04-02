## Context

`calcMultiColumn()` uses `const minColWidth = 30` and divides available width by this minimum to get the column count. It doesn't account for the actual number of items — 4 items on a wide terminal still creates 6+ narrow columns with empty ones.

## Goals / Non-Goals

**Goals:**
- Wider, more readable columns
- No wasted space from empty columns

**Non-Goals:**
- Making column width configurable (the adaptive behavior should handle most cases)

## Decisions

### Increase minColWidth to 45

45 chars gives ~43 usable chars per item after padding. This comfortably fits a pane name + truncated breadcrumb without aggressive ellipsis.

### Cap columns to item count

```go
numColumns = min(numColumns, totalItems)
```

This ensures every column has at least one item and redistributes the extra width to make each column wider.
