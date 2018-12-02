#!/bin/sh

echo "CREATE DATABASE IF NOT EXISTS \`gogollellero_test\` ;" | "${mysql[@]}"
echo "GRANT ALL ON \`gogollellero_test\`.* TO '"$MYSQL_USER"'@'%' ;" | "${mysql[@]}"
echo 'FLUSH PRIVILEGES ;' | "${mysql[@]}"
