#!/usr/bin/bash

# the substring match this does is case-sensitive
busctl --user call \
    org.gnome.Shell \
    /de/lucaswerkmeister/ActivateWindowByTitle \
    de.lucaswerkmeister.ActivateWindowByTitle \
    activateBySubstring \
    s 'goraylib-studies - vscode'
