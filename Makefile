# 单元测试
.PHONY: ut
ut:
	@go test -race ./...

.PHONY: setup
setup:
	@sh ./script/setup.sh

.PHONY: lint
lint:
	golangci-lint run

.PHONY: fmt
fmt:
	@sh ./script/fmt.sh

.PHONY: tidy
tidy:
	@go mod tidy -v

.PHONY: check
check:
	@$(MAKE) --no-print-directory fmt
	@$(MAKE) --no-print-directory tidy

# e2e 测试
.PHONY: e2e
e2e:
	sh ./script/integrate_test.sh

.PHONY: dev_3rd_up
dev_3rd_up:
	docker-compose -f script/integration_test_compose.yml up -d

.PHONY: dev_3rd_down
dev_3rd_down:
	docker-compose -f script/integration_test_compose.yml down -v