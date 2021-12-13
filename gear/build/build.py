from threading import Thread
import os
import time
import shutil
import zipfile
    
version = "alpha_1.0.0"

started_thread_counter = 0
finished_thread_counter = 0

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
    global started_thread_counter, finished_thread_counter
    started_thread_counter += 1
    now = time.time()
    os.system(f"go test {package}")
    finished_thread_counter += 1
    print(f"[T] finish {package} in {time.time() - now} s")

def compile(package, platform):
    global started_thread_counter, finished_thread_counter
    started_thread_counter += 1
    now = time.time()
    os.system(f"env GOOS={platform[0]} GOARCH={platform[1]} go build -o .build/{platform[0]}.{platform[1]}/{package.split('/')[-1]}{platform[2]} {package}")
    finished_thread_counter += 1
    print(f"[C] finish {package.split('/')[-1]} [OS]: {platform[0]} [ARCH]: {platform[1]} in {time.time() - now} s")

for package in packages_to_test:
    Thread(target=test, args=(package, )).start()
    print(f"[T] {package}")

try:
    os.mkdir(".build")
except:
    pass

for directory in os.listdir("cmd"):
    package = f"{package_prefix}/{directory}"
    for platform in platforms:
        try:
            os.mkdir(f".build/{platform[0]}.{platform[1]}")
        except:
            pass
        thread = Thread(target = compile, args=(package, platform))
        thread.start()
        print(f"[C] started {directory} [OS]: {platform[0]} [ARCH]: {platform[1]}")

while started_thread_counter != finished_thread_counter:
    pass

print(f"[I] finished building")

for platform in platforms:
    print(f'[C] .build/{platform[0]}.{platform[1]} ==> .build/gear_{version}v_{platform[0]}.{platform[1]}.zip')
    zipf = zipfile.ZipFile(f'.build/gear_{version}v_{platform[0]}.{platform[1]}.zip', 'w')
    zipf.write(f'.build/{platform[0]}.{platform[1]}')
    zipf.close()
