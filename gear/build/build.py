from threading import Thread
import os

package_prefix = "gear/cmd"
packages_to_test = ["gear/pkg/parser", "gear/pkg/lexer", "gear/pkg/docparser"]

platforms = [
    ('linux', '386'),
    ('linux', 'amd64'),
    ('linux', 'arm'),
    ('linux', 'arm64'),
    ('windows', '386'),
    ('windows', 'amd64')
]

print("⚙️  hike build started")

threads = []

def test(package):
    os.system(f"go test {package}")
    print(f"😃 finished testing thread [package={package}]")

def compile(package, platform):
    os.system(f"env GOOS={platform[0]} GOARCH={platform[1]} go build {package}")

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
        for file in os.listdir("."):
            if file.split(".")[-1] == directory:
                os.system(f"mv {file} .build/{platform[0]}.{platform[1]}/{file}")
