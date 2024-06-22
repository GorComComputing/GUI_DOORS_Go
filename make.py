#!/usr/bin/python3
'''
Эта программа на Python предназначена для управления сборкой и развертыванием проекта под названием "DOORS", 
который включает веб-сервер и WASM (WebAssembly) компоненты. 
Программа выполняет различные задачи, такие как сборка, загрузка на сервер, запуск, 
остановка и проверка состояния сервера, а также коммит и пуш изменений в репозиторий на GitHub. 
'''

import paramiko 
import subprocess
import sys
import re
import secret

userGit = "GorComComputing"
repoGit = "GUI_DOORS_Go"

port = 22
remotepathWww = "/DOORS/www/"
remotepathServer = "/DOORS/"
localpathWww = "./Server/www/"
localpathServer = "./Server/"


# Удалить __pycache__
def del_pycache():	
	subprocess.run(["rm", "-rf", "__pycache__"])
	
	
# Вывод сообщения об ошибке
def print_error_message():
    print("ERROR. Выбери один из параметров:")
    print("  wasm    - собрать DOORS.wasm")
    print("  srv     - собрать сервер DOORS")
    print("  start   - запустить сервер DOORS")
    print("  stop    - остановить сервер DOORS")
    print("  stat    - узнать запущен ли сервер DOORS")
    print("  git     - commit to GitHub")
    print("  restart - перезапустить сервер DOORS")
    print("  docker  - создать Docker контейнер")
    print("  hub     - отправить контейнер в Docker Hub")


# Подключение по SSH
def connect_ssh():
    ssh = paramiko.SSHClient()
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    ssh.connect(hostname=secret.host, username=secret.user, password=secret.password, port=port)
    return ssh
    
    
# Подключение по SFTP
def connect_sftp():
    transport = paramiko.Transport((secret.host, port))
    transport.connect(username=secret.user, password=secret.password)
    sftp = paramiko.SFTPClient.from_transport(transport)
    return sftp, transport
    
    
# Убить процессы launch.sh и DOORS
def kill_processes(ssh):
    ssh.exec_command("killall launch.sh")
    ssh.exec_command("killall DOORS")
    print("-----> Stop Server OK")
    
    
# Загрузить файл на сервер
def upload_file(sftp, localpath, remotepath):
    sftp.put(localpath, remotepath)
    print(f"-----> Uploaded {localpath} to {remotepath}")
    
    
# Выполнить команду make
def make_command(command, directory=None):
    cmd = ["make"]
    if directory:
        cmd.extend(["-C", directory])
    subprocess.run(cmd)
    print(f"MAKE   {command} OK")
    
    
# Выполнить несколько команд по SSH
def run_ssh_commands(ssh, commands):
    for command in commands:
        ssh.exec_command(command)
        print(f"-----> Run command: {command}")
        
        
# Создание Docker контейнера
def docker_build():
	print("DOCKER BUILD ...")
	proc1 = subprocess.run([f"docker", "ps"], stdout=subprocess.PIPE)
	#stdout = proc1.communicate()
	print("!!!")
	print(f"{proc1.stdout}")
	
	
	proc2 = subprocess.run([f"grep", f"'{secret.dockerUser}/{secret.dockerRepo}'"], input=proc1.stdout, stdout=subprocess.PIPE)
	
	#stdout2 = proc2.communicate()
	print("!!!2")
	print(f"{proc2.stdout}")
	
#	proc3 = subprocess.run([f"awk", "'{{print $1}}'"], input = proc2.stdout.encode("utf-8"), capture_output = True, text=True)
	#out, _ = proc3.communicate()
#	print("!!!3")
#	print("{0}".format(proc3.stdout))
	
	
	
	
	
	
	#subprocess.run(["docker", "stop", f"{proc3.stdout}"])
	print(f"-----> Stop container: {secret.dockerUser}/{secret.dockerRepo}")
	subprocess.run(["docker", "rm", "$(docker", "ps", "-a", "-q", "-f", "status=exited)"])
	print(f"-----> Remove all exited containers")
	subprocess.run(["docker", "rmi", "$(docker", "images", "|", "grep", f"'{secret.dockerUser}/{secret.dockerRepo}'", "|", "awk", "'{print $3}')"])
	print(f"-----> Remove container: {secret.dockerUser}/{secret.dockerRepo}")
	subprocess.run(["docker", "build", "-t", f"{secret.dockerUser}/{secret.dockerRepo}", "."])
	print("DOCKER BUILD OK")
        
        
# Отправка контейнера в Docker Hub
def docker_push():    
	subprocess.run(["docker", "login", "-u", f"{secret.dockerLogin}", "-p", f"{secret.dockerPassword}", "docker.io"])
	print("DOCKER login OK")

	subprocess.run(["docker", "push", f"{secret.dockerUser}/{secret.dockerRepo}"])
	print("DOCKER push OK")


if len(sys.argv) != 2 and not (len(sys.argv) == 3 and sys.argv[1] == "git"):
    print_error_message()
    del_pycache()
    sys.exit(1)

valid_commands = [
				  "wasm", 
				  "srv", 
				  "start", 
				  "stop", 
				  "stat", 
				  "git", 
				  "restart", 
				  "docker",
				  "hub",
				  ]
command = sys.argv[1]

if command not in valid_commands:
    print_error_message()
    del_pycache()
    sys.exit(1)

if (sys.argv[1] == "wasm"):	
	# Сборка WASM
	make_command("WASM")
	print("UPLOAD WASM ...")
	sftp, transport = connect_sftp()
	upload_file(sftp, localpathWww + "DOORS.wasm", remotepathWww + "DOORS.wasm")
	upload_file(sftp, localpathWww + "index.html", remotepathWww + "index.html")
	sftp.close()
	transport.close()
	del transport, sftp 
	print("UPLOAD WASM OK")
	del_pycache()
    
	
elif (sys.argv[1] == "srv"):
	# Сборка сервера
    make_command("SERVER", localpathServer)
    print("UPLOAD SERVER ...")
    ssh = connect_ssh()
    kill_processes(ssh)
    sftp, transport = connect_sftp()
    upload_file(sftp, localpathWww + "index.html", remotepathWww + "index.html")
    upload_file(sftp, localpathServer + "DOORS", remotepathServer + "DOORS")
    upload_file(sftp, localpathServer + "auto_doors", "/etc/init.d/auto_doors")
    upload_file(sftp, localpathServer + "launch.sh", remotepathServer + "launch.sh")
    upload_file(sftp, localpathServer + "startDOORS.sh", remotepathServer + "startDOORS.sh")
    sftp.close()
    transport.close()
    print("UPLOAD SERVER OK")
    print("START  SERVER ...")
    run_ssh_commands(ssh, [
    					   f"chmod +x {remotepathServer}launch.sh", 
    					   f"chmod +x {remotepathServer}startDOORS.sh", 
    					   f"chmod +x {remotepathServer}DOORS", 
    					   f"{remotepathServer}launch.sh &"
    					   ])
    
    print("START  SERVER OK")
    ssh.close()
    del ssh
    del transport, sftp
    del_pycache()
	
elif (sys.argv[1] == "start"):
	# Запуск сервера
	print("START  SERVER ...")
	ssh = connect_ssh()
	kill_processes(ssh)
	run_ssh_commands(ssh, ["cd " + remotepathServer, "chmod +x DOORS", "/DOORS/launch.sh &"])
	print("START  SERVER OK")
	ssh.close()
	del ssh
	del_pycache()
	
elif (sys.argv[1] == "stop"):
	# Остановка сервера
	print("STOP   SERVER ...")
	ssh = connect_ssh()
	kill_processes(ssh)
	print("STOP   SERVER OK")
	ssh.close()
	del ssh
	del_pycache()
	
elif (sys.argv[1] == "stat"):	
	# Проверка статуса сервера
    ssh = connect_ssh()
    stdin, stdout, stderr = ssh.exec_command("ps | grep DOORS")
    strings = stdout.read().decode().split('\n')
    
    is_run = False
    for str in strings:
        str = re.sub(r'\s+', ' ', str).strip()
        words = str.split(' ')
        try:
            if words[4] == "./DOORS":
                print("Server is RUN")
                is_run = True
        except IndexError:
            pass
    if not is_run:
        print("Server is STOP")
    
    ssh.close()
    del ssh, stdin, stdout, stderr
    del_pycache()
	
elif (sys.argv[1] == "git"):
	# Шифрование файлов GPG
    files_to_encrypt = ["secret.go", "secret.py", "Server/secret.go"]
    for file in files_to_encrypt:
        subprocess.run(["gpg", "-e", "-r", "Eugeny Goryachev", file])
        print(f"OK: gpg -e -r \"Eugeny Goryachev\" {file}")

    # Коммит в GitHub
    subprocess.run(["git", "add", "."])
    print("OK: git add .")
    
    comment = sys.argv[2] if len(sys.argv) == 3 else "-//-"
    subprocess.run(["git", "commit", "-m", comment])
    print(f"OK: git commit -m \"{comment}\"")
    
    subprocess.run(["git", "push", f"https://{secret.tokenGit}@github.com/{userGit}/{repoGit}.git"])
    print(f"OK: git push to {repoGit}")

    # Удаление зашифрованных файлов
    for file in files_to_encrypt:
        subprocess.run(["rm", f"{file}.gpg"])
        print(f"OK: rm {file}.gpg")
    
    del_pycache()

	
elif command == "docker":
    # Создание Docker контейнера
    docker_build()
    del_pycache()
    
    
elif command == "hub":
    # Отправка контейнера в Docker Hub
    docker_push()
    del_pycache()


elif command == "restart":
    # Перезапуск сервера
    print("RESTART  SERVER ...")
    ssh = connect_ssh()
    kill_processes(ssh)
    run_ssh_commands(ssh, ["cd " + remotepathServer, "chmod +x DOORS", "/DOORS/launch.sh &"])
    print("RESTART  SERVER OK")
    ssh.close()
    del ssh
    del_pycache()


else:
    print_error_message()
    del_pycache()
    sys.exit(1)
	
	
	
	

	


