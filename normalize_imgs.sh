#! /bin/bash

# This script take all images of Lilou and make sure they are 1024x1024 webp images
# It is meant to be run by me, because I manually put the images in the public folder

IMG_SZ=1024

for file in public/lilou/*.jpg; do
   hash=$(md5sum "$file" | awk '{print $1}')
   out="./public/lilou/${hash}.webp"
   cwebp "$file" -resize $IMG_SZ $IMG_SZ -o "$out"
   magick "$out" -rotate 90 "$out"
   rm "$file"
done
