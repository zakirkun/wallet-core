docker run --name walletkita_stag_db -e MYSQL_ROOT_PASSWORD=W@lle3tKit4 -e MYSQL_DATABASE=walletkita_main_core -e MYSQL_USER=walletkita -e MYSQL_PASSWORD=W@Ll3tKita! -p 3306:3306 -d mysql:latest --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci