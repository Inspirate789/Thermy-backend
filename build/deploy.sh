#!/bin/bash
# EXECUTE FROM PROJECT ROOT

set -e

rm -f .env

# Get host machine IP address:
cur_ip=$(nmcli device show | grep IP4.ADDRESS | head -1 | awk '{print $2}' | rev | cut -c 4- | rev)
echo -e "CURRENT_IP=${cur_ip}" >> .env


# Init Influx service:
influx_container_name="influx_log"
echo -e "INFLUXDB_CONTAINER_NAME=${influx_container_name}" >> .env

# Init Influx database:
influx_username="root"
influx_password=$(echo $RANDOM | md5sum | head -c 20)
influx_org="thermy"
influx_host=$cur_ip
influx_port=8086
influx_url="http://${influx_host}:${influx_port}"
frontend_bucket_name="frontend"
backend_bucket_name="backend"

echo -e "INFLUXDB_USERNAME=${influx_username}" >> .env
echo -e "INFLUXDB_PASSWORD=${influx_password}" >> .env
echo -e "INFLUXDB_ORG=${influx_org}" >> .env
echo -e "INFLUXDB_PORT=${influx_port}" >> .env
echo -e "INFLUXDB_URL=${influx_url}" >> .env
echo -e "INFLUXDB_FRONTEND_BUCKET_NAME=${frontend_bucket_name}" >> .env
echo -e "INFLUXDB_BACKEND_BUCKET_NAME=${backend_bucket_name}" >> .env

# Create Influx token:
gpg --gen-random -a 0 25
influx_token=$(gpg --gen-random -a 0 25)
echo -e "INFLUXDB_TOKEN=${influx_token}" >> .env



# Init PostgreSQL database:
postgres_user="root"
postgres_password=$(echo $RANDOM | md5sum | head -c 20)
postgres_host=$cur_ip
postgres_port=5432
postgres_dbname="thermy_db"
postgres_ssl_mode="disable"
postgres_driver_name="postgres"

echo -e "POSTGRES_USER=${postgres_user}" >> .env
echo -e "POSTGRES_PASSWORD=${postgres_password}" >> .env
echo -e "POSTGRES_HOST=${postgres_host}" >> .env
echo -e "POSTGRES_PORT=${postgres_port}" >> .env
echo -e "POSTGRES_DBNAME=${postgres_dbname}" >> .env
echo -e "POSTGRES_SSL_MODE=${postgres_ssl_mode}" >> .env
echo -e "POSTGRES_DRIVER_NAME=${postgres_driver_name}" >> .env



# Init Redis database:
redis_host="redis"
redis_port=6379
redis_password=$(echo $RANDOM | md5sum | head -c 20)

echo -e "REDIS_HOST=${redis_host}" >> .env
echo -e "REDIS_PORT=${redis_port}" >> .env
echo -e "REDIS_PASSWORD=${redis_password}" >> .env



# Init backend server port:
backend_port=8080
echo -e "BACKEND_PORT=${backend_port}" >> .env



echo -e ".env before run: "
cat .env



# Setup environment for backend program:
rm -f backend.env

echo -e "INFLUXDB_ORG=${influx_org}" >> backend.env
echo -e "INFLUXDB_URL=${influx_url}" >> backend.env
echo -e "INFLUXDB_BACKEND_BUCKET_NAME=${backend_bucket_name}" >> backend.env
echo -e "INFLUXDB_TOKEN=${influx_token}" >> backend.env
echo -e "POSTGRES_HOST=${postgres_host}" >> backend.env
echo -e "POSTGRES_PORT=${postgres_port}" >> backend.env
echo -e "POSTGRES_DBNAME=${postgres_dbname}" >> backend.env
echo -e "POSTGRES_SSL_MODE=${postgres_ssl_mode}" >> backend.env
echo -e "POSTGRES_DRIVER_NAME=${postgres_driver_name}" >> backend.env
echo -e "REDIS_HOST=${redis_host}" >> backend.env
echo -e "REDIS_PORT=${redis_port}" >> backend.env
echo -e "REDIS_PASSWORD=${redis_password}" >> backend.env
echo -e "BACKEND_PORT=${backend_port}" >> backend.env

echo -e "backend.env before run: "
cat backend.env



# Run system
docker-compose up -d
sleep 3



# Setup Influx database:
# Create buckets for services:
docker exec -ti $influx_container_name influx bucket create -n $frontend_bucket_name -o $influx_org -r 0
docker exec -ti $influx_container_name influx bucket create -n $backend_bucket_name -o $influx_org -r 0



# Setup PostgreSQL database:
admin_name="initial_admin"
admin_password=$(echo $RANDOM | md5sum | head -c 20)

echo -e "POSTGRES_ADMIN_USERNAME=${admin_name}" >> .env
echo -e "POSTGRES_ADMIN_PASSWORD=${admin_password}" >> .env

psql postgresql://$postgres_user:$postgres_password@$postgres_host:$postgres_port/$postgres_dbname \
  -f sql/init_db/create_fixed_tables.sql
psql postgresql://$postgres_user:$postgres_password@$postgres_host:$postgres_port/$postgres_dbname \
  -f sql/init_db/create_roles.sql
psql postgresql://$postgres_user:$postgres_password@$postgres_host:$postgres_port/$postgres_dbname \
  -f sql/init_db/create_admin.sql \
  -v username="${admin_name}" \
  -v quoted_username="'${admin_name}'" \
  -v quoted_password="'${admin_password}'"



echo -e "setup completed"
echo -e "complete .env: "
cat .env

# rm -f .env
