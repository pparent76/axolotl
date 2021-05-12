#!/bin/bash

# This is a helper script to start the axolotl backend and axolotl frontend with electron

programs=(
  'echo starting axolotl-web'
  'electron /app/bin/axolotl-web/dist'
  'echo starting axolotl'
  'axolotl -e server'
)

for program in "${programs[@]}";
do
  eval "${program}";
done
