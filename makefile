
help:
	echo "ToDo - Make help"

mysql-run: ## Spins up a mysql instance for dev purposes
	docker run --name mysql01 \
	--rm \
	-p 3306:3306 \
	-v /storage/docker/mysql-datadir:/var/lib/mysql \
	-e MYSQL_ROOT_PASSWORD=Welcome1234 \
	-e MYSQL_ROOT_HOST=% \
	-d \
	mysql:8.0

mysql-kill: ## Kills the mysql database
	docker kill mysql01

mysql-login: ## Logs in to the local mysql instance
	mysql -uroot -h 127.0.0.1 -pWelcome1234

test-mysql: ## Spins ups a mysql instance + runs integration tests

	# Spin up new MySQL
	docker run -p 3306:3306 \
	  --rm \
      --name acc-test-mysql \
      -e MYSQL_ROOT_PASSWORD=Welcome1234 \
      -d mysql:8.0

	sleep 1s

    # Run tests
	go test ./stores/... -tags=mysql

	docker kill acc-test-mysql