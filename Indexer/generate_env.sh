# genere le .env
MYSQL_ROOT_PASSWORD=`openssl rand -base64 30`
MYSQL_DATABASE=eagle_db
MYSQL_USER=admin
MYSQL_PASSWORD=`openssl rand -base64 30`

cat > .env <<EOF
MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD
MYSQL_DATABASE=$MYSQL_DATABASE
MYSQL_USER=$MYSQL_USER
MYSQL_PASSWORD=$MYSQL_PASSWORD
EOF

# utilise le init.sql.template pour généré le vrais init.sql
DEFAULT_API_KEY_READ=`openssl rand -base64 30`
DEFAULT_API_KEY_WRITE=`openssl rand -base64 30`
cp init.sql.template init.sql

sed -i "s|{DEFAULT_API_KEY_READ}|$DEFAULT_API_KEY_READ|g; s|{DEFAULT_API_KEY_WRITE}|$DEFAULT_API_KEY_WRITE|g" init.sql

# afficher les creds
cat .env
echo "DEFAULT_API_KEY_READ=$DEFAULT_API_KEY_READ"
echo "DEFAULT_API_KEY_READ=$DEFAULT_API_KEY_WRITE"