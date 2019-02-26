# Run Go server in debug mode

The most important part is to make sure that the target path where we copy
the sources MUST match the relative path under `$GOPATH` so that the debugger
can map the sources correctly. If we don't do this, we ARE NOT able to set/create
breakpoints in the sources.

## Gotchas

- The target path should match the path under `$GOPATH`, so that
The debugger can map the sources correctly.

# Links

- Example taken from: https://blog.jetbrains.com/go/2018/04/30/debugging-containerized-go-applications/
