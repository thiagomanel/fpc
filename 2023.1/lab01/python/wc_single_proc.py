import os
import string


def wc(content):
    count = 0
    inword = False

    for char in content:
        if char.isspace():
            if inword:
                inword = False
        else:
            if not inword:
                inword = True
                count += 1

    return count


def wc_file(filename):
    try:
        with open(filename, 'r', encoding='utf-8') as f:
            file_content = f.read()
            return wc(file_content)
    except FileNotFoundError:
        return -1


def wc_dir(dir_path):
    count = 0

    for filename in os.listdir(dir_path):
        filepath = os.path.join(dir_path, filename)
        if os.path.isfile(filepath):
            count += wc_file(filepath)

    return count


def main():
    import sys

    if len(sys.argv) != 2:
        print("Usage: python script_name.py root_directory_path")
        return

    root_path = sys.argv[1]
    count = 0

    for subdir in os.listdir(root_path):
        subdir_path = os.path.join(root_path, subdir)
        if os.path.isdir(subdir_path):
            count += wc_dir(subdir_path)

    print(count)


if __name__ == "__main__":
    main()
