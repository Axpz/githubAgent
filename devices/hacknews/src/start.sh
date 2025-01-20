#!/bin/bash

set -x

python main_hacknews_job.py

git config --global user.name "robot"
git config --global user.email "axpzhang@gmail.com"

git clone git@github.com:Axpz/xMinima.git || true
cd xMinima
cp ../reports/*hackernews_report.md ./_posts/
git add . && git commit -m "add hacknews $date"
git push --force