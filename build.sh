#!/usr/bin/env bash

set -e

export PATH="$(pwd)/php/bin:$PATH"
export LD_LIBRARY_PATH="$(pwd)/php/lib:$LD_LIBRARY_PATH"
# Use php-config to get the correct includes and linker flags


if [[ $1 == "build" ]]; then
    export CGO_CFLAGS="$($(pwd)/php/bin/php-config --includes)"
    export CGO_LDFLAGS="-Lphp/lib -lphp $($(pwd)/php/bin/php-config --libs)"

    echo $CGO_CFLAGS
    echo $CGO_LDFLAGS
    go build .
elif [[ $1 == "run" ]]; then
    export CGO_CFLAGS="$($(pwd)/php/bin/php-config --includes)"
    export CGO_LDFLAGS="-Lphp/lib -lphp $($(pwd)/php/bin/php-config --libs)"

    shift
    go run . "$@"
elif [[ $1 == "conf" ]]; then
    echo "Make sure you're running this in the PHP source file directory!"
    ./configure --prefix=/home/tristan/Development/whoa/php --enable-static=yes \
            --disable-shared \
            --enable-embed=static \
             --enable-cli \
            --with-zlib \
            --with-openssl \
            --enable-mbstring \
            --with-curl \
            --enable-pdo \
            --with-pdo-mysql \
            --with-mysqli \
            --with-sqlite3 \
            --enable-ctype \
            --enable-json \
            --enable-session \
            --enable-simplexml \
            --enable-tokenizer \
            --enable-xml \
            --enable-xmlreader \
            --enable-xmlwriter \
            --enable-cli \
            --enable-zts \
            --disable-zend-signals \
            --enable-zend-max-execution-timers 
    # --with-curl --with-zlib --with-openssl --with-zip --enable-mbstring --with-mysqli --with-pdo-mysql
    make -j$(nproc)
    sudo make install
    # make test
else 
    echo "Please run with cmds: 'run' or 'build'."
fi
