#!/usr/bin/python3

from argparse import ArgumentParser
import requests
import concurrent.futures
import json
import sys
import time

def get(s: str, i: int) -> dict:
    a = {}
    a["Html"] = requests.get(s).content
    a["InputOrder"] = i
    return a

def test(l: list):
    export = {"Articles": None}
    articles = []
    c = 0
    with concurrent.futures.ThreadPoolExecutor() as executor:
        futures = []
        for url in l:
            futures.append(executor.submit(get, s=url, i=c))
            c = c + 1
        for future in concurrent.futures.as_completed(futures):
            articles.append(future.result())
    export["Articles"] = articles
    print(len(export["Articles"]))

if __name__ == "__main__":
    parser = ArgumentParser()
    parser.add_argument("data")
    args = parser.parse_args()
    t = time.perf_counter()
    test(args.data.split(","))
    print(round(time.perf_counter() - t, 2))
