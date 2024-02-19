echo "***** Starting development environment *****"
date
echo "[1] Create docker network (cool-organic_network)"
docker network create -d bridge cool-organic_network

# Start Mysql
echo "[2] Create mysql"
docker run -d --name cool-organic_mysql --privileged=true --network cool-organic_network -v cool-organic-data:/var/lib/mysql  -e MYSQL_ROOT_PASSWORD="@Klov3x124n" -e MYSQL_USER="cool_organic" -e MYSQL_PASSWORD="@Klov3x124n" -e MYSQL_DATABASE="cool_organic" -p 3307:3306 bitnami/mysql:5.7
#docker run -d --name cool-organic_mysql --privileged=true -v cool-organic-data:/var/lib/mysql  -e MYSQL_ROOT_PASSWORD="@Klov3x124n" -e MYSQL_USER="cool_organic" -e MYSQL_PASSWORD="@Klov3x124n" -e MYSQL_DATABASE="cool_organic" -p 3307:3306 bitnami/mysql:5.7

# Start Backend Service
echo "[3] Create Backend"
docker build -t cool-organic_backend .
docker run --name cool-organic_backend -dp 8080:8080 --network cool-organic_network cool-organic_backend
#docker run --name cool-organic_backend -dp 8080:8080 cool-organic_backend


# Start API Gateway
#echo "[4] Create API Gateway"
#docker stop api-gateway
#docker rm api-gateway
#docker build -t api-gateway gateway/
#docker run -d --name api-gateway --net=bookstore-network  -p 443:443 api-gateway


echo "The following containers are running..."

docker ps

