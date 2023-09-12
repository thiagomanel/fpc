import os
import sys

def wc(content):
    return len(content.split())

def wc_file(filename):
    try:
        with open(filename, 'r', encoding='latin-1') as f:
            file_content = f.read()
        return wc(file_content)
    except FileNotFoundError:
        return 0

def wc_dir(dir_path):
    count = 0
    for filename in os.listdir(dir_path):
        filepath = os.path.join(dir_path, filename)
        if os.path.isfile(filepath):
            count += wc_file(filepath)
        elif os.path.isdir(filepath):
            count += wc_dir(filepath)  # Chamada recursiva para diretórios
    return count

def main():
    if len(sys.argv) != 2:
        print("Usage: python", sys.argv[0], "root_directory_path")
        return
    root_path = os.path.abspath(sys.argv[1])
    count = wc_dir(root_path)
    print(count)

if __name__ == "__main__":
    main()
