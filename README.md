# nethsm-console-log

This utility logs the output from the NetHSM serial console to a file. When starting it sends two commands to the NetHSM console:

- `la` - enable all log channels
- `st` - print debug server status
