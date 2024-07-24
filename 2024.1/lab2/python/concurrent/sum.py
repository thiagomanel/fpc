import sys
import threading

def do_sum(path):
    _sum = 0
    with open(path, 'rb') as f:
        byte = f.read(1)
        while byte:
            _sum += int.from_bytes(byte, byteorder='big', signed=False)
            byte = f.read(1)
    print(path + " : " + str(_sum))

#many error could be raised error. we don't care       
if __name__ == "__main__":
    paths = sys.argv[1:]
    for path in paths:
        try:
            thread = threading.Thread(target=do_sum, args=(path,))
            thread.start()
        except Exception as e:
            print(f"Erro ao processar {path}: {e}")