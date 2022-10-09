#!/bin/bash
addr="46.101.46.156" 
addr1="repowatchdog.local" 
addr2="46.101.46.156" 
curl -XPOST -H 'Content-Type: application/x-www-form-urlencoded' \
-H 'User-Agent: GitHub-Hookshot/abbd694' \
-H 'X-GitHub-Delivery: 675ea740-465e-11ed-8747-21ffdd301373' \
-H 'X-GitHub-Event: pull_request' \
-H 'Accept: */*' \
-H 'X-GitHub-Hook-ID: 383154436' \
-H 'X-GitHub-Hook-Installation-Target-ID: 325321654' \
-H 'X-GitHub-Hook-Installation-Target-Type: repository' \
-H 'X-Hub-Signature: sha1=7ee3fd0256233ff0f8002dfd33fe1cbc4062d104' \
-H 'X-Hub-Signature-256: sha256=3ac63e367247eb05c33fab988cb4e8d1ad078bf530ed4300c6fdf823f0ebe997' \
-d @payload.json http://${addr1}/hook 

