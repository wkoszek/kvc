bin_path := ./bin/rcli
config_path := ~/.configrcli.json
install_path := /usr/bin/

build:
	go build -o $(bin_path)

rmconfig:
	if [ -f $(config_path) ]; then rm $(config_path); fi;

install: build rmconfig
	mv $(bin_path) $(install_path)