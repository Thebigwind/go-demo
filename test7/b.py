import subprocess
def test_exec():
    return_code = subprocess.run(['ls', '-al', '.'], stdout=subprocess.DEVNULL)
test_exec()