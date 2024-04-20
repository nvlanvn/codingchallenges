import sys
import os


arguments = sys.argv
argument_total = len(arguments)

if argument_total != 3:
    print("Argument invalid")
    sys.exit(0)


option = arguments[1]
file_name = arguments[2]
file_path = os.getcwd()+"/wc/"+file_name

print(option)
print(file_path)

try:
    file_stats = os.stat(file_path)
    print(f"{file_stats.st_size} {file_name}")
except FileNotFoundError:
    print("File Not Found")
