#!/bin/zsh

script_dir=$(dirname "$0")
project_root=$(realpath "$script_dir")

build_dir="$project_root/target/build"

web_static_dir="$project_root/src/main/web-static"
web_static_build_dir="$build_dir/dist-web-static"
go_build_dir="$build_dir/src-go"

echo "===== Preparing build folders"
rm -rf "${build_dir:?}"
[ -d "$build_dir" ] || mkdir "$build_dir"
[ -d "$go_build_dir" ] || mkdir "$go_build_dir"
[ -d "$web_static_build_dir" ] || mkdir "$web_static_build_dir"

echo "===== Building frontend"
cd "$web_static_dir" || exit
npm run build

echo "===== Copying frontend resources for build"
cp -rf "$web_static_dir"/dist/* "$web_static_build_dir"

echo "===== Copying backend resources for build"
cp -rf "$project_root"/src/main/go/* "$go_build_dir"

echo "===== Copying docker resources for build"
cp "$project_root"/src/main/docker/backend/Dockerfile "$build_dir"

echo "===== Building docker image"
cd "$build_dir" || exit
docker build --tag open-asset-allocator .