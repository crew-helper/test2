#!/bin/sh

target_dir="deploy"
version="v1t"
mkdir "${target_dir}"

#create all-in-one config
controller-gen crd:crdVersions=v1 rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases
kustomize build config/crd > "${target_dir}"/all-in-one.yaml
cd config/manager && kustomize edit set image controller="${INPUT_IMAGE_URL}"
cd - && kustomize build config/default >> "${target_dir}"/all-in-one.yaml

cp "${target_dir}"/all-in-one.yaml "${target_dir}"/all-in-one-${version}.yaml
cat "${target_dir}"/all-in-one.yaml
