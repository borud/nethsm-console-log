# nethsm-console-log

This utility logs the output from the NetHSM serial console to a file. When starting it sends two commands to the NetHSM console:

- `la` - enable all log channels
- `st` - print debug server status

## Installing

If you have Go installed you can install `nethsm-console-log` with:

```shell
go install github.com/borud/nethsm-console-log/cmd/nethsm-console-log@latest
```
