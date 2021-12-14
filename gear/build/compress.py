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
problem = False
problem_command = ""


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

started_thread_counter = 0
finished_thread_counter = 0

def compress(platform):
    global started_thread_counter, finished_thread_counter
    started_thread_counter += 1
    # os.chdir(f"cd .build/{platform[0]}.{platform[1]}")
    shutil.make_archive(f'.build/gear_{version}v_{platform[0]}.{platform[1]}.zip', 'zip', f'.build/{platform[0]}.{platform[1]}')
    # os.system(f"zip -r ../gear_{version}v_{platform[0]}.{platform[1]}.zip .")
    # os.chdir("./../../")
    finished_thread_counter += 1
    print(f'{Fore.YELLOW}{Style.BRIGHT}[compress]{Style.RESET_ALL}{Fore.YELLOW}: finished .build/{platform[0]}.{platform[1]} -> .build/gear_{version}v_{platform[0]}.{platform[1]}.zip')
    


for platform in platforms:
    print(f'{Fore.YELLOW}{Style.BRIGHT}[compress]{Style.RESET_ALL}{Fore.YELLOW}: .build/{platform[0]}.{platform[1]} -> .build/gear_{version}v_{platform[0]}.{platform[1]}.zip')
    # Thread(target = compress, args = (platform,)).start()    
    compress(platform)

print("Finished in:", time.time() - now, "s.")