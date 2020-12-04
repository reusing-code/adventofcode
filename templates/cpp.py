#!/usr/bin/python3

import argparse, os, shutil
from pathlib import Path


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("year", type=str)
    parser.add_argument("folder", help="name for folder and files. eg. day1", type=str)
    args = parser.parse_args()

    basepath = os.path.dirname(os.path.realpath(__file__))

    destpath = Path(basepath)/ ".." / args.year / args.folder
    destpath.mkdir(parents=True, exist_ok=False)

    for entry in os.scandir(Path(basepath) / "cpp"):
        if entry.is_file():
            newName = Path(entry.path).name.replace("@DAY@", args.folder).replace("@BUILD@", "BUILD")
            #shutil.copyfile(entry.path, destpath / newName)
            fin = open(entry.path, "rt")
            fout = open(destpath / newName, "wt")
            for line in fin:
                fout.write(line.replace('@DAY@', args.folder))
            fin.close()
            fout.close()


if __name__ == "__main__":
    main()
