#!/bin/bash

if ! command -v swagger &> /dev/null
then
	sudo make install_swagger
	chmod +x /usr/local/bin/swagger
fi

make docui