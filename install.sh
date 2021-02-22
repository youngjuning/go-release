#!/bin/sh
# Copyright 2020 youngjuning. All rights reserved. MIT license.
# TODO(everyone): Keep this script simple and easily auditable.
set -e

if ! command -v unzip >/dev/null; then
	echo "Error: unzip is required to install tpc (see: https://github.com/youngjuning/go-release#unzip-is-required)." 1>&2
	exit 1
fi

if [ "$OS" = "Windows_NT" ]; then
	target="x86_64-pc-windows-msvc"
else
	case $(uname -sm) in
	"Darwin x86_64") target="x86_64-apple-darwin" ;;
	*) target="x86_64-unknown-linux-gnu" ;;
	esac
fi

release_uri="$1/download/$2-${target}.zip"

if [ ! $3 ]; then
  home_dir=".$1"
else
  home_dir="$3"
fi
install_home="${install_home:-$HOME/$home_dir}"
bin_dir="$install_home/bin"
exe="$bin_dir/$2"

if [ ! -d "$bin_dir" ]; then
	mkdir -p "$bin_dir"
fi

curl --fail --location --progress-bar --output "$exe.zip" "$release_uri"
unzip -d "$bin_dir" -o "$exe.zip"
chmod +x "$exe"
rm "$exe.zip"

echo "$2 was installed successfully to $exe"
if command -v $2 >/dev/null; then
	echo "Run '$2 --help' to get started"
else
	case $SHELL in
	/bin/zsh) shell_profile=".zshrc" ;;
	*) shell_profile=".bash_profile" ;;
	esac
	echo "Manually add the directory to your \$HOME/$shell_profile (or similar)"
	echo "  export install_home=\"$install_home\""
	echo "  export PATH=\"\$install_home/bin:\$PATH\""
	echo "Run 'source $HOME/$shell_profile' to restart"
	echo "Run '$exe --help' to get started"
fi
