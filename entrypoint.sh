echo "starting up dbg"

/go/bin/dlv --listen=:30123 --headless=true --api-version=2 exec ./app

#./app