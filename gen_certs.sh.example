echo "Creating certs folder ..."
mkdir certs && cd certs

echo "Generating certificates ..."

openssl genrsa -passout pass:1111 -des3 -out ca.key 4096

openssl req -passin pass:1111 -new -x509 -days 3650 -key ca.key -out ca.crt -subj  "/C=CL/ST=RM/L=Santiago/O=Test/OU=Test/CN=ca"

openssl genrsa -passout pass:1111 -des3 -out server.key 4096

openssl req -passin pass:1111 -new -key server.key -out server.csr -subj  "/C=CL/ST=RM/L=Santiago/O=Test/OU=Server/CN=localhost"

openssl x509 -req -passin pass:1111 -days 3650 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt

openssl rsa -passin pass:1111 -in server.key -out server.key