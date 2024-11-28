import os
import sys

def read_passwords_from_dir(directory: str) -> dict:
    passwords = {}
    try:
        for file_name in os.listdir(directory):
            file_path = os.path.join(directory, file_name)
            if os.path.isfile(file_path):
                with open(file_path, 'r') as file:
                    passwords[file_name] = [line.strip() for line in file.readlines()]
    except FileNotFoundError:
        print(f"Erro: Diretório não encontrado no caminho {directory}.")
    except Exception as e:
        print(f"Ocorreu um erro: {e}")
    return passwords

def rot13_obfuscation(password: str) -> str:
    return "".join(
        chr((ord(char) - 65 + 13) % 26 + 65) if char.isupper() else
        chr((ord(char) - 97 + 13) % 26 + 97) if char.islower() else char
        for char in password
    )

def process_file_and_write(file_name: str, file_path: str, passwords: list):
    """Process and overwrite the file with ROT13-obfuscated passwords."""
    obfuscated_passwords = [rot13_obfuscation(password) for password in passwords]

    try:
        with open(file_path, 'w') as file:
            file.write("\n".join(obfuscated_passwords) + "\n")
            print(f"Processed and updated file: {file_name}")
    except Exception as e:
        print(f"Erro ao escrever no arquivo {file_path}: {e}")

def process_passwords_serially(directory: str, passwords_by_file: dict):
    for file_name, passwords in passwords_by_file.items():
        file_path = os.path.join(directory, file_name)
        process_file_and_write(file_name, file_path, passwords)

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Uso: python script.py <caminho_do_diretorio>")
        sys.exit(1)

    directory_path = sys.argv[1]

    passwords_by_file = read_passwords_from_dir(directory_path)

    process_passwords_serially(directory_path, passwords_by_file)

