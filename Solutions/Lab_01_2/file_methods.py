def read(path):
    f_contents = list()
    with open(path, 'r') as f:
        for line in f:
            f_contents.append(line.replace('\n', ''))
    return f_contents


def write(path, contents):
    with open(path, 'w') as f:
        for str in contents:
            f.write(str + '\n')
