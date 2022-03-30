<?php

$ws_url = "http://127.0.0.1:9091/dsvs/pkcs7/v1?wsdl";
$pkcs7 = "MIIikAYJKoZI.................";

// Initialize WS with the WSDL
$client = new SoapClient($ws_url);

// Invoke WS method (verifyPkcs7) with the request params 
$response = $client->__soapCall("verifyPkcs7", array($pkcs7));

// Print WS response
var_dump($response);

?>
