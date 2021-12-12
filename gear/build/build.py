from threading import Thread
import os

package_prefix = "gear/cmd"
packages_to_test = ["gear/pkg/parser", "gear/pkg/lexer", "gear/pkg/docparser"]

print("⚙️  hike build started")

threads = []

def test(package):
    os.system(f"go test {package}")

def compile(package):
    os.system(f"go build {package}")

for package in packages_to_test:
    Thread(target=test, args=(package, )).start()
    print(f"⚙️  >>> testing thread started [package={package}]")

for directory in os.listdir("cmd"):
    package = f"{package_prefix}/{directory}"
    print(f"⚙️  >>> found {package}")
    thread = Thread(target = compile, args=(package,))
    thread.start()
    print(f"⚙️  >>> compilation thread started [package={package}]")
