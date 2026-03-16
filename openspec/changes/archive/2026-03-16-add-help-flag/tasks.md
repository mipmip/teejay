## 1. Help Flag Detection

- [x] 1.1 Add help flag detection in parseFlags to check for --help or -h anywhere in args
- [x] 1.2 Return a boolean from parseFlags indicating if help was requested

## 2. Help Text Display

- [x] 2.1 Create printHelp function with formatted usage text showing commands and flags
- [x] 2.2 Call printHelp and exit(0) when help flag is detected, before other processing

## 3. Testing

- [x] 3.1 Manually test `tj --help` displays help and exits cleanly
- [x] 3.2 Manually test `tj -h` produces same output
- [x] 3.3 Manually test `tj add --help` shows help (not add command)
