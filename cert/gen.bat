rm *.pem
echo 1. Generate CA's private key and self-signed certificate

openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=CH/ST=Zhejiang/L=Taizhou/O=Jiaget/OU=Jiaget/CN=*.jiaget.org/emailAddress=jiaget@outlook.com"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

echo 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=CH/ST=Zhejiang/L=Taizhou/O=pc book/OU=computer/CN=*.pcbook.org/emailAddress=pcbook@outlook.com"

echo 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in server-req.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf

echo "server's signed certificate"
openssl x509 -in server-cert.pem -noout -text

echo 3. Generate web client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj "/C=CH/ST=aa/L=bb/O=pc client/OU=computer/CN=*.pcclient.org/emailAddress=pcclient@outlook.com"

echo 3. Use CA's private key to sign web client's CSR and get back the signed certificate
openssl x509 -req -in client-req.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile client-ext.cnf

echo "client's signed certificate"
openssl x509 -in client-cert.pem -noout -text