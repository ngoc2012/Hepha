#!/bin/sh

# Some commands to test directly on the container:
# mysqld --user=mysql > ./mysql.log 2>&1 &
# mysql_secure_installation
# mysqld restart --user=mysql > ./mysql.log 2>&1 &

# Configuration files
# vim /etc/my.cnf
# vim /etc/my.cnf.d/mariadb-server.cnf

# Inialization mysql
chown hefa:hefa /var/lib/mysql
chmod 750 /var/lib/mysql

mkdir -p /run/mysqld
chown -R hefa:hefa /run/mysqld
chmod 750 /var/lib/mysql

# Run with mysql as user
mysql_install_db --user=hefa --datadir=/var/lib/mysql

mysqld --user=hefa &

# Wait for MySQL to start
until mysqladmin ping --silent; do
  echo "Waiting for MySQL to start..."
  sleep 2
done

# Testing server
# mysqladmin version
# mysqlshow mysql

echo "Set root passwordm "$DB_ROOT_PASSWORD
# Start MySQL secure installation with Expect
expect << EOF
spawn mysql_secure_installation

expect "Enter current password for root (enter for none):"
send "\r"

expect "Switch to unix_socket authentication"
send "y\r"

expect "Change the root password?"
send "y\r"

expect "New password:"
send "$DB_ROOT_PASSWORD\r"

expect "Re-enter new password:"
send "$DB_ROOT_PASSWORD\r"

expect "Remove anonymous users?"
send "y\r"

expect "Disallow root login remotely?"
send "y\r"

expect "Remove test database and access to it?"
send "y\r"

expect "Reload privilege tables now?"
send "y\r"

expect eof
EOF

mysqladmin -u root -p"$DB_ROOT_PASSWORD" shutdown
mysqld_safe --user=hefa --datadir=/var/lib/mysql &

# Wait for MySQL to start again
until mysqladmin ping --silent; do
  echo "Waiting for MySQL to restart..."
  sleep 2
done

# echo "Create database" $DB_ROOT_PASSWORD
# echo "Create user" $DB_USER
# mysql -u root -p"$DB_ROOT_PASSWORD" <<EOF
# CREATE USER '$DB_USER'@'%' IDENTIFIED BY '$DB_PASSWORD';
# GRANT ALL PRIVILEGES ON *.* TO '$DB_USER'@'%' WITH GRANT OPTION;
# FLUSH PRIVILEGES;
# EOF
# mysql -u root -p"$DB_ROOT_PASSWORD" < /app/init.sql
# mysql -u root -p"$DB_ROOT_PASSWORD" < /app/create_db.sql

mysqladmin -u root -p"$DB_ROOT_PASSWORD" shutdown

mysqld_safe --user=hefa --datadir=/var/lib/mysql --skip-networking=0 --bind-address=0.0.0.0 --skip-name-resolve

# #skip-networking
# If you enable skip-networking by removing the hash (#) symbol, the database server will only accept connections from local clients using Unix sockets or named pipes. This means remote clients won't be able to connect to the database server.
# Conversely, if you disable skip-networking (which is the default setting in most cases), the database server will listen for connections on a network interface, allowing both local and remote clients to connect, provided they have the necessary permissions.
