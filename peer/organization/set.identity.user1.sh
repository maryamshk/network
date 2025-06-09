# The env.sh file sets the identity to Peer
# This will set the identity to Admin 

export CORE_PEER_MSPCONFIGPATH=$CONFIG_DIRECTORY/crypto-config/peerOrganizations/Org1.com/users/User1@Org1.com/msp

echo ' Switched identity to User1 !!!'
