#!/usr/bin/env bash

cd "$(dirname $0)/src"

for i in test_*; do
  echo "Creating snapshots for $i"
  woext="${i%.*}"
  treelike -f "$i" -D > "../snapshots/${woext}_snapshot_no_root.txt"
  treelike -f "$i" -p > "../snapshots/${woext}_snapshot_full_path.txt"
  treelike -f "$i" -s > "../snapshots/${woext}_snapshot_trailing_slash.txt"
  treelike -f "$i" -c ascii > "../snapshots/${woext}_snapshot_ascii.txt"
done

echo "Done"
