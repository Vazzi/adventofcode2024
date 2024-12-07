#!/bin/bash
if [ -z "$1" ]; then
  echo "Usage: $0 <number>"
  exit 1
fi

NUMBER=$(printf "%02d" "$1")

DAY_FOLDER="day${NUMBER}"
mkdir -p "$DAY_FOLDER"

FIRST_LINE="package day${NUMBER}"

echo "${FIRST_LINE}" > "$DAY_FOLDER/main.go"
echo "${FIRST_LINE}" > "$DAY_FOLDER/main_test.go"
echo "${FIRST_LINE}" > "$DAY_FOLDER/firstSolution.go"

INPUT_FOLDER="./inputData"
mkdir -p "$INPUT_FOLDER/test"
touch "$INPUT_FOLDER/day${NUMBER}.txt"
touch "$INPUT_FOLDER/test/day${NUMBER}.txt"

echo "Folder and files for day${NUMBER} created successfully."