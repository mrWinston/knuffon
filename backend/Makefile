build-dev:
	docker build --target dev -t knuffon-backend:dev .

run: build-dev
	docker run -it -v $(pwd):/code -p 8000:8000 knuffon-backend:dev
