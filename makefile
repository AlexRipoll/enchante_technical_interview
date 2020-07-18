
mysql_create_tables:
	docker exec -i $(docker-compose ps -q db) mysql -u root -proot enchainte_db < config/mysql/database.sql

mysql-shell:
	docker-compose exec mysql mysql -u root -proot
