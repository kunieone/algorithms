class TreeNode {
  L: TreeNode;
  R: TreeNode;
  P: TreeNode | null;
  value: number;
  lNum = 0;
  rNum = 0;
  insert(n: TreeNode): void {
    n.P = this;

    if (n.value > this.value) {
      this.R = n;
      this.rNum++;
    } else {
      this.L = n;
      this.lNum++;
    }

    leftRotate: () => {
      let tmp = this.L;
    };
  }
  getBalanceFactor = (n: TreeNode): number => this.lNum - this.rNum;
}
