build:
	sh build/build.sh

install:
	go get -u -v github.com/kardianos/govendor
	govendor sync
	
publish:
	docker tag lanago:latest registry.heroku.com/lanago/web 
	docker push registry.heroku.com/lanago/web 
	
release:
	heroku container:release web -a lanago