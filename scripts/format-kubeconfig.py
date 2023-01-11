# This script allows you to remove lines with \n at the end of each line after encoding 
# a kube config file with base64 format
# Credits: https://aliartiza75.medium.com/gitlab-integration-with-kubernetes-46c2473a4396
f = open("config.txt", "r")
config = ""
for line in f:
    config += line.split("\n")[0]
print(config)