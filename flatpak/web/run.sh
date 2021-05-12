#!/bin/bash

# This is a helper script to start the axolotl backend and axolotl frontend with electron

programs=('echo starting axolotl' 'axolotl -e server' 'electron /app/bin/axolotl-web/dist')

for program in "${programs[@]}";
do
  eval "${program}";
done
