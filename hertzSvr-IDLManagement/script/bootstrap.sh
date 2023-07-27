
#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=hertzSvr-IDLManagement
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}
