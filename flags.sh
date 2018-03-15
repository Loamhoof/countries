#!/bin/zsh

for file in data/*.svg
do
    echo $file
    inkscape -e flags/${$(basename $file)%.*}.png $file
done
