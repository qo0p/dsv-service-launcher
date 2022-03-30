
# soap_client.py
# coding=utf-8

# pip install suds-py3

from suds.client import Client
from suds.cache import NoCache

ws_url = 'http://127.0.0.1:9091/dsvs/pkcs7/v1?wsdl'
pkcs7 = 'MIIikAYJKoZI.................'

ws_client = Client(ws_url, cache=NoCache())

ws_desc = ws_client
print(ws_desc)

reply = ws_client.service.verifyPkcs7(pkcs7)
print(reply)
