VERSION := 1.4.0

build:
	docker build -t ingress-monitor:$(VERSION) .

publish: build
	docker tag ingress-monitor:$(VERSION) mangoweb/ingress-monitor:$(VERSION)
	docker push mangoweb/ingress-monitor:$(VERSION)
