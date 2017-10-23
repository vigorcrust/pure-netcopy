appname := pnc
sources_client := $(wildcard ./ncclient/*.go)
sources_server := $(wildcard ./ncserver/*.go)

build = GOOS=$(1) GOARCH=$(2) go build -o build/$(appname)$(3) $(4)
tar = mkdir -p dist && cd build && tar -cvzf $(appname)_$(1)_$(2).tar.gz * && cp *.tar.gz ../dist && cd .. && rm -rf ./build
zip = mkdir -p dist && cd build && zip $(appname)_$(1)_$(2).zip * && cp *.zip ../dist && cd .. && rm -rf ./build

.PHONY: release

release: linux_build windows_build

copy_doc:
	mkdir -p build/
	cp README.md build/README.md
	cp LICENSE build/LICENSE

##### LINUX BUILDS #####
linux_build: build/linux_arm.tar.gz build/linux_arm64.tar.gz build/linux_386.tar.gz build/linux_amd64.tar.gz

build/linux_386.tar.gz: 
	$(MAKE) copy_doc
	$(call build,linux,386,client,$(sources_client))
	$(call build,linux,386,server,$(sources_server))
	$(call tar,linux,386)

build/linux_amd64.tar.gz: 
	$(MAKE) copy_doc
	$(call build,linux,amd64,client,$(sources_client))
	$(call build,linux,amd64,server,$(sources_server))
	$(call tar,linux,amd64)

build/linux_arm.tar.gz: 
	$(MAKE) copy_doc
	$(call build,linux,arm,client,$(sources_client))
	$(call build,linux,arm,server,$(sources_server))
	$(call tar,linux,arm)

build/linux_arm64.tar.gz: 
	$(MAKE) copy_doc
	$(call build,linux,arm64,client,$(sources_client))
	$(call build,linux,arm64,server,$(sources_server))
	$(call tar,linux,arm64)


##### WINDOWS BUILDS #####
windows_build: build/windows_386.zip build/windows_amd64.zip

build/windows_386.zip: 
	$(MAKE) copy_doc
	$(call build,windows,386,client,$(sources_client))
	$(call build,windows,386,server,$(sources_server))
	$(call zip,windows,386)

build/windows_amd64.zip: 
	$(MAKE) copy_doc
	$(call build,windows,amd64,client,$(sources_client))
	$(call build,windows,amd64,server,$(sources_server))
	$(call zip,windows,amd64)