
./clean.sh

BASE_CONFIG_DIR=/vagrant/setup
export FABRIC_CFG_PATH=$PWD

# Change this to see log meesage details
export ORDERER_GENERAL_LOGLEVEL=INFO

#2 Copy the YAML files
if [ ! -z $1 ]; then
    if [ $1 == "all" ]; then

        cp $BASE_CONFIG_DIR/crypto-config.yaml .
        cp $BASE_CONFIG_DIR/configtx.yaml .
        cp $BASE_CONFIG_DIR/orderer.yaml .
        echo    'Copied: crypto-config.yaml, configtx & orderer YAML files.'
        
    fi
else
    echo 'Use ./init.sh   all      to initialize crypto-config.yaml, configtx and orderer YAML'
fi


#3. Setup cryptogen for 
echo    '================ Generating crypto-config ================'
rm -rf ./crypto-config 2> /dev/null
cryptogen generate --config=./crypto-config.yaml

#4. Setup the genesis block
echo    '================ Writing genesis ================'
configtxgen -outputBlock ./channel-artifacts/genesis.block -profile OrdererGenesis -channelID ordererchannel

echo    '================ Writing channel ================'
configtxgen -profile Channel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID channel


