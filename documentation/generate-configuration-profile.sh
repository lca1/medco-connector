#!/bin/bash
set -e

# usage: bash generate-configuration-profile.sh CONFIGURATION_FOLDER KEYSTORE_PASSWORD NODE_DNS_1 NODE_IP_1 NODE_DNS_2 NODE_IP_3 NODE_DNS_3 NODE_IP_3 ...
if [ $# -lt 4 ]
then
    echo "Wrong number of arguments, usage: bash generate-configuration-profile.sh CONFIGURATION_FOLDER KEYSTORE_PASSWORD NODE_DNS_1 NODE_IP_1 NODE_DNS_2 NODE_IP_3 NODE_DNS_3 NODE_IP_3 ..."
    exit
fi

# arguments
CONF_FOLDER="$1"
KEYSTORE_PW="$2"
shift
shift

# clean up previous entries
mkdir -p "$CONF_FOLDER"
rm -f "$CONF_FOLDER/shrine.keystore" "$CONF_FOLDER/shrine_downstream_nodes.conf"

while [ $# -gt 0 ]
do
    NODE_DNS="$1"
    NODE_IP="$2"
    shift
    shift

    # generate the node certificate in the keystore and export it
    keytool -genkeypair -keysize 2048 -alias "$NODE_DNS" -validity 7300 \
        -dname "CN=$NODE_DNS, OU=$NODE_DNS, O=SHRINE Network, L=Lausanne, S=VD, C=CH" -ext "SAN=IP:$NODE_IP" \
        -keyalg RSA -keypass "$KEYSTORE_PW" -storepass "$KEYSTORE_PW" -keystore "$CONF_FOLDER/shrine.keystore"

    keytool -export -alias "$NODE_DNS" -storepass "$KEYSTORE_PW" -file "$CONF_FOLDER/$NODE_DNS.cer" -keystore "$CONF_FOLDER/shrine.keystore"
    #keytool -import -v -trustcacerts -alias "$NODE_IP" -file "$CONF_FOLDER/$NODE_IP.cer" -keystore "$CONF_FOLDER/shrine.keystore"  -keypass "$KEYSTORE_PW"  -storepass "$KEYSTORE_PW"


    # add entry in the downstream nodes
    echo "\"$NODE_DNS\" = \"https://$NODE_DNS:6443/shrine/rest/adapter/requests\"" >> "$CONF_FOLDER/shrine_downstream_nodes.conf"

    #todo: unlynx keys
done

keytool -list -v -keystore "$CONF_FOLDER/shrine.keystore" -storepass "$KEYSTORE_PW"
