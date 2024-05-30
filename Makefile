INPUT = 40

run-go:
	@go run cmd/go/main.go $(INPUT) > go_output.txt

run-js:
	@node cmd/js/index.js $(INPUT) > js_output.txt

run-py:
	@python3 cmd/python/main.py $(INPUT) > py_output.txt

run:
	@parellel --joblog joblog.log make run-go make run-js make run-py &
	@echo "Go output:"
	@cat go_output.txt
	@echo ""
	@echo "JS output:"
	@cat js_output.txt
	@echo ""
	@echo "PY output:"
	@cat py_output.txt

