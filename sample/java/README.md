# Import WSDL and generate SOAP client code

```
wsimport -keep -verbose http://127.0.0.1:9091/dsvs/pkcs7/v1?wsdl
```

# Compile Client.java file

```
javac Client.java
```

# Run Client

```
java Client
```
