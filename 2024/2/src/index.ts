import * as fs from 'fs';
import * as readline from 'readline';

// Day 2: Red-Nosed Reports
// Given an input file containing  many reports, one report per line.
// Each report is a list of numbers called levels that are separated by spaces.
// The engineers are trying to figure out which reports are safe.
// The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing.
// So, a report only counts as safe if both of the following are true:
// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.

const FILE_PATH = 'src/input.txt';

class Solution {
  countSafeReports(puzzleInput: number[][]): number {
    return puzzleInput.filter(report => this.isReportSafeWithDampener(report))
      .length;
  }

  isSafe(report: number[]): boolean {
    const isIncreasing = this.isIncreasing(report);
    const isDecreasing = this.isDecreasing(report);

    if (!(isIncreasing || isDecreasing)) {
      return false;
    }

    for (let i = 1; i < report.length; i++) {
      const diff = Math.abs(report[i] - report[i - 1]);
      if (diff === 0 || diff > 3) {
        return false;
      }
    }

    return true;
  }

  isReportSafeWithDampener(levels: number[]): boolean {
    if (this.isSafe(levels)) {
      return true;
    }

    for (let i = 0; i < levels.length; i++) {
      const modifiedLevels = levels.slice(0, i).concat(levels.slice(i + 1));
      if (this.isSafe(modifiedLevels)) {
        return true;
      }
    }

    return false;
  }

  private isDecreasing(report: number[]) {
    return report.every(
      (level, index) => index === 0 || level <= report[index - 1],
    );
  }

  private isIncreasing(report: number[]) {
    return report.every(
      (level, index) => index === 0 || level >= report[index - 1],
    );
  }

  async processInputFile(filePath: string): Promise<number[][]> {
    const fileStream = fs.createReadStream(filePath);
    const rl = readline.createInterface({
      input: fileStream,
      crlfDelay: Infinity,
    });

    const reports: number[][] = [];

    for await (const line of rl) {
      const report = line
        .trim()
        .split(/\s+/)
        .map(num => parseInt(num, 10));
      reports.push(report);
    }

    return reports;
  }

  async solve(filePath: string): Promise<number> {
    const puzzleInput = await this.processInputFile(filePath);

    return this.countSafeReports(puzzleInput);
  }
}

async function main() {
  const solution = new Solution();
  try {
    const safeReportCount = await solution.solve(FILE_PATH);
    console.log(`Number of safe reports: ${safeReportCount}`);
  } catch (error) {
    console.error('Error solving the puzzle:', error);
  }
}

main()
  .then(() => {
    console.log('Ran');
  })
  .catch(err => {
    console.log(err);
  });
