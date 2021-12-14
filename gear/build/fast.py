from threading import Thread
import os
import sys
import time
import shutil
import zipfile
from colorama import *
from rich.console import Console

console = Console()

now = time.time()
started_thread_counter = 0
finished_thread_counter = 0
problem = False
problem_command = ""


package_prefix = "gear/cmd"
packages_to_test = ["gear/pkg/parser", "gear/pkg/lexer", "gear/pkg/docparser"]

print("⚙️  hike fast build started")

threads = []

def test(package):
	global started_thread_counter, finished_thread_counter, problem_command, problem
	started_thread_counter += 1
	now = time.time()
	r = os.system(f"go test {package}")
	if r != 0:
		problem = True
		# problem_command = f"env GOOS={platform[0]} GOARCH={platform[1]} go build -o .build/{platform[0]}.{platform[1]}/{package.split('/')[-1]}{platform[2]} {package}" 
	finished_thread_counter += 1
	print(f"{Fore.GREEN}{Style.BRIGHT}[test]{Style.RESET_ALL}{Fore.GREEN}: finish {package} in {time.time() - now} s{Fore.RESET}")

def compile(package):
	global started_thread_counter, finished_thread_counter, problem_command, problem
	started_thread_counter += 1
	now = time.time()
	r = os.system(f"go build {package}")
	if r != 0:
		problem = True
		problem_command = f"go build {package}" 
	finished_thread_counter += 1
	print(f"{Fore.GREEN}{Style.BRIGHT}[build]{Style.RESET_ALL}{Fore.GREEN}: finish {package.split('/')[-1]} in {time.time() - now} s{Fore.RESET}")

for package in packages_to_test:
	Thread(target=test, args=(package, )).start()
	print(f"{Fore.BLUE}{Style.BRIGHT}[test]{Style.RESET_ALL}{Fore.BLUE}: {package}{Fore.RESET}")

try:
	os.mkdir(".build")
except:
	pass

# def errors_check():
#     global problem, problem_command
#     while True:        
#         if problem == True:
#             print(f"found problem while executing `{problem_command}`")
#             os.kill(os.getpid(), signal.SIGINT)
#         if started_thread_counter == finished_thread_counter:
#             break

# Thread(target = errors_check).start()

for directory in os.listdir("cmd"):
	package = f"{package_prefix}/{directory}"
	thread = Thread(target = compile, args=(package, ))
	thread.start()
	print(f"{Fore.BLUE}{Style.BRIGHT}[build]{Style.RESET_ALL}{Fore.BLUE}: start building {directory}")

# while True:        
#     if problem == True:
#         os.system("clear")
#         print(f"found problem while executing `{problem_command}`")
#         sys.exit(1)
#     if started_thread_counter == finished_thread_counter:
# #         break
# def print_statusline(msg: str):
#     last_msg_length = len(print_statusline.last_msg) if hasattr(print_statusline, 'last_msg') else 0
#     print(' ' * last_msg_length, end='\r')
#     print(msg, end='\r')
#     sys.stdout.flush()  # Some say they needed this, I didn't.
#     print_statusline.last_msg = msg
while started_thread_counter != finished_thread_counter:
	with console.status(f"[bold yellow]waiting for build and test threads") as st:
		while started_thread_counter != finished_thread_counter:
			pass
