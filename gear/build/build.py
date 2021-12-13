from threading import Thread
import os
import shutil
import zipfile
    
version = "alpha_1.0.0"

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
    os.system(f"go test {package}")
    print(f"😃 finished testing thread [package={package}]")

def compile(package, platform):
    os.system(f"env GOOS={platform[0]} GOARCH={platform[1]} go build -o .build/{platform[0]}.{platform[1]}/{package.split('/')[-1]}{platform[2]} {package}")

for package in packages_to_test:
    Thread(target=test, args=(package, )).start()
    print(f"test {package}")

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
        print(f"compile {directory}, os: {platform[0]}, arch: {platform[1]}")

for platform in platforms:
    print(f'compress .build/{platform[0]}.{platform[1]} ==> .build/gear_{version}v_{platform[0]}.{platform[1]}.zip')
    zipf = zipfile.ZipFile(f'.build/gear_{version}v_{platform[0]}.{platform[1]}.zip', 'w')
    zipf.write(f'.build/{platform[0]}.{platform[1]}')
    zipf.close()
