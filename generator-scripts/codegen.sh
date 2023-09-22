#!/bin/bash
#
# Copyright (c) 2022 by Delphix. All rights reserved.
#

OPEN_API_GENERATOR_URL=https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/6.3.0/openapi-generator-cli-6.3.0.jar

function die
{
    echo "$(basename $0): $*" >&2
    exit 1
}

function usage() {
    echo "usage: $(basename $0) api.yaml"
    exit 2
}

[[ $# -eq 1 ]] || usage

if [[ ! -f openapi-generator-cli.jar ]]; then
    echo "No generator tool found. Downloading it"
    wget $OPEN_API_GENERATOR_URL -O openapi-generator-cli.jar \
        || die "failed to download JAR"
fi

# generate a TMP directory
[[ ! -d "/tmp" ]] && mktemp
script_base_dir=$(dirname $0)

rm -rf $script_base_dir/../../test/*

java -jar openapi-generator-cli.jar generate \
    -i $1 -g go --global-property=skipFormModel=false -p enumClassPrefix=true -o tmp/ --git-host github.com --git-user-id delphix --git-repo-id dct-sdk-go --package-name delphix_dct_api  || die "failed to generate code"


# navigating back to the sdk directory
pushd $script_base_dir/tmp || die "failed to enter $script_base_dir/.."

#remove all pre-generated go files
rm -rf $script_base_dir/../../*.go

# move all the files from temp directory to outer except README
mv -v $(ls | grep -v README.md) $script_base_dir/../..

# move and rename README to outer directory
mv README.md $script_base_dir/../../OPENAPI-README.md

popd

# setting context to the outer directory
pushd $script_base_dir/../

# removing the template go mod file autogenerated.
rm -rf go.mod || die "failed to remove existing go.mod"

# initializing the module with 'github.com/delphix/dct-sdk-go'
go mod init github.com/delphix/dct-sdk-go && go mod tidy \
    || die "module initialization failed"

popd

# running the sanity test file
go run test.go || die "test failed"

# deleting the tmp directory
rm -rf $script_base_dir/tmp