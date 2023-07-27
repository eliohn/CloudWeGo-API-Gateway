
#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=hertzSvr-Gateway
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}
