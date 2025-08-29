#!/usr/bin/env dagger

container |
from alpine |
with-exec cat /etc/os-release |
stdout
