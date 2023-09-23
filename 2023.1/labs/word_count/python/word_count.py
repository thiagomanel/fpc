import os
import sys
from threading import Thread

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
    threads = []

    def process_file(filename):
        nonlocal count
        word_count = wc_file(os.path.join(dir_path, filename))
        count += word_count

    for filename in os.listdir(dir_path):
        if os.path.isfile(os.path.join(dir_path, filename)):
            thread = Thread(target=process_file, args=(filename,))
            threads.append(thread)
            thread.start()

    for thread in threads:
        thread.join()

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
