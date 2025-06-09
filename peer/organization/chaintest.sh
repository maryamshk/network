# Run the test
# Install & Instantiate the chaincode

export CORE_CHAINCODE_LOGGING_LEVEL=INFO

#1  Install - will fail if already installed
peer chaincode install -n goccdev -v 1.0  -p chaincode_example02

#2  Instantiate - will fail if already instantiated
peer chaincode instantiate  -n goccdev -v 1.0 -C channel -c '{"Args":["init","a","100","b","200"]}'

peer chaincode invoke  -C channel -n goccdev  -c '{"Args":["invoke","b","a","5"]}'
peer chaincode query -C channel -n goccdev  -c '{"Args":["query","a"]}'