image_name = knffon-backend

build: dev
	docker build --target prod -t $(image_name):prod .

dev:
	docker build --target dev -t $(image_name):dev .

lint: dev
	docker run -it --rm -v $(pwd):/code $(image_name):dev \
		bash -c "go vet && golint"

run: dev
	docker run -it --rm -v $(pwd):/code -p 8000:8000 $(image_name):dev

