import os
import sys

def wc(content):
	count = 0
	inword = False

	for char in content:
		if char.isspace():
			if inword:
				inword = False
				count += 1  # Increment the word count here
		else:
			if not inword:
				inword = True

	if inword:
		count += 1  # Increment the word count if the content ends with a word

	return count

def wc_file(filename):
	try:
		with open(filename, 'r', encoding='utf-8') as f:
			file_content = f.read()
			return wc(file_content)
	except FileNotFoundError as e:
		raise FileNotFoundError(f"File not found: {filename}") from e

def wc_dir(dir_path):
	count = 0

	for filename in os.listdir(dir_path):
		filepath = os.path.join(dir_path, filename)
		if os.path.isfile(filepath):
			count += wc_file(filepath)

	return count

def main():
	if len(sys.argv) != 2:
		print("Usage: python ", sys.argv[0], "root_directory_path")
		return

	root_path = os.path.abspath(sys.argv[1])
	count = wc_dir(root_path)

	print(count)

if __name__ == "__main__":
	main()
