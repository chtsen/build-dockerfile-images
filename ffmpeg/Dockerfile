FROM ubuntu:20.04

RUN set -ex && \
    apt update && \
    echo 'Yes, do as I say!'|apt install libtinfo5 && \
    apt install ncurses-base blender ffmpeg ghostscript imagemagick libglu1-mesa-dev -y
