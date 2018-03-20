#!/usr/bin/python3

import csv
import os


def main():
    with open('zooms.csv', 'w') as csvfile:
        csvwriter = csv.writer(csvfile)

        for filename in sorted(os.listdir('maps')):
            csvwriter.writerow([filename[0:3], filename[4]])


if __name__ == '__main__':
    main()
