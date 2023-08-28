# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOLINUXBUILD=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build
APP_NAME=airport

all: build
build:
	# 编译
	$(GOBUILD) -o ${APP_NAME} main.go
	@# 输出到output
	@rm -rf output
	@mkdir output
	@# 守护进程设置执行权限
	#@chmod 755 supervise
	#@chmod 775 control.sh
	@cp -rf ./conf/ ./output/conf
	#@rm -rf ./output/conf/conf.go
	#@rm -rf ./output/conf/app/app.go
	#@rm -rf ./output/conf/README.md
	@mv ./${APP_NAME} ./output
	#@cp -rf ./control.sh ./output
	#@cp -rf ./supervise ./output/supervise.${APP_NAME}
	@cp -rf ./locales/ ./output/locales
	# 打包
	#cd output && tar -zcf ${APP_NAME}.tar.gz ./conf control.sh supervise.${APP_NAME} ${APP_NAME}

linux:
	# 编译
	$(GOLINUXBUILD) -o ${APP_NAME} main.go
	@# 输出到output
	@rm -rf output
	@mkdir output
	@# 守护进程设置执行权限
	@chmod 755 supervise
	@chmod 775 control.sh
	@cp -rf ./conf/ ./output/conf
	@rm -rf ./output/conf/conf.go
	@rm -rf ./output/conf/README.md
	@mv ./${APP_NAME} ./output
	@cp -rf ./control.sh ./output
	@cp -rf ./supervise ./output/supervise.${APP_NAME}
	# 打包
	cd output && tar -zcf ${APP_NAME}.tar.gz ./conf control.sh supervise.${APP_NAME} ${APP_NAME}

# go get -u github.com/swaggo/swag/cmd/swag
swagger:
	swag init -d ./,../../pkg/axgolib/req,../../pkg/axgolib/res  --output "./docs/swagger"

test:
	$(GOTEST) -v ./...

clean:
	@rm -rf output
	@if [ "$(shell uname)" == "Darwin" ]; then \
		find ./ -name "*.php~" |xargs -0 rm -rf {}; \
		find ./ -name "*.swo" |xargs -0 rm -rf {}; \
		find ./ -name "*.swp" |xargs -0 rm -rf {}; \
		find ./ -name "*~" |xargs -0 rm -rf {}; \
		find ./ -type d -name .git |xargs -0 rm -rf {}; \
	else \
		find ./ -name "*.php~" |xargs -i rm -rf {}; \
		find ./ -name "*.swo" |xargs -i rm -rf {}; \
		find ./ -name "*.swp" |xargs -i rm -rf {}; \
		find ./ -name "*~" |xargs -i rm -rf {}; \
		find ./ -type d -name .git |xargs -i rm -rf {}; \
	fi
