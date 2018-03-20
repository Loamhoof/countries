#!/usr/bin/python3

import os.path
import re


def main():
    for filename in os.listdir('maps'):
        os.rename(os.path.join('maps', filename), os.path.join('maps', re.sub(r'(\w+)-\d', r'\1', filename)))


if __name__ == '__main__':
    main()
