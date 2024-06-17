#!/usr/bin/python3
import paramiko 
import subprocess
import sys
import re
import secret
import time


port = 22
remotepathWww = "/Server/www/"
remotepathServer = "/Server/"
localpathWww = "./Server/www/"
localpathServer = "./Server/"


# Удалить __pycache__
def delPycache():	
	subprocess.run(["rm", "-r", "__pycache__"])


# Connect SSH
ssh = paramiko.SSHClient()
ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
ssh.connect(hostname = secret.host, username = secret.user, password = secret.password, port = port)
	
delPycache()

t = time.time()
local = time.localtime(t)

print("DOORS Server upper is START:", time.asctime(local))
print("-----------------------------------------------------")
	
while (True):
	isRun = False
	# ps | grep DOORS
	stdin, stdout, stderr = ssh.exec_command("ps | grep DOORS")
	strings = stdout.read().decode().split('\n') 
	
	for str in strings:
		str = re.sub(r'\s+', ' ', str)
		str = str.strip()
		words = str.split(' ')
		try:
			if words[4] == "./DOORS":
				isRun = True
		except:
			pass
	if isRun != True:
		t = time.time()
		local = time.localtime(t)
		print("DOWN: Server is DOWN:       ", time.asctime(local))
		
		# Killall DOORS
		stdin, stdout, stderr = ssh.exec_command("killall DOORS")
		#print("Stop   SERVER OK")
	
		# Chmod +x DOORS
		stdin, stdout, stderr = ssh.exec_command("cd " + remotepathServer)
		stdin, stdout, stderr = ssh.exec_command("chmod +x DOORS")
		#print(stdout.read().decode())
		#print("Chmod  SERVER OK")

		stdin, stdout, stderr = ssh.exec_command("/root/startDOORS.sh > /dev/null 2>&1")
		print("UP:   Start  SERVER:        ", time.asctime(local))
		
		isRun = True




ssh.close()
del ssh, stdin, stdout, stderr

	
	


	

	
	
	

	


	
	


	
	
	
	

	


