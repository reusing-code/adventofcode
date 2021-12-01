#!/usr/bin/python3

import argparse
import os
import shutil
from pathlib import Path


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "language", help="The target language. Options: go, cpp", type=str)
    parser.add_argument("year", help="the year folder eg. 2020", type=str)
    parser.add_argument(
        "folder", help="name for folder and files. eg. day1", type=str)
    args = parser.parse_args()

    if (args.language != "go" and args.language != "cpp"):
        print("incorrect language, needs to be either go or cpp")
        return 1

    basepath = os.path.dirname(os.path.realpath(__file__))

    destpath = Path(basepath) / ".." / args.year / args.folder
    destpath.mkdir(parents=True, exist_ok=False)

    for entry in os.scandir(Path(basepath) / args.language):
        if entry.is_file():
            newName = Path(entry.path).name.replace(
                "@DAY@", args.folder).replace("@BUILD@", "BUILD")
            with open(entry.path, "rt") as fin, open(destpath / newName, "wt") as fout:
                for line in fin:
                    fout.write(line.replace(
                        '@DAY@', args.folder).replace('@YEAR@', args.year).replace('@EXT@', args.language))


if __name__ == "__main__":
    main()
