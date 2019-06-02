# We shold run this after we install Go in the machine. See https://golang.org/doc/install

# Installing go on path
export PATH=$PATH:/usr/local/go/bin

CURRENT_DIR="$PWD"
echo $CURRENT_DIR

#Variables needed to build
export $GOPATH = $CURRENT_DIR
export PATH=$PATH:$GOPATH/bin