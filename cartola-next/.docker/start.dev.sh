#!/bin/bash

if [ ! -f ".env.local" ]; then
    cp .env.example .env.local
fi

yarn install

tail -f /dev/null