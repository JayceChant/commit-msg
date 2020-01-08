# commit-msg

A git commit-msg hook by go.



## Install

Just put the executable binary into the hook directory of your project which using git as VCS.

Assuming the project root is `/`, the hook directory should be `/.git/hooks/` . Please make sure the binary is named as `commit-msg` or `commit-msg.exe` (win) , and the executable permission is granted.

If you can't find the hook directory, make sure the hidden directories are shown.



After installation, git will call it automatically to validate commit message, no manual calls are required.

### From pre-built binary

You can download the latest pre-built binary from [here](https://github.com/JayceChant/commit-msg/releases) . However, only win64 binaries are available.



### Build and install from source code

Please setup the go environment and build it yourself.

Sorry for the poor instructions.