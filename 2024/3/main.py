import re

# Day 3: Mull It Over
# Given an input of corrupted memory
# 1. Find the real multiplying instructions "mul(x,y)"
# 2. Ignore all other invalid instructions
# 3. Do the valid multiplications
# 4. Sum up all the values and return

FILE_PATH = "input.txt"
REGEX = r"mul[(]\d+,\d+[)]|don't[(][)]|do[(][)]"

class Solution:
    text = None
    is_disabled = False
    
    def __init__(self, file_path):
        try: 
            with open(file_path, "r") as f:
                self.text = f.read()
        except:
            raise Exception("Invalid file path")
        
    def solve(self):
        total = 0
        input_set = re.findall(REGEX, self.text)

        for instruction in input_set:
            if instruction == 'don\'t()':
                self.is_disabled = True
                continue
            
            if instruction == 'do()':
                self.is_disabled = False
                continue
            
            if not self.is_disabled:
                a, b = tuple(map(int, re.findall(r'\d+', instruction)))
                total += self.multiply(a, b)

        return total
    
    def multiply(self, a, b):
        return a * b

solution = Solution(FILE_PATH)
print(solution.solve())

