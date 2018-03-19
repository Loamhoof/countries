#!/bin/zsh

for file in data/flags/*.svg
do
    echo $file
    inkscape -e flags/${$(basename $file)%.*}.png $file
done
