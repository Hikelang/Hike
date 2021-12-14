from threading import Thread
import os
import sys
import time
import shutil
import zipfile
from colorama import *
from rich.console import Console

version = "alpha_0.1.0"
console = Console()

now = time.time()
started_thread_counter = 0
finished_thread_counter = 0
problem = False
problem_command = ""

def zipdir(path, ziph):
    # ziph is zipfile handle
    for root, dirs, files in os.walk(path):
        for file in files:
            if not file.endswith('.zip'):
                ziph.write(os.path.join(root, file), 
                       os.path.relpath(os.path.join(root, file), 
                                       os.path.join(path, '..')))

package_prefix = "gear/cmd"
packages_to_test = ["gear/pkg/parser", "gear/pkg/lexer", "gear/pkg/docparser"]

platforms = [
    # windows
    ('windows', 'amd64', '.exe'),
    ('windows', '386', '.exe'),

    # linux
    ('linux', '386', ''),
    ('linux', 'amd64', ''),
    ('linux', 'arm', ''),
    ('linux', 'arm64', ''),

    # ('darwin', '386', '.app'),
    # ('darwin', 'amd64', '.app'),
    # ('darwin', 'arm', '.app'),
    # ('darwin', 'arm64', '.app'),
]

print("⚙️  hike build started")

threads = []

def test(package):
    global started_thread_counter, finished_thread_counter, problem_command, problem
    started_thread_counter += 1
    now = time.time()
    r = os.system(f"go test {package}")
    if r != 0:
        problem = True
        problem_command = f"env GOOS={platform[0]} GOARCH={platform[1]} go build -o .build/{platform[0]}.{platform[1]}/{package.split('/')[-1]}{platform[2]} {package}" 
    finished_thread_counter += 1
    print(f"{Fore.GREEN}{Style.BRIGHT}[test]{Style.RESET_ALL}{Fore.GREEN}: finish {package} in {time.time() - now} s{Fore.RESET}")

def compile(package, platform):
    global started_thread_counter, finished_thread_counter, problem_command, problem
    started_thread_counter += 1
    now = time.time()
    r = os.system(f"env GOOS={platform[0]} GOARCH={platform[1]} go build -o .build/{platform[0]}.{platform[1]}/{package.split('/')[-1]}{platform[2]} {package}")
    if r != 0:
        problem = True
        problem_command = f"env GOOS={platform[0]} GOARCH={platform[1]} go build -o .build/{platform[0]}.{platform[1]}/{package.split('/')[-1]}{platform[2]} {package}" 
    finished_thread_counter += 1
    print(f"{Fore.GREEN}{Style.BRIGHT}[build]{Style.RESET_ALL}{Fore.GREEN}: finish {package.split('/')[-1]} [OS]: {platform[0]} [ARCH]: {platform[1]} in {time.time() - now} s{Fore.RESET}")

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
    for platform in platforms:
        try:
            os.mkdir(f".build/{platform[0]}.{platform[1]}")
        except:
            pass
        thread = Thread(target = compile, args=(package, platform))
        thread.start()
        print(f"{Fore.BLUE}{Style.BRIGHT}[build]{Style.RESET_ALL}{Fore.BLUE}: start building {directory} [OS]: {platform[0]} [ARCH]: {platform[1]}")

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

    # print(f"finished: {int(finished_thread_counter * 100 / started_thread_counter)} %", end="\r")
    # curses.setupterm(fd=sys.stdout.fileno()); print(hex(curses.tigetstr('really cool')));

print("======== DONE BUILD AND TEST THREADS ========")
print("[I] 3 SECOND SLEEP (for all threads to finish)")
time.sleep(3)
print("========== START COMPRESS THREADS ===========")
