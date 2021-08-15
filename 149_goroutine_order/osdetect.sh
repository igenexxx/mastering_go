#!/usr/bin/env bash

#
# Bash и кибербезопасность
# osdetect.sh
#
# Описание:
# Определение типа ОС: MS-Windows/Linux/MacOS
#
# Использование: bash osdetect.sh
# вывод будет одним из таких: Linux MSWin macOS

function detectOS() {
    if type -t wevtutil &>/dev/null; then
      OS=MSWin
    elif type -t scutil &>/dev/null; then
      OS=macOS
    else
      OS=Linux
    fi
}

detectOS

echo $OS