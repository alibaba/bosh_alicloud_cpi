#!/usr/bin/env bash

set -e

: ${ALICLOUD_ACCESS_KEY_ID:?}
: ${ALICLOUD_SECRET_ACCESS_KEY:?}
: ${ALICLOUD_DEFAULT_REGION:?}
: ${GIT_USER_EMAIL:?}
: ${GIT_USER_NAME:?}

CURRENT_PATH=$(pwd)
SOURCE_PATH=$CURRENT_PATH/bosh-cpi-src
TERRAFORM_PATH=$CURRENT_PATH/terraform
TERRAFORM_MODULE=$SOURCE_PATH/ci/assets/terraform
TERRAFORM_METADATA=$CURRENT_PATH/terraform-metadata
METADATA=metadata
TERRAFORM_VERSION=0.10.0
TERRAFORM_PROVIDER_VERSION=1.2.4

wget -N https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip
wget -N https://github.com/alibaba/terraform-provider/releases/download/V${TERRAFORM_PROVIDER_VERSION}/terraform-provider-alicloud_linux-amd64.tgz

mkdir -p ${TERRAFORM_PATH}

unzip -o terraform_${TERRAFORM_VERSION}_linux_amd64.zip -d ${TERRAFORM_PATH}
tar -xzvf terraform-provider-alicloud_linux-amd64.tgz
mv -f bin/terraform* ${TERRAFORM_PATH}
rm -rf ./bin
export PATH="${TERRAFORM_PATH}:$PATH"

cd ${TERRAFORM_MODULE}

echo "Destroy terraform environment......"
terraform init
echo terraform destroy -var alicloud_access_key=${ALICLOUD_ACCESS_KEY_ID} -var alicloud_secret_key=${ALICLOUD_SECRET_ACCESS_KEY} -var alicloud_region=${ALICLOUD_DEFAULT_REGION}  \<\< EOF > terraform_destroy.sh
echo yes >> terraform_destroy.sh
echo EOF >> terraform_destroy.sh
chmod +x terraform_destroy.sh
./terraform_destroy.sh
echo "Destroy terraform environment successfully."
rm -rf ./terraform_destroy.sh

function copyToOutput(){

    cp -rf $1/. $2

    cd $2
    ls -la

    git config --global user.email ${GIT_USER_EMAIL}
    git config --global user.name ${GIT_USER_NAME}
    git config --local -l

    echo "&&&&&&&&git status"
    git status
    git status | sed |while read LINE
    do
        echo "&&&&&&&&hahah"
        echo $LINE
        if [[ $LINE == *detached* ]];
        then
            read -r -a Words <<< $LINE
            echo "**** $Words"
            echo "**** $Words[2]"
            echo "**** $Words[3]"
            git branch temp $Words[3]
            git checkout concourse_ci_tmp
            git merge temp
            git push origin concourse_ci_tmp
        fi
    done
    echo "***** git sattus"
    git status
    git add .
    git commit -m 'commit metadata'
    git status
    return 0
}

echo "Copy to output ......"
copyToOutput ${SOURCE_PATH} ${TERRAFORM_METADATA}