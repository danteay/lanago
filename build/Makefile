build:
	sh build/build.sh

install:
	go get -u -v github.com/kardianos/govendor
	govendor sync

publish: build
	docker push registry.heroku.com/lanago/web

release:
	heroku container:release web -a lanago